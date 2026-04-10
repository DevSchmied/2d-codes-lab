package datamatrix

import (
	"crypto/sha256"
	"encoding/hex"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	dm "github.com/boombuler/barcode/datamatrix"
)

func GenerateDataMatrixWithBoombuler(payload, outputFile string) error {

	sum := sha256.Sum256([]byte(payload))
	hash := hex.EncodeToString(sum[:])

	code, err := dm.Encode(hash)
	if err != nil {
		log.Println("encode failed:", err)
		return err
	}

	scaled, err := barcode.Scale(code, 300, 300)
	if err != nil {
		log.Println("scale failed:", err)
		return err
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Println("file creation failed:", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("file close failed:", err)
		}
	}()

	if err := png.Encode(file, scaled); err != nil {
		log.Println("png encoding failed:", err)
		return err
	}

	log.Println("Data Matrix successfully created:", outputFile)
	return nil
}
