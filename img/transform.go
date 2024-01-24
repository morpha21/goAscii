package img

import (
	"image"
	"image/color"
)

func Ascii(fig *image.Image, scale int) {

}

// ------------------------------------------------------// does ordered dithering
func OrderedDithering(img *image.Image) *image.RGBA {

	ordered, x_max, y_max := copyImage(*img)

	threshold_map := [8][8]uint8{
		{0, 32, 8, 40, 2, 34, 10, 42},
		{48, 16, 56, 24, 50, 18, 58, 26},
		{12, 44, 4, 36, 14, 46, 6, 38},
		{60, 28, 52, 20, 62, 30, 54, 22},
		{3, 35, 11, 43, 1, 33, 9, 41},
		{51, 19, 59, 27, 49, 17, 57, 25},
		{15, 47, 7, 39, 13, 45, 5, 37},
		{63, 31, 55, 23, 61, 29, 53, 21}}

	for i := 0; i < len(threshold_map); i++ {
		for j := 0; j < len(threshold_map); j++ {
			threshold_map[i][j] = threshold_map[i][j] * 4 // (aij/64)*max(threshold_map)
		}
	}

	for y := 0; y < y_max; y += len(threshold_map) {
		for x := 0; x < x_max; x += len(threshold_map) {
			for i := 0; i < len(threshold_map) && x+i < x_max; i++ {
				for j := 0; j < len(threshold_map) && y+j < y_max; j++ {
					if grayscalePixel(ordered.At(x+i, y+j)) <= int(threshold_map[i][j]) {
						ordered.Set(x+i, y+j, color.RGBA{28, 28, 2, 255})
						// continue
					} else {
						ordered.Set(x+i, y+j, color.RGBA{233, 234, 164, 255})
						// continue
					}
				}
			}
		}
	}

	return ordered
}
