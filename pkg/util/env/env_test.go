// +build identity_test, payment_test

package env

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	EnvSet   string = "ENV_SET_VAR"
	EnvUnset string = "ENV_UNSET_VAR"
)

func TestGetEnvAsString(t *testing.T) {
	const expected = "foo"

	assert := assert.New(t)

	os.Setenv(EnvSet, expected)
	assert.Equal(expected, GetEnvAsString(EnvSet, expected))

	assert.Equal(expected, GetEnvAsString(EnvUnset, expected))
}

func TestGetEnvAsIntOrFallback(t *testing.T) {
	const expected = 1

	assert := assert.New(t)

	os.Setenv(EnvSet, strconv.Itoa(expected))
	returnVal, _ := GetEnvAsInt(EnvSet, expected)
	assert.Equal(expected, returnVal)

	returnVal, _ = GetEnvAsInt(EnvUnset, expected)
	assert.Equal(expected, returnVal)

	os.Setenv(EnvSet, "not-an-int")
	returnVal, err := GetEnvAsInt(EnvSet, expected)
	assert.Equal(expected, returnVal)
	if err == nil {
		t.Error("expected error")
	}
}

func TestGetEnvAsFloat64OrFallback(t *testing.T) {
	const expected = 1.0

	assert := assert.New(t)

	os.Setenv(EnvSet, "1.0")
	returnVal, _ := GetEnvAsFloat64(EnvSet, expected)
	assert.Equal(expected, returnVal)

	returnVal, _ = GetEnvAsFloat64(EnvUnset, expected)
	assert.Equal(expected, returnVal)

	os.Setenv(EnvSet, "not-a-flaat")
	returnVal, err := GetEnvAsFloat64(EnvSet, 1.0)
	assert.Equal(expected, returnVal)
	if err == nil {
		t.Error("expected error")
	}
}
