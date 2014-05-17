package filepreviews

import (
	"testing"
)

func TestNew(t *testing.T) {
	fp := New()
	if fp == nil {
		t.Error("New() should not return nil")
	}
}

func TestGenerate(t *testing.T) {
	fp := New()
	opts := &FilePreviewsOptions{}
	_, err := fp.Generate("http://www.getblimp.com/images/screenshot1.png", opts)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestGenerateWithMetadata(t *testing.T) {
	fp := New()
	opts := &FilePreviewsOptions{
		Size: map[string]int{
			"width":  50,
			"height": 100,
		},
		Metadata: []string{"all"},
	}
	_, err := fp.Generate("http://www.getblimp.com/images/screenshot1.png", opts)
	if err != nil {
		t.Errorf("%v", err)
	}
}
