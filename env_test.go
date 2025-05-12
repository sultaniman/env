package env

import (
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	firstVar  = "SOME_INT"
	secondVar = "SECRET_INT"
)

const failMessage = `
Wanted: %v
Got: %v
Test case: %s::%s:%d
`

func assertEqual(t *testing.T, expected any, having any) {
	if ok := cmp.Equal(expected, having); !ok {
		_, f, l, _ := runtime.Caller(1)
		t.Errorf(failMessage, expected, having, f, t.Name(), l)
		t.Fail()
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		_, f, l, _ := runtime.Caller(1)
		t.Errorf(failMessage, "error", err, f, t.Name(), l)
		t.Fail()
	}
}

func TestGetInt(t *testing.T) {
	t.Setenv(firstVar, "1")
	t.Setenv(secondVar, "1")

	someInt := GetInt(firstVar)
	assertEqual(t, 1, someInt)

	otherInt := GetInt(secondVar)
	assertEqual(t, 1, otherInt)

	randomInt := GetInt("OTHER_INTS")
	assertEqual(t, 0, randomInt)

	someInt8 := GetInt8(firstVar)
	assertEqual(t, int8(1), someInt8)

	someInt16 := GetInt16(firstVar)
	assertEqual(t, int16(1), someInt16)

	someInt32 := GetInt32(firstVar)
	assertEqual(t, int32(1), someInt32)

	someInt64 := GetInt64(firstVar)
	assertEqual(t, int64(1), someInt64)
}

func TestGetUInt(t *testing.T) {
	t.Setenv(firstVar, "111")
	t.Setenv(secondVar, "101")

	someUInt := GetUInt(firstVar)
	assertEqual(t, uint(111), someUInt)

	otherUInt := GetUInt(secondVar)
	assertEqual(t, uint(101), otherUInt)

	randomInt := GetUInt("OTHER_INTS")
	assertEqual(t, uint(0), randomInt)

	someInt8 := GetUInt8(firstVar)
	assertEqual(t, uint8(111), someInt8)

	someInt16 := GetUInt16(firstVar)
	assertEqual(t, uint16(111), someInt16)

	someInt32 := GetUInt32(firstVar)
	assertEqual(t, uint32(111), someInt32)

	someInt64 := GetUInt64(firstVar)
	assertEqual(t, uint64(111), someInt64)
}

func TestGetFloat(t *testing.T) {
	t.Setenv(firstVar, "111.001")
	t.Setenv(secondVar, "101.0012")

	someFloat32 := GetFloat32(firstVar)
	assertEqual(t, float32(111.001), someFloat32)

	otherFloat64 := GetFloat64(secondVar)
	assertEqual(t, float64(101.0012), otherFloat64)

	randomFloat32 := GetFloat32("RANDOM_FLOAT")
	assertEqual(t, float32(0.0), randomFloat32)

	randomFloat64 := GetFloat64("RANDOM_FLOAT")
	assertEqual(t, float64(0.0), randomFloat64)
}

func TestGetBool(t *testing.T) {
	t.Setenv(firstVar, "true")
	t.Setenv(secondVar, "false")

	firstBool := GetBool(firstVar)
	assertEqual(t, true, firstBool)

	secondBool := GetBool(secondVar)
	assertEqual(t, false, secondBool)

	randomBool := GetBool("random")
	assertEqual(t, false, randomBool)
}

func TestGetBoolE(t *testing.T) {
	t.Setenv(firstVar, "TTrue")
	t.Setenv(secondVar, "FFalse")

	firstBool, err := GetBoolE(firstVar)
	assertEqual(t, false, firstBool)
	assertError(t, err)

	secondBool, err := GetBoolE(secondVar)
	assertEqual(t, false, secondBool)
	assertError(t, err)

}
func TestGetWithPrefix(t *testing.T) {
	SetEnvPrefix("PREFIX_")
	t.Setenv("PREFIX_"+firstVar, "123")
	firstInt := GetInt(firstVar)
	assertEqual(t, 123, firstInt)
}
