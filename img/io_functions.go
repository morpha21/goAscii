package img

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
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

func SaveImage(img **image.RGBA, name string) {
	f, err := os.Create("output/" + name)
	checkErr(err, "failed to create file: ")
	defer f.Close()

	checkErr(err, "failed to encode file: ")

	err = png.Encode(f, *img)
}

func checkErr(err error, msg string) {
	if err != nil {
		err = fmt.Errorf(msg+"%w", err)
		log.Fatal(err)
	}
}
