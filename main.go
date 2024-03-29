package main

import (
	"fmt"
	"goAscii/img"
	"os"
)

func main() {
	fmt.Println("loading image...")
	fig := img.LoadImage(os.Args[1])

	processed_fig := img.Shrink(&fig, 140)

	img.SaveImage(&processed_fig, "output.png")

	fmt.Println("done.")
}
