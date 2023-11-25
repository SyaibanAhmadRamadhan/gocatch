package jsql

import (
	"database/sql"
	"time"
)

// JSQLNullBool is a type that can handle bool null values in databases
type JSQLNullBool sql.NullBool

// NewJSQLNullBool creates a new JSQLNullBool object
func NewJSQLNullBool(i *bool) JSQLNullBool {
	if i == nil {
		return JSQLNullBool{}
	}

	return JSQLNullBool{
		Bool:  *i,
		Valid: true,
	}
}

// GetBool returns the bool value if it's valid, or false if it's not.
func (j JSQLNullBool) GetBool() bool {
	if !j.Valid {
		return false
	}

	return j.Bool
}

// Value returns a pointer to the false value if it's valid, or nil if it's not.
func (j JSQLNullBool) Value() *bool {
	if !j.Valid {
		return nil
	}

	return &j.Bool
}

// JSQLNullString is a type that can handle string null values in databases
type JSQLNullString sql.NullString

// NewJSQLNullString creates a new JSQLNullString object
func NewJSQLNullString(s *string) JSQLNullString {
	if s == nil {
		return JSQLNullString{}
	}

	return JSQLNullString{
		String: *s,
		Valid:  true,
	}
}

// GetString returns the String value if it's valid, or empty string if it's not.
func (j JSQLNullString) GetString() string {
	if !j.Valid {
		return ""
	}

	return j.String
}

// Value returns a pointer to the string value if it's valid, or nil if it's not.
func (j JSQLNullString) Value() *string {
	if !j.Valid {
		return nil
	}

	return &j.String
}

// JSQLNullFloat64 is a type that can handle float64 null values in databases
type JSQLNullFloat64 sql.NullFloat64

// NewJSQLNullFloat64 creates a new JSQLNullFloat64 object
func NewJSQLNullFloat64(f *float64) JSQLNullFloat64 {
	if f == nil {
		return JSQLNullFloat64{}
	}

	return JSQLNullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

// GetFloat64 returns the Float64 value if it's valid, or 0.0 if it's not.
func (j JSQLNullFloat64) GetFloat64() float64 {
	if !j.Valid {
		return 0.0
	}

	return j.Float64
}

// Value returns a pointer to the float64 value if it's valid, or nil if it's not.
func (j JSQLNullFloat64) Value() *float64 {
	if !j.Valid {
		return nil
	}

	return &j.Float64
}

// JSQLNullInt64 is a type that can handle int64 null values in databases
type JSQLNullInt64 sql.NullInt64

// NewJSQLNullInt64 creates a new JSQLNullInt64 object
func NewJSQLNullInt64(i *int64) JSQLNullInt64 {
	if i == nil {
		return JSQLNullInt64{}
	}

	return JSQLNullInt64{
		Int64: *i,
		Valid: true,
	}
}

// GetInt64 returns the GetInt64 value if it's valid, or 0 if it's not.
func (j JSQLNullInt64) GetInt64() int64 {
	if !j.Valid {
		return 0
	}

	return j.Int64
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j JSQLNullInt64) Value() *int64 {
	if !j.Valid {
		return nil
	}

	return &j.Int64
}

// JSQLNullInt32 is a type that can handle null values int64 in databases
type JSQLNullInt32 sql.NullInt32

// NewJSQLNullInt32 creates a new JSQLNullInt32 object
func NewJSQLNullInt32(i *int32) JSQLNullInt32 {
	if i == nil {
		return JSQLNullInt32{}
	}

	return JSQLNullInt32{
		Int32: *i,
		Valid: true,
	}
}

// GetInt32 returns the GetInt32 value if it's valid, or 0 if it's not.
func (j JSQLNullInt32) GetInt32() int32 {
	if !j.Valid {
		return 0
	}

	return j.Int32
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j JSQLNullInt32) Value() *int32 {
	if !j.Valid {
		return nil
	}

	return &j.Int32
}

// JSQLNullInt16 is a type that can handle null values int16 in databases
type JSQLNullInt16 sql.NullInt16

// NewJSQLNullInt16 creates a new JSQLNullInt16 object
func NewJSQLNullInt16(i *int16) JSQLNullInt16 {
	if i == nil {
		return JSQLNullInt16{}
	}

	return JSQLNullInt16{
		Int16: *i,
		Valid: true,
	}
}

// GetInt16 returns the GetInt16 value if it's valid, or 0 if it's not.
func (j JSQLNullInt16) GetInt16() int16 {
	if !j.Valid {
		return 0
	}

	return j.Int16
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j JSQLNullInt16) Value() *int16 {
	if !j.Valid {
		return nil
	}

	return &j.Int16
}

// JSQLNullByte is a type that can handle byte null values byte in databases
type JSQLNullByte sql.NullByte

// NewJSQLNullByte creates a new JSQLNullByte object
func NewJSQLNullByte(b *byte) JSQLNullByte {
	if b == nil {
		return JSQLNullByte{}
	}

	return JSQLNullByte{
		Byte:  *b,
		Valid: true,
	}
}

// GetByte returns the byte value if it's Valid, or 0 if it's not.
func (j JSQLNullByte) GetByte() byte {
	if !j.Valid {
		return 0
	}

	return j.Byte
}

// GetByteString returns the string from byte value if it's Valid, or empty string if it's not.
func (j JSQLNullByte) GetByteString() string {
	if !j.Valid {
		return ""
	}

	return string(j.Byte)
}

// Value returns a pointer to the byte value if it's Valid, or nil if it's not.
func (j JSQLNullByte) Value() *byte {
	if !j.Valid {
		return nil
	}

	return &j.Byte
}

// JSQLNullTime is a type that can handle time.Time null values in databases
type JSQLNullTime sql.NullTime

// NewJSQLNullTime creates a new JSQLNullTime
func NewJSQLNullTime(t *time.Time) JSQLNullTime {
	if t == nil {
		return JSQLNullTime{}
	}

	return JSQLNullTime{
		Time:  *t,
		Valid: true,
	}
}

// GetTime returns the time value if it's valid, or zero time if it's not.
func (j JSQLNullTime) GetTime() time.Time {
	if !j.Valid {
		return time.Time{}
	}

	return j.Time
}

// Value returns a pointer to the time value if it's valid, or nil if it's not.
func (j JSQLNullTime) Value() *time.Time {
	if !j.Valid {
		return nil
	}

	return &j.Time
}
