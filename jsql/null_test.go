package jsql

import (
	"testing"
	"time"
)

// Test cases for JSQLNullBool
func TestNewJSQLNullBool(t *testing.T) {
	i := true
	jnb := NewJSQLNullBool(&i)
	if !jnb.Valid || !jnb.Bool {
		t.FailNow()
	}

	i = false
	jnb = NewJSQLNullBool(&i)
	if !jnb.Valid || jnb.Bool {
		t.FailNow()
	}

	jnb = NewJSQLNullBool(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullBool_Value(t *testing.T) {
	i := true
	jnb := NewJSQLNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = false
	jnb = NewJSQLNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullBool(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullBool_GetBool(t *testing.T) {
	i := true
	jnb := NewJSQLNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	i = false
	jnb = NewJSQLNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullBool(nil)
	if p := jnb.GetBool(); p != false {
		t.FailNow()
	}
}

// Test cases for JSQLNullString
func TestNewJSQLNullString(t *testing.T) {
	s := "rama"
	jnb := NewJSQLNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	s = ""
	jnb = NewJSQLNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	jnb = NewJSQLNullString(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullString_Value(t *testing.T) {
	i := "rama"
	jnb := NewJSQLNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewJSQLNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullString(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullString_GetString(t *testing.T) {
	i := "rama"
	jnb := NewJSQLNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewJSQLNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullString(nil)
	if p := jnb.GetString(); p != "" {
		t.FailNow()
	}
}

// Test cases for JSQLNullFloat64
func TestNewJSQLNullFloat64(t *testing.T) {
	f := 8.5
	jnb := NewJSQLNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewJSQLNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	jnb = NewJSQLNullFloat64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullFloat64_Value(t *testing.T) {
	f := 8.5
	jnb := NewJSQLNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewJSQLNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	jnb = NewJSQLNullFloat64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullFloat64_GetFloat64(t *testing.T) {
	f := 8.5
	jnb := NewJSQLNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewJSQLNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	jnb = NewJSQLNullFloat64(nil)
	if p := jnb.GetFloat64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for JSQLNullInt64
func TestNewJSQLNullInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewJSQLNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewJSQLNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullInt64_Value(t *testing.T) {
	var i int64 = 8
	jnb := NewJSQLNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullInt64_GetInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewJSQLNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt64(nil)
	if p := jnb.GetInt64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for JSQLNullInt32
func TestNewJSQLNullInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewJSQLNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewJSQLNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt32(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullInt32_Value(t *testing.T) {
	var i int32 = 8
	jnb := NewJSQLNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt32(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullInt32_GetInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewJSQLNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt32(nil)
	if p := jnb.GetInt32(); p != 0 {
		t.FailNow()
	}
}

// Test cases for JSQLNullInt16
func TestNewJSQLNullInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewJSQLNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewJSQLNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt16(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullInt16_Value(t *testing.T) {
	var i int16 = 8
	jnb := NewJSQLNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt16(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullInt16_GetInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewJSQLNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewJSQLNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	jnb = NewJSQLNullInt16(nil)
	if p := jnb.GetInt16(); p != 0 {
		t.FailNow()
	}
}

// Test cases for JSQLNullByte
func TestNewJSQLNullByte(t *testing.T) {
	var i byte = 'r'
	jnb := NewJSQLNullByte(&i)
	if !jnb.Valid || jnb.GetByte() != i {
		t.FailNow()
	}

	var nullByte byte
	jnb = NewJSQLNullByte(&nullByte)
	if !jnb.Valid || jnb.GetByte() != nullByte {
		t.FailNow()
	}

	jnb = NewJSQLNullByte(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestJSQLNullByte_Value(t *testing.T) {
	var i byte = 'a'
	jnb := NewJSQLNullByte(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewJSQLNullByte(&nullByte)
	if p := jnb.Value(); p == nil || *p != nullByte {
		t.FailNow()
	}

	jnb = NewJSQLNullByte(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullByte_GetByte(t *testing.T) {
	var i byte = 'a'
	jnb := NewJSQLNullByte(&i)
	if p := jnb.GetByte(); p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewJSQLNullByte(&nullByte)
	if p := jnb.GetByte(); p != nullByte {
		t.FailNow()
	}

	jnb = NewJSQLNullByte(nil)
	if p := jnb.GetByte(); p != 0 {
		t.FailNow()
	}
}

func TestJSQLNullByte_GetByteString(t *testing.T) {
	var i byte = 'a'
	jnb := NewJSQLNullByte(&i)
	// convert to string
	if p := jnb.GetByteString(); p != string(i) {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewJSQLNullByte(&nullByte)
	if p := jnb.GetByteString(); p != string(nullByte) {
		t.FailNow()
	}

	jnb = NewJSQLNullByte(nil)
	if p := jnb.GetByteString(); p != "" {
		t.FailNow()
	}
}

// Test cases for JSQLNullTime
func TestNewJSQLNullTime(t *testing.T) {
	now := time.Now()
	jnt := NewJSQLNullTime(&now)
	if !jnt.Valid || jnt.GetTime() != now {
		t.FailNow()
	}

	jnt = NewJSQLNullTime(nil)
	if jnt.Valid {
		t.FailNow()
	}
}

func TestJSQLNullTime_Value(t *testing.T) {
	now := time.Now()
	jnt := NewJSQLNullTime(&now)
	if p := jnt.Value(); p == nil || *p != now {
		t.FailNow()
	}

	jnt = NewJSQLNullTime(nil)
	if p := jnt.Value(); p != nil {
		t.FailNow()
	}
}

func TestJSQLNullTime_GetTime(t *testing.T) {
	now := time.Now()
	jnt := NewJSQLNullTime(&now)
	if p := jnt.GetTime(); !p.Equal(now) {
		t.FailNow()
	}

	jnt = NewJSQLNullTime(nil)
	if p := jnt.GetTime(); !p.IsZero() {
		t.FailNow()
	}
}
