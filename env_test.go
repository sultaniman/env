package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	firstVar  = "SOME_INT"
	secondVar = "SECRET_INT"
)

func TestGetInt(t *testing.T) {
	t.Setenv(firstVar, "1")
	t.Setenv(secondVar, "1")

	someInt := GetInt(firstVar)
	assert.Equal(t, 1, someInt)

	otherInt := GetInt(secondVar)
	assert.Equal(t, 1, otherInt)

	randomInt := GetInt("OTHER_INTS")
	assert.Equal(t, 0, randomInt)

	someInt8 := GetInt8(firstVar)
	assert.IsType(t, int8(0), someInt8)
	assert.Equal(t, int8(1), someInt8)

	someInt16 := GetInt16(firstVar)
	assert.IsType(t, int16(0), someInt16)
	assert.Equal(t, int16(1), someInt16)

	someInt32 := GetInt32(firstVar)
	assert.IsType(t, int32(0), someInt32)
	assert.Equal(t, int32(1), someInt32)

	someInt64 := GetInt64(firstVar)
	assert.IsType(t, int64(0), someInt64)
	assert.Equal(t, int64(1), someInt64)
}

func TestGetUInt(t *testing.T) {
	t.Setenv(firstVar, "111")
	t.Setenv(secondVar, "101")

	someUInt := GetUInt(firstVar)
	assert.Equal(t, uint(111), someUInt)

	otherUInt := GetUInt(secondVar)
	assert.Equal(t, uint(101), otherUInt)

	randomInt := GetUInt("OTHER_INTS")
	assert.Equal(t, uint(0), randomInt)

	someInt8 := GetUInt8(firstVar)
	assert.IsType(t, uint8(0), someInt8)
	assert.Equal(t, uint8(111), someInt8)

	someInt16 := GetUInt16(firstVar)
	assert.IsType(t, uint16(0), someInt16)
	assert.Equal(t, uint16(111), someInt16)

	someInt32 := GetUInt32(firstVar)
	assert.IsType(t, uint32(0), someInt32)
	assert.Equal(t, uint32(111), someInt32)

	someInt64 := GetUInt64(firstVar)
	assert.IsType(t, uint64(0), someInt64)
	assert.Equal(t, uint64(111), someInt64)
}

func TestGetFloat(t *testing.T) {
	t.Setenv(firstVar, "111.001")
	t.Setenv(secondVar, "101.0012")

	someFloat32 := GetFloat32(firstVar)
	assert.Equal(t, float32(111.001), someFloat32)

	otherFloat64 := GetFloat64(secondVar)
	assert.Equal(t, float64(101.0012), otherFloat64)

	randomFloat32 := GetFloat32("RANDOM_FLOAT")
	assert.Equal(t, float32(0.0), randomFloat32)

	randomFloat64 := GetFloat64("RANDOM_FLOAT")
	assert.Equal(t, float64(0.0), randomFloat64)
}

func TestGetBool(t *testing.T) {
	t.Setenv(firstVar, "true")
	t.Setenv(secondVar, "false")

	firstBool := GetBool(firstVar)
	assert.Equal(t, true, firstBool)

	secondBool := GetBool(secondVar)
	assert.Equal(t, false, secondBool)

	randomBool := GetBool("random")
	assert.Equal(t, false, randomBool)
}

func TestGetWithPrefix(t *testing.T) {
	SetEnvPrefix("PREFIX_")
	t.Setenv("PREFIX_"+firstVar, "123")
	firstInt := GetInt(firstVar)
	assert.Equal(t, 123, firstInt)
}
