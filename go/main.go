package main

import (
	"datamatrix/datamatrix"
	"fmt"
	"log"
)

func main() {
	gasBalloonID := "GB-100045@!$%^&"
	balloonSerialNumber := "SN-2026-000381@!$%^&"

	payload := fmt.Sprintf("%s:%s", gasBalloonID, balloonSerialNumber)
	outputFile := "datamatrix.png"
	if err := datamatrix.GenerateDataMatrixWithBoombuler(payload, outputFile); err != nil {
		log.Println("generation failed:", err)
		return
	}
}
