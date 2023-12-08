package gredis

import (
	"context"

	"github.com/redis/rueidis"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

type RueidisAdapter struct {
	Conn rueidis.Client
}

func OpenConnRueidis(opt rueidis.ClientOption) *RueidisAdapter {
	conn, err := rueidis.NewClient(opt)
	gcommon.PanicIfError(err)

	return &RueidisAdapter{
		Conn: conn,
	}
}

func (r *RueidisAdapter) Hset(ctx context.Context, key, field, value string) rueidis.RedisResult {
	return r.Conn.Do(ctx, r.Conn.B().Hset().Key(key).FieldValue().FieldValue(field, value).Build())
}

func (r *RueidisAdapter) Hsetnx(ctx context.Context, key, field, value string) rueidis.RedisResult {
	return r.Conn.Do(ctx, r.Conn.B().Hsetnx().Key(key).Field(field).Value(value).Build())
}
