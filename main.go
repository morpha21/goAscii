package main

import (
	"fmt"
	"goAscii/img"
	"os"
)

func main() {
	fmt.Println("loading image...")
	fig := img.LoadImage(os.Args[1])

	fmt.Println("done.")
}
