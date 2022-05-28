package utils

import (
	"testing"
)

func TestGetHash(t *testing.T) {
	data := "test-data"
	expected := 20
	hash := GetHash(data)

	if hash != uint8(expected) {
		t.Errorf("string %s should have hash of %d but found %d", data, expected, hash)
	}
}
