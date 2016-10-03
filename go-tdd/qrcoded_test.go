package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

// Second Test
func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368")

	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}

	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	e := GenerateQRCode(w, "555-2368")

	if e == nil || e.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, go %v", e)
	}
}
