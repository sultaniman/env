package env

import (
	"os"
	"strconv"
)

func asInt(v string) (int, error) {
	i, err := strconv.ParseInt(v, 10, getIntBits())
	return int(i), err
}

func asInt8(v string) (int8, error) {
	i, err := strconv.ParseInt(v, 10, 8)
	return int8(i), err
}

func asInt16(v string) (int16, error) {
	i, err := strconv.ParseInt(v, 10, 16)
	return int16(i), err
}

func asInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(v, 10, 32)
	return int32(i), err
}

func asInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func asUInt(v string) (uint, error) {
	i, err := strconv.ParseUint(v, 10, getIntBits())
	return uint(i), err
}

func asUInt8(v string) (uint8, error) {
	i, err := strconv.ParseUint(v, 10, 8)
	return uint8(i), err
}

func asUInt16(v string) (uint16, error) {
	i, err := strconv.ParseUint(v, 10, 16)
	return uint16(i), err
}

func asUInt32(v string) (uint32, error) {
	i, err := strconv.ParseUint(v, 10, 32)
	return uint32(i), err
}

func asUInt64(v string) (uint64, error) {
	i, err := strconv.ParseUint(v, 10, 64)
	return i, err
}

func asFloat32(v string) (float32, error) {
	f, err := strconv.ParseFloat(v, 32)
	return float32(f), err
}

func asFloat64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func getIntBits() int {
	return 32 << (^uint(0) >> 63)
}

func getEnv(name string) string {
	if envPrefix != "" {
		return os.Getenv(envPrefix + name)
	}

	return os.Getenv(name)
}
