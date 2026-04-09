package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
)

func main() {
	gasBalloonID := "GB-100045"
	balloonSerialNumber := "SN-2026-000381"

	payload := fmt.Sprintf("%s:%s", gasBalloonID, balloonSerialNumber)
	outputFile := "datamatrix.png"

	code, err := datamatrix.Encode(payload)
	if err != nil {
		log.Println("encode failed:", err)
		return
	}

	scaled, err := barcode.Scale(code, 300, 300)
	if err != nil {
		log.Println("scale failed:", err)
		return
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Println("file creation failed:", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("file close failed:", err)
		}
	}()

	if err := png.Encode(file, scaled); err != nil {
		log.Println("png encoding failed:", err)
		return
	}

	log.Println("Data Matrix successfully created:", outputFile)
}
