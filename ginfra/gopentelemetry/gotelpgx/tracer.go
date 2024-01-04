package gotelpgx

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	tracer trace.Tracer
	attrs  []attribute.KeyValue
}

func NewTracer() *Tracer {
	tp := otel.GetTracerProvider()
	return &Tracer{
		tracer: tp.Tracer(TracerName, trace.WithInstrumentationVersion(InstrumentVersion)),
		attrs:  []attribute.KeyValue{semconv.DBSystemPostgreSQL},
	}
}

func recordError(span trace.Span, err error) {
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func connAttrFromCfgPgx(config *pgx.ConnConfig) []trace.SpanStartOption {
	if config != nil {
		return []trace.SpanStartOption{
			trace.WithAttributes(
				attribute.String("database", config.Database),
				semconv.NetPeerName(config.Host),
				semconv.NetPeerPort(int(config.Port)),
				semconv.DBUser(config.User),
			),
		}
	}

	return nil
}

func makeParamAttr(args []any) attribute.KeyValue {
	ss := make([]string, len(args))
	for i := range args {
		ss[i] = fmt.Sprintf("%+v", args[i])
	}

	attrKey := attribute.Key("pgx.query.parameter")
	return attrKey.StringSlice(ss)
}

func (t *Tracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx
	}

	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
		trace.WithAttributes(semconv.DBStatement(data.SQL)),
		trace.WithAttributes(makeParamAttr(data.Args)),
	}

	if conn != nil {
		opts = append(opts, connAttrFromCfgPgx(conn.Config())...)
	}

	spanName := "query | " + data.SQL
	ctx, _ = t.tracer.Start(ctx, spanName, opts...)

	return ctx
}

func (t *Tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	span := trace.SpanFromContext(ctx)
	recordError(span, data.Err)

	rowsAffectedKey := attribute.Key("pgx.rows_affected")
	if data.Err != nil {
		span.SetAttributes(rowsAffectedKey.Int64(data.CommandTag.RowsAffected()))
	}

	span.End()
}

func (t *Tracer) TraceCopyFromStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceCopyFromStartData) context.Context {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx
	}

	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
		trace.WithAttributes(semconv.DBSQLTable(data.TableName.Sanitize())),
	}

	if conn != nil {
		attrColumnKey := attribute.Key("copy.columns")
		opts = append(opts, connAttrFromCfgPgx(conn.Config())...)
		opts = append(opts, trace.WithAttributes(
			attrColumnKey.StringSlice(data.ColumnNames),
		))

	}

	ctx, _ = t.tracer.Start(ctx, "copy_from "+data.TableName.Sanitize(), opts...)

	return ctx
}

func (t *Tracer) TraceCopyFromEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceCopyFromEndData) {
	span := trace.SpanFromContext(ctx)
	recordError(span, data.Err)

	rowsAffectedKey := attribute.Key("pgx.rows_affected")
	if data.Err != nil {
		span.SetAttributes(rowsAffectedKey.Int64(data.CommandTag.RowsAffected()))
	}

	span.End()
}

func (t *Tracer) TraceBatchStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx
	}

	var size int
	if b := data.Batch; b != nil {
		size = b.Len()
	}

	batchSizeKey := attribute.Key("pgx.batch_size")
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
		trace.WithAttributes(batchSizeKey.Int(size)),
	}

	if conn != nil {
		opts = append(opts, connAttrFromCfgPgx(conn.Config())...)
	}

	ctx, _ = t.tracer.Start(ctx, "batch_start", opts...)

	return ctx
}

func (t *Tracer) TraceBatchQuery(ctx context.Context, conn *pgx.Conn, data pgx.TraceBatchQueryData) {
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
	}

	if conn != nil {
		opts = append(opts, connAttrFromCfgPgx(conn.Config())...)
		opts = append(opts, trace.WithAttributes(semconv.DBStatement(data.SQL)))
		opts = append(opts, trace.WithAttributes(makeParamAttr(data.Args)))
	}

	spanName := "batch_query | " + data.SQL

	_, span := t.tracer.Start(ctx, spanName, opts...)
	recordError(span, data.Err)

	span.End()
}

func (t *Tracer) TraceBatchEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceBatchEndData) {
	span := trace.SpanFromContext(ctx)
	recordError(span, data.Err)

	span.End()
}

func (t *Tracer) TraceConnectStart(ctx context.Context, data pgx.TraceConnectStartData) context.Context {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx
	}

	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
	}

	if data.ConnConfig != nil {
		opts = append(opts, connAttrFromCfgPgx(data.ConnConfig)...)
	}

	ctx, _ = t.tracer.Start(ctx, "connect", opts...)

	return ctx
}

func (t *Tracer) TraceConnectEnd(ctx context.Context, data pgx.TraceConnectEndData) {
	span := trace.SpanFromContext(ctx)
	recordError(span, data.Err)

	span.End()
}

func (t *Tracer) TracePrepareStart(ctx context.Context, conn *pgx.Conn, data pgx.TracePrepareStartData) context.Context {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx
	}

	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(t.attrs...),
	}

	prepareStmtNameKey := attribute.Key("pgx.prepared_statement")
	if data.Name != "" {
		trace.WithAttributes(prepareStmtNameKey.String(data.Name))
	}

	if conn != nil {
		opts = append(opts, connAttrFromCfgPgx(conn.Config())...)
		opts = append(opts, trace.WithAttributes(semconv.DBStatement(data.SQL)))
	}

	spanName := "prepared_statement | " + data.SQL

	ctx, _ = t.tracer.Start(ctx, spanName, opts...)

	return ctx
}

func (t *Tracer) TracePrepareEnd(ctx context.Context, _ *pgx.Conn, data pgx.TracePrepareEndData) {
	span := trace.SpanFromContext(ctx)
	recordError(span, data.Err)

	span.End()
}
