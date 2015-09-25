package filepreviews

import (
	"testing"
)

func TestNew(t *testing.T) {
	fp := New("1", "2")
	if fp == nil {
		t.Error("New() should not return nil")
	}
}
