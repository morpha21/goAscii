package img

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func LoadImage(imagePath string) image.Image {
	imageFile, err := os.Open(imagePath)
	checkErr(err, "failed to open file:")
	defer imageFile.Close()

	i, _, err := image.Decode(imageFile)
	checkErr(err, "failed to decode image:")

	return i
}

func checkErr(err error, msg string) {
	if err != nil {
		err = fmt.Errorf(msg+"%w", err)
		log.Fatal(err)
	}
}
