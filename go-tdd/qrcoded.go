package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

// GenerateQRCode generate a QR Code with a phone number
func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}

func main() {
	fmt.Println("Hello QR code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368")
	if err != nil {
		log.Fatal(err)
	}
}
