package gdb

import (
	"testing"
	"time"
)

// Test cases for SqlNullBool
func TestNewNullBool(t *testing.T) {
	i := true
	jnb := NewSqlNullBool(&i)
	if !jnb.Valid || !jnb.Bool {
		t.FailNow()
	}

	i = false
	jnb = NewSqlNullBool(&i)
	if !jnb.Valid || jnb.Bool {
		t.FailNow()
	}

	jnb = NewSqlNullBool(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullBool_Value(t *testing.T) {
	i := true
	jnb := NewSqlNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = false
	jnb = NewSqlNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewSqlNullBool(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullBool_GetBool(t *testing.T) {
	i := true
	jnb := NewSqlNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	i = false
	jnb = NewSqlNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	jnb = NewSqlNullBool(nil)
	if p := jnb.GetBool(); p != false {
		t.FailNow()
	}
}

// Test cases for SqlNullString
func TestNewNullString(t *testing.T) {
	s := "rama"
	jnb := NewSqlNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	s = ""
	jnb = NewSqlNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	jnb = NewSqlNullString(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullString_Value(t *testing.T) {
	i := "rama"
	jnb := NewSqlNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewSqlNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewSqlNullString(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullString_GetString(t *testing.T) {
	i := "rama"
	jnb := NewSqlNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewSqlNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	jnb = NewSqlNullString(nil)
	if p := jnb.GetString(); p != "" {
		t.FailNow()
	}
}

// Test cases for SqlNullFloat64
func TestNewNullFloat64(t *testing.T) {
	f := 8.5
	jnb := NewSqlNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewSqlNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	jnb = NewSqlNullFloat64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullFloat64_Value(t *testing.T) {
	f := 8.5
	jnb := NewSqlNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewSqlNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	jnb = NewSqlNullFloat64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullFloat64_GetFloat64(t *testing.T) {
	f := 8.5
	jnb := NewSqlNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewSqlNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	jnb = NewSqlNullFloat64(nil)
	if p := jnb.GetFloat64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for SqlNullInt64
func TestNewNullInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewSqlNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewSqlNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt64_Value(t *testing.T) {
	var i int64 = 8
	jnb := NewSqlNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt64_GetInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewSqlNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt64(nil)
	if p := jnb.GetInt64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for SqlNullInt32
func TestNewNullInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewSqlNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewSqlNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt32(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt32_Value(t *testing.T) {
	var i int32 = 8
	jnb := NewSqlNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt32(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt32_GetInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewSqlNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt32(nil)
	if p := jnb.GetInt32(); p != 0 {
		t.FailNow()
	}
}

// Test cases for SqlNullInt16
func TestNewNullInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewSqlNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewSqlNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt16(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt16_Value(t *testing.T) {
	var i int16 = 8
	jnb := NewSqlNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt16(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt16_GetInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewSqlNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewSqlNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	jnb = NewSqlNullInt16(nil)
	if p := jnb.GetInt16(); p != 0 {
		t.FailNow()
	}
}

// Test cases for SqlNullByte
func TestNewNullByte(t *testing.T) {
	var i byte = 'r'
	jnb := NewSqlNullByte(&i)
	if !jnb.Valid || jnb.GetByte() != i {
		t.FailNow()
	}

	var nullByte byte
	jnb = NewSqlNullByte(&nullByte)
	if !jnb.Valid || jnb.GetByte() != nullByte {
		t.FailNow()
	}

	jnb = NewSqlNullByte(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullByte_Value(t *testing.T) {
	var i byte = 'a'
	jnb := NewSqlNullByte(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewSqlNullByte(&nullByte)
	if p := jnb.Value(); p == nil || *p != nullByte {
		t.FailNow()
	}

	jnb = NewSqlNullByte(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullByte_GetByte(t *testing.T) {
	var i byte = 'a'
	jnb := NewSqlNullByte(&i)
	if p := jnb.GetByte(); p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewSqlNullByte(&nullByte)
	if p := jnb.GetByte(); p != nullByte {
		t.FailNow()
	}

	jnb = NewSqlNullByte(nil)
	if p := jnb.GetByte(); p != 0 {
		t.FailNow()
	}
}

func TestNullByte_GetByteString(t *testing.T) {
	var i byte = 'a'
	jnb := NewSqlNullByte(&i)
	// convert to string
	if p := jnb.GetByteString(); p != string(i) {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewSqlNullByte(&nullByte)
	if p := jnb.GetByteString(); p != string(nullByte) {
		t.FailNow()
	}

	jnb = NewSqlNullByte(nil)
	if p := jnb.GetByteString(); p != "" {
		t.FailNow()
	}
}

// Test cases for SqlNullTime
func TestNewNullTime(t *testing.T) {
	now := time.Now()
	jnt := NewSqlNullTime(&now)
	if !jnt.Valid || jnt.GetTime() != now {
		t.FailNow()
	}

	jnt = NewSqlNullTime(nil)
	if jnt.Valid {
		t.FailNow()
	}
}

func TestNullTime_Value(t *testing.T) {
	now := time.Now()
	jnt := NewSqlNullTime(&now)
	if p := jnt.Value(); p == nil || *p != now {
		t.FailNow()
	}

	jnt = NewSqlNullTime(nil)
	if p := jnt.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullTime_GetTime(t *testing.T) {
	now := time.Now()
	jnt := NewSqlNullTime(&now)
	if p := jnt.GetTime(); !p.Equal(now) {
		t.FailNow()
	}

	jnt = NewSqlNullTime(nil)
	if p := jnt.GetTime(); !p.IsZero() {
		t.FailNow()
	}
}
