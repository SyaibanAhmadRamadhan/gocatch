package Jsql

import (
	"database/sql"
	"time"
)

// NullBool is a type that can handle bool null values in databases
type NullBool sql.NullBool

// NewNullBool creates a new NullBool object
func NewNullBool(i *bool) NullBool {
	if i == nil {
		return NullBool{}
	}

	return NullBool{
		Bool:  *i,
		Valid: true,
	}
}

// GetBool returns the bool value if it's valid, or false if it's not.
func (j NullBool) GetBool() bool {
	if !j.Valid {
		return false
	}

	return j.Bool
}

// Value returns a pointer to the false value if it's valid, or nil if it's not.
func (j NullBool) Value() *bool {
	if !j.Valid {
		return nil
	}

	return &j.Bool
}

// NullString is a type that can handle string null values in databases
type NullString sql.NullString

// NewNullString creates a new NullString object
func NewNullString(s *string) NullString {
	if s == nil {
		return NullString{}
	}

	return NullString{
		String: *s,
		Valid:  true,
	}
}

// GetString returns the String value if it's valid, or empty string if it's not.
func (j NullString) GetString() string {
	if !j.Valid {
		return ""
	}

	return j.String
}

// Value returns a pointer to the string value if it's valid, or nil if it's not.
func (j NullString) Value() *string {
	if !j.Valid {
		return nil
	}

	return &j.String
}

// NullFloat64 is a type that can handle float64 null values in databases
type NullFloat64 sql.NullFloat64

// NewNullFloat64 creates a new NullFloat64 object
func NewNullFloat64(f *float64) NullFloat64 {
	if f == nil {
		return NullFloat64{}
	}

	return NullFloat64{
		Float64: *f,
		Valid:   true,
	}
}

// GetFloat64 returns the Float64 value if it's valid, or 0.0 if it's not.
func (j NullFloat64) GetFloat64() float64 {
	if !j.Valid {
		return 0.0
	}

	return j.Float64
}

// Value returns a pointer to the float64 value if it's valid, or nil if it's not.
func (j NullFloat64) Value() *float64 {
	if !j.Valid {
		return nil
	}

	return &j.Float64
}

// NullInt64 is a type that can handle int64 null values in databases
type NullInt64 sql.NullInt64

// NewNullInt64 creates a new NullInt64 object
func NewNullInt64(i *int64) NullInt64 {
	if i == nil {
		return NullInt64{}
	}

	return NullInt64{
		Int64: *i,
		Valid: true,
	}
}

// GetInt64 returns the GetInt64 value if it's valid, or 0 if it's not.
func (j NullInt64) GetInt64() int64 {
	if !j.Valid {
		return 0
	}

	return j.Int64
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j NullInt64) Value() *int64 {
	if !j.Valid {
		return nil
	}

	return &j.Int64
}

// NullInt32 is a type that can handle null values int64 in databases
type NullInt32 sql.NullInt32

// NewNullInt32 creates a new NullInt32 object
func NewNullInt32(i *int32) NullInt32 {
	if i == nil {
		return NullInt32{}
	}

	return NullInt32{
		Int32: *i,
		Valid: true,
	}
}

// GetInt32 returns the GetInt32 value if it's valid, or 0 if it's not.
func (j NullInt32) GetInt32() int32 {
	if !j.Valid {
		return 0
	}

	return j.Int32
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j NullInt32) Value() *int32 {
	if !j.Valid {
		return nil
	}

	return &j.Int32
}

// NullInt16 is a type that can handle null values int16 in databases
type NullInt16 sql.NullInt16

// NewNullInt16 creates a new NullInt16 object
func NewNullInt16(i *int16) NullInt16 {
	if i == nil {
		return NullInt16{}
	}

	return NullInt16{
		Int16: *i,
		Valid: true,
	}
}

// GetInt16 returns the GetInt16 value if it's valid, or 0 if it's not.
func (j NullInt16) GetInt16() int16 {
	if !j.Valid {
		return 0
	}

	return j.Int16
}

// Value returns a pointer to the int64 value if it's valid, or nil if it's not.
func (j NullInt16) Value() *int16 {
	if !j.Valid {
		return nil
	}

	return &j.Int16
}

// NullByte is a type that can handle byte null values byte in databases
type NullByte sql.NullByte

// NewNullByte creates a new NullByte object
func NewNullByte(b *byte) NullByte {
	if b == nil {
		return NullByte{}
	}

	return NullByte{
		Byte:  *b,
		Valid: true,
	}
}

// GetByte returns the byte value if it's Valid, or 0 if it's not.
func (j NullByte) GetByte() byte {
	if !j.Valid {
		return 0
	}

	return j.Byte
}

// GetByteString returns the string from byte value if it's Valid, or empty string if it's not.
func (j NullByte) GetByteString() string {
	if !j.Valid {
		return ""
	}

	return string(j.Byte)
}

// Value returns a pointer to the byte value if it's Valid, or nil if it's not.
func (j NullByte) Value() *byte {
	if !j.Valid {
		return nil
	}

	return &j.Byte
}

// NullTime is a type that can handle time.Time null values in databases
type NullTime sql.NullTime

// NewNullTime creates a new NullTime
func NewNullTime(t *time.Time) NullTime {
	if t == nil {
		return NullTime{}
	}

	return NullTime{
		Time:  *t,
		Valid: true,
	}
}

// GetTime returns the time value if it's valid, or zero time if it's not.
func (j NullTime) GetTime() time.Time {
	if !j.Valid {
		return time.Time{}
	}

	return j.Time
}

// Value returns a pointer to the time value if it's valid, or nil if it's not.
func (j NullTime) Value() *time.Time {
	if !j.Valid {
		return nil
	}

	return &j.Time
}
