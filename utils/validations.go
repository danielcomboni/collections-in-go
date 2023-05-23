package utils

import (
	"fmt"
	"strings"
)

func IsNullOrEmpty(i interface{}) bool {

	if strings.TrimSpace(fmt.Sprintf("%v", i)) == "" {
		return true
	}

	if strings.TrimSpace(fmt.Sprintf("%v", i)) == "{}" {
		return true
	}

	// map[string]interface{}{}
	if fmt.Sprintf("%v", i) == "map[]" { // result is map[] and length is 5
		return true
	}

	str := strings.Fields(fmt.Sprintf("%v", i))

	if len(str) == 0 {
		return true
	}

	if len(str) == 2 {
		if str[0] == "{" && str[1] == "}" {
			return true
		}
	}

	return false
}

type ValidationsCheck struct {
	Value interface{} `json:"value"`
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// IsGreaterThan returns true if `value` is greater than `expected`
func IsGreaterThan[T Number](value, expected T) bool {
	return value > expected
}

// IsGreaterThanOrEqualTo returns true if `value` is greater than or equal to `expected`
func IsGreaterThanOrEqualTo[T Number](value, expected T) bool {
	return value >= expected
}

// IsLessThan returns true if `value` is less than `expected`
func IsLessThan[T Number](value, expected T) bool {
	return value < expected
}

// IsLessThanOrEqualTo returns true if `value` is less than or equal to `expected`
func IsLessThanOrEqualTo[T Number](value, expected T) bool {
	return value <= expected
}

// StringContains returns true if `str` contains `containedSubStr`
func StringContains(str, containedSubStr string) bool {
	return strings.Contains(str, containedSubStr)
}
