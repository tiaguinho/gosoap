package gosoap

import (
	"testing"
)

func TestNewRequestByStruct(t *testing.T) {
	_, err := NewRequestByStruct(nil)
	if err == nil {
		t.Error("An err was expected")
	}
}
