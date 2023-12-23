package gdb

import (
	"database/sql"
	"time"
)

// SqlNullBool is a type that can handle bool null values in databases
type SqlNullBool sql.NullBool

// NewSqlNullBool creates a new SqlNullBool object
func NewSqlNullBool(i *bool) SqlNullBool {
	if i == nil {
		return SqlNullBool{}
	}

	return SqlNullBool{
		Bool:  *i,
		Valid: true,
	}
}

// GetBool returns the bool value if it's valid, or false if it's not.
func (j SqlNullBool) GetBool() bool {
	if !j.Valid {
		return false
	}

	return j.Bool
}

// Value returns a pointer to the false value if it's valid, or nil if it's not.
func (j SqlNullBool) Value() *bool {
	if !j.Valid {
		return nil
	}

	return &j.Bool
}

// SqlNullString is a type that can handle string null values in databases
type SqlNullString sql.NullString

// NewSqlNullString creates a new SqlNullString object
func NewSqlNullString(s *string) SqlNullString {
	if s == nil {
		return SqlNullString{}
	}

	return SqlNullString{
		String: *s,
		Valid:  true,
	}
}

// GetString returns the String value if it's valid, or empty string if it's not.
func (j SqlNullString) GetString() string {
	if !j.Valid {
		return ""
	}

	return j.String
}

// Value returns a pointer to the string value if it's valid, or nil if it's not.
func (j SqlNullString) Value() *string {
	if !j.Valid {
		return nil
	}

	return &j.String
}

// SqlNullFloat64 is a type that can handle float64 null values in databases
type SqlNullFloat64 sql.NullFloat64

// NewSqlNullFloat64 creates a new SqlNullFloat64 object
func NewSqlNullFloat64(f *float64) SqlNullFloat64 {
	if f == nil {
		return SqlNullFloat64{}
	}

	return SqlNullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

// GetFloat64 returns the Float64 value if it's valid, or 0.0 if it's not.
func (j SqlNullFloat64) GetFloat64() float64 {
	if !j.Valid {
		return 0.0
	}

	return j.Float64
}

// Value returns a pointer to the float64 value if it's valid, or nil if it's not.
func (j SqlNullFloat64) Value() *float64 {
	if !j.Valid {
		return nil
	}

	return &j.Float64
}

// SqlNullInt64 is a type that can handle int64 null values in databases
type SqlNullInt64 sql.NullInt64

// NewSqlNullInt64 creates a new SqlNullInt64 object
func NewSqlNullInt64(i *int64) SqlNullInt64 {
	if i == nil {
		return SqlNullInt64{}
	}

	return SqlNullInt64{
		Int64: *i,
		Valid: true,
	}
}

// GetInt64 returns the GetInt64 value if it's valid, or 0 if it's not.
func (j SqlNullInt64) GetInt64() int64 {
	if !j.Valid {
		return 0
	}

	return j.Int64
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j SqlNullInt64) Value() *int64 {
	if !j.Valid {
		return nil
	}

	return &j.Int64
}

// SqlNullInt32 is a type that can handle null values int64 in databases
type SqlNullInt32 sql.NullInt32

// NewSqlNullInt32 creates a new SqlNullInt32 object
func NewSqlNullInt32(i *int32) SqlNullInt32 {
	if i == nil {
		return SqlNullInt32{}
	}

	return SqlNullInt32{
		Int32: *i,
		Valid: true,
	}
}

// GetInt32 returns the GetInt32 value if it's valid, or 0 if it's not.
func (j SqlNullInt32) GetInt32() int32 {
	if !j.Valid {
		return 0
	}

	return j.Int32
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j SqlNullInt32) Value() *int32 {
	if !j.Valid {
		return nil
	}

	return &j.Int32
}

// SqlNullInt16 is a type that can handle null values int16 in databases
type SqlNullInt16 sql.NullInt16

// NewSqlNullInt16 creates a new SqlNullInt16 object
func NewSqlNullInt16(i *int16) SqlNullInt16 {
	if i == nil {
		return SqlNullInt16{}
	}

	return SqlNullInt16{
		Int16: *i,
		Valid: true,
	}
}

// GetInt16 returns the GetInt16 value if it's valid, or 0 if it's not.
func (j SqlNullInt16) GetInt16() int16 {
	if !j.Valid {
		return 0
	}

	return j.Int16
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j SqlNullInt16) Value() *int16 {
	if !j.Valid {
		return nil
	}

	return &j.Int16
}

// SqlNullByte is a type that can handle byte null values byte in databases
type SqlNullByte sql.NullByte

// NewSqlNullByte creates a new SqlNullByte object
func NewSqlNullByte(b *byte) SqlNullByte {
	if b == nil {
		return SqlNullByte{}
	}

	return SqlNullByte{
		Byte:  *b,
		Valid: true,
	}
}

// GetByte returns the byte value if it's Valid, or 0 if it's not.
func (j SqlNullByte) GetByte() byte {
	if !j.Valid {
		return 0
	}

	return j.Byte
}

// GetByteString returns the string from byte value if it's Valid, or empty string if it's not.
func (j SqlNullByte) GetByteString() string {
	if !j.Valid {
		return ""
	}

	return string(j.Byte)
}

// Value returns a pointer to the byte value if it's Valid, or nil if it's not.
func (j SqlNullByte) Value() *byte {
	if !j.Valid {
		return nil
	}

	return &j.Byte
}

// SqlNullTime is a type that can handle time.Time null values in databases
type SqlNullTime sql.NullTime

// NewSqlNullTime creates a new SqlNullTime
func NewSqlNullTime(t *time.Time) SqlNullTime {
	if t == nil {
		return SqlNullTime{}
	}

	return SqlNullTime{
		Time:  *t,
		Valid: true,
	}
}

// GetTime returns the time value if it's valid, or zero time if it's not.
func (j SqlNullTime) GetTime() time.Time {
	if !j.Valid {
		return time.Time{}
	}

	return j.Time
}

// Value returns a pointer to the time value if it's valid, or nil if it's not.
func (j SqlNullTime) Value() *time.Time {
	if !j.Valid {
		return nil
	}

	return &j.Time
}
