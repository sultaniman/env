package env

import (
	"os"
	"strconv"
)

// GetInt tries to retrieve and parse int value
// uses `strconv.ParseInt` function, ignores error.
func GetInt(name string) int {
	value, _ := asInt(os.Getenv(name))
	return value
}

// GetInt8 tries to retrieve and parse int8 value
// uses `strconv.ParseInt` function, ignores error.
func GetInt8(name string) int8 {
	value, _ := asInt8(os.Getenv(name))
	return value
}

// GetInt16 tries to retrieve and parse int16 value
// uses `strconv.ParseInt` function, ignores error.
func GetInt16(name string) int16 {
	value, _ := asInt16(os.Getenv(name))
	return value
}

// GetInt32 tries to retrieve and parse int32 value
// uses `strconv.ParseInt` function, ignores error.
func GetInt32(name string) int32 {
	value, _ := asInt32(os.Getenv(name))
	return value
}

// GetInt64 tries to retrieve and parse int64 value
// uses `strconv.ParseInt` function, ignores error.
func GetInt64(name string) int64 {
	value, _ := asInt64(os.Getenv(name))
	return value
}

// GetUInt tries to retrieve and parse uint value
// uses `strconv.ParseUint` function, ignores error.
func GetUInt(name string) uint {
	value, _ := asUInt(os.Getenv(name))
	return value
}

// GetUInt8 tries to retrieve and parse uint8 value
// uses `strconv.ParseUint` function, ignores error.
func GetUInt8(name string) uint8 {
	value, _ := asUInt8(os.Getenv(name))
	return value
}

// GetUInt16 tries to retrieve and parse uint16 value
// uses `strconv.ParseUint` function, ignores error.
func GetUInt16(name string) uint16 {
	value, _ := asUInt16(os.Getenv(name))
	return value
}

// GetUInt32 tries to retrieve and parse uint32 value
// uses `strconv.ParseUint` function, ignores error.
func GetUInt32(name string) uint32 {
	value, _ := asUInt32(os.Getenv(name))
	return value
}

// GetUInt64 tries to retrieve and parse uint64 value
// uses `strconv.ParseUint` function, ignores error.
func GetUInt64(name string) uint64 {
	value, _ := asUInt64(os.Getenv(name))
	return value
}

// GetFloat32 tries to retrieve and parse float32 value
// uses `strconv.ParseFloat` function, ignores error.
func GetFloat32(name string) float32 {
	value, _ := asFloat32(os.Getenv(name))
	return value
}

// GetFloat64 tries to retrieve and parse float64 value
// uses `strconv.ParseFloat` function, ignores error.
func GetFloat64(name string) float64 {
	value, _ := asFloat64(os.Getenv(name))
	return value
}

// GetBool tries to retrieve and parse bool value
// uses `strconv.ParseBool` function, ignores error.
func GetBool(name string) bool {
	value, _ := strconv.ParseBool(os.Getenv(name))
	return value
}

// Same functions with error values

// GetIntE tries to retrieve and parse int value
// returns the result of `strconv.ParseInt`.
func GetIntE(name string) (int, error) {
	return asInt(os.Getenv(name))
}

// GetInt8E tries to retrieve and parse int8 value
// returns the result of `strconv.ParseInt`.
func GetInt8E(name string) (int8, error) {
	return asInt8(os.Getenv(name))
}

// GetInt16E tries to retrieve and parse int16 value
// returns the result of `strconv.ParseInt`.
func GetInt16E(name string) (int16, error) {
	return asInt16(os.Getenv(name))
}

// GetInt32E tries to retrieve and parse int32 value
// returns the result of `strconv.ParseInt`.
func GetInt32E(name string) (int32, error) {
	return asInt32(os.Getenv(name))
}

// GetInt64E tries to retrieve and parse int64 value
// returns the result of `strconv.ParseInt`.
func GetInt64E(name string) (int64, error) {
	return asInt64(os.Getenv(name))
}

// GetUIntE tries to retrieve and parse uint value
// returns the result of `strconv.ParseUint`.
func GetUIntE(name string) (uint, error) {
	return asUInt(os.Getenv(name))
}

// GetUInt8E tries to retrieve and parse uint8 value
// returns the result of `strconv.ParseUint`.
func GetUInt8E(name string) (uint8, error) {
	return asUInt8(os.Getenv(name))
}

// GetUInt16E tries to retrieve and parse uint16 value
// returns the result of `strconv.ParseUint`.
func GetUInt16E(name string) (uint16, error) {
	return asUInt16(os.Getenv(name))
}

// GetUInt32E tries to retrieve and parse uint32 value
// returns the result of `strconv.ParseUint`.
func GetUInt32E(name string) (uint32, error) {
	return asUInt32(os.Getenv(name))
}

// GetUInt64E tries to retrieve and parse uint64 value
// returns the result of `strconv.ParseUint`.
func GetUInt64E(name string) (uint64, error) {
	return asUInt64(os.Getenv(name))
}

// GetFloat32E tries to retrieve and parse float32 value
// returns the result of `strconv.ParseFloat`.
func GetFloat32E(name string) (float32, error) {
	return asFloat32(os.Getenv(name))
}

// GetFloat64E tries to retrieve and parse float64 value
// returns the result of `strconv.ParseFloat`.
func GetFloat64E(name string) (float64, error) {
	return asFloat64(os.Getenv(name))
}
