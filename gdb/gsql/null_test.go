package gsql

import (
	"testing"
	"time"
)

// Test cases for NullBool
func TestNewNullBool(t *testing.T) {
	i := true
	jnb := NewNullBool(&i)
	if !jnb.Valid || !jnb.Bool {
		t.FailNow()
	}

	i = false
	jnb = NewNullBool(&i)
	if !jnb.Valid || jnb.Bool {
		t.FailNow()
	}

	jnb = NewNullBool(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullBool_Value(t *testing.T) {
	i := true
	jnb := NewNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = false
	jnb = NewNullBool(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewNullBool(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullBool_GetBool(t *testing.T) {
	i := true
	jnb := NewNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	i = false
	jnb = NewNullBool(&i)
	if p := jnb.GetBool(); p != i {
		t.FailNow()
	}

	jnb = NewNullBool(nil)
	if p := jnb.GetBool(); p != false {
		t.FailNow()
	}
}

// Test cases for NullString
func TestNewNullString(t *testing.T) {
	s := "rama"
	jnb := NewNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	s = ""
	jnb = NewNullString(&s)
	if !jnb.Valid || jnb.String != s {
		t.FailNow()
	}

	jnb = NewNullString(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullString_Value(t *testing.T) {
	i := "rama"
	jnb := NewNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewNullString(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewNullString(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullString_GetString(t *testing.T) {
	i := "rama"
	jnb := NewNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	i = ""
	jnb = NewNullString(&i)
	if p := jnb.GetString(); p != i {
		t.FailNow()
	}

	jnb = NewNullString(nil)
	if p := jnb.GetString(); p != "" {
		t.FailNow()
	}
}

// Test cases for NullFloat64
func TestNewNullFloat64(t *testing.T) {
	f := 8.5
	jnb := NewNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewNullFloat64(&f)
	if !jnb.Valid || jnb.Float64 != f {
		t.FailNow()
	}

	jnb = NewNullFloat64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullFloat64_Value(t *testing.T) {
	f := 8.5
	jnb := NewNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewNullFloat64(&f)
	if p := jnb.Value(); p == nil || *p != f {
		t.FailNow()
	}

	jnb = NewNullFloat64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullFloat64_GetFloat64(t *testing.T) {
	f := 8.5
	jnb := NewNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	f = 0.0
	jnb = NewNullFloat64(&f)
	if p := jnb.GetFloat64(); p != f {
		t.FailNow()
	}

	jnb = NewNullFloat64(nil)
	if p := jnb.GetFloat64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for NullInt64
func TestNewNullInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewNullInt64(&i)
	if !jnb.Valid || jnb.Int64 != i {
		t.FailNow()
	}

	jnb = NewNullInt64(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt64_Value(t *testing.T) {
	var i int64 = 8
	jnb := NewNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt64(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewNullInt64(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt64_GetInt64(t *testing.T) {
	var i int64 = 8
	jnb := NewNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt64(&i)
	if p := jnb.GetInt64(); p != i {
		t.FailNow()
	}

	jnb = NewNullInt64(nil)
	if p := jnb.GetInt64(); p != 0 {
		t.FailNow()
	}
}

// Test cases for NullInt32
func TestNewNullInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewNullInt32(&i)
	if !jnb.Valid || jnb.Int32 != i {
		t.FailNow()
	}

	jnb = NewNullInt32(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt32_Value(t *testing.T) {
	var i int32 = 8
	jnb := NewNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt32(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewNullInt32(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt32_GetInt32(t *testing.T) {
	var i int32 = 8
	jnb := NewNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt32(&i)
	if p := jnb.GetInt32(); p != i {
		t.FailNow()
	}

	jnb = NewNullInt32(nil)
	if p := jnb.GetInt32(); p != 0 {
		t.FailNow()
	}
}

// Test cases for NullInt16
func TestNewNullInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	i = 0
	jnb = NewNullInt16(&i)
	if !jnb.Valid || jnb.Int16 != i {
		t.FailNow()
	}

	jnb = NewNullInt16(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullInt16_Value(t *testing.T) {
	var i int16 = 8
	jnb := NewNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt16(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	jnb = NewNullInt16(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullInt16_GetInt16(t *testing.T) {
	var i int16 = 8
	jnb := NewNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	i = 0.0
	jnb = NewNullInt16(&i)
	if p := jnb.GetInt16(); p != i {
		t.FailNow()
	}

	jnb = NewNullInt16(nil)
	if p := jnb.GetInt16(); p != 0 {
		t.FailNow()
	}
}

// Test cases for NullByte
func TestNewNullByte(t *testing.T) {
	var i byte = 'r'
	jnb := NewNullByte(&i)
	if !jnb.Valid || jnb.GetByte() != i {
		t.FailNow()
	}

	var nullByte byte
	jnb = NewNullByte(&nullByte)
	if !jnb.Valid || jnb.GetByte() != nullByte {
		t.FailNow()
	}

	jnb = NewNullByte(nil)
	if jnb.Valid {
		t.FailNow()
	}
}

func TestNullByte_Value(t *testing.T) {
	var i byte = 'a'
	jnb := NewNullByte(&i)
	if p := jnb.Value(); p == nil || *p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewNullByte(&nullByte)
	if p := jnb.Value(); p == nil || *p != nullByte {
		t.FailNow()
	}

	jnb = NewNullByte(nil)
	if p := jnb.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullByte_GetByte(t *testing.T) {
	var i byte = 'a'
	jnb := NewNullByte(&i)
	if p := jnb.GetByte(); p != i {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewNullByte(&nullByte)
	if p := jnb.GetByte(); p != nullByte {
		t.FailNow()
	}

	jnb = NewNullByte(nil)
	if p := jnb.GetByte(); p != 0 {
		t.FailNow()
	}
}

func TestNullByte_GetByteString(t *testing.T) {
	var i byte = 'a'
	jnb := NewNullByte(&i)
	// convert to string
	if p := jnb.GetByteString(); p != string(i) {
		t.FailNow()
	}

	var nullByte byte = 0
	jnb = NewNullByte(&nullByte)
	if p := jnb.GetByteString(); p != string(nullByte) {
		t.FailNow()
	}

	jnb = NewNullByte(nil)
	if p := jnb.GetByteString(); p != "" {
		t.FailNow()
	}
}

// Test cases for NullTime
func TestNewNullTime(t *testing.T) {
	now := time.Now()
	jnt := NewNullTime(&now)
	if !jnt.Valid || jnt.GetTime() != now {
		t.FailNow()
	}

	jnt = NewNullTime(nil)
	if jnt.Valid {
		t.FailNow()
	}
}

func TestNullTime_Value(t *testing.T) {
	now := time.Now()
	jnt := NewNullTime(&now)
	if p := jnt.Value(); p == nil || *p != now {
		t.FailNow()
	}

	jnt = NewNullTime(nil)
	if p := jnt.Value(); p != nil {
		t.FailNow()
	}
}

func TestNullTime_GetTime(t *testing.T) {
	now := time.Now()
	jnt := NewNullTime(&now)
	if p := jnt.GetTime(); !p.Equal(now) {
		t.FailNow()
	}

	jnt = NewNullTime(nil)
	if p := jnt.GetTime(); !p.IsZero() {
		t.FailNow()
	}
}
