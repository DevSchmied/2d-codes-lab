package main

import (
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
)

func main() {
	gasBalloonID := "GB-100045"
	balloonSerialNumber := "SN-2026-000381"

	combinedData := gasBalloonID + ":" + balloonSerialNumber

	// Encode the text into a Data Matrix code.
	code, err := datamatrix.Encode(combinedData)
	if err != nil {
		log.Println("encode failed:", err)
		return
	}

	// Scale the code to a readable image size.
	scaled, err := barcode.Scale(code, 300, 300)
	if err != nil {
		log.Println("scale failed:", err)
		return
	}

	// Create the output PNG file.
	file, err := os.Create("datamatrix.png")
	if err != nil {
		log.Println("file creation failed:", err)
		return
	}
	// Ensure the file is closed before exiting.
	defer file.Close()

	// Write the scaled image as PNG.
	if err := png.Encode(file, scaled); err != nil {
		log.Println("png encoding failed:", err)
		return
	}

	// Report success.
	log.Println("Data Matrix successfully created: datamatrix.png")
}
