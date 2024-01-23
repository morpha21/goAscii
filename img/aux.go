package img

import (
	"image"
	"image/color"
)

// //////////////////////////////////////////////////////////// creates a copy of the image, and returns the copy and its dimensions
func copyImage(img image.Image) (*image.RGBA, int, int) {
	copy := image.NewRGBA(img.Bounds())
	x_max := copy.Bounds().Max.X
	y_max := copy.Bounds().Max.Y
	for y := 0; y < y_max; y++ {
		for x := 0; x < x_max; x++ {
			copy.Set(x, y, img.At(x, y))

		}
	}
	return copy, x_max, y_max
}

func grayscalePixel(c color.Color) float32 {
	pixel_color := color.RGBAModel.Convert(c).(color.RGBA)
	r, g, b := pixel_color.R, pixel_color.G, pixel_color.B

	return ((float32(r) + float32(g) + float32(b)) / 3)
}

func pixelMatrix(img *image.Image) [][][3]uint8 {
	x_max := (*img).Bounds().Max.X
	y_max := (*img).Bounds().Max.Y

	matrix := make([][][3]uint8, x_max)

	for x := 0; x < x_max; x++ {
		matrix[x] = make([][3]uint8, y_max)
	}

	for y := 0; y < y_max; y++ {
		for x := 0; x < x_max; x++ {
			col := color.RGBAModel.Convert((*img).At(x-1+i, y-1+j)).(color.RGBA)
			matrix[x][y][0] = col.R
			matrix[x][y][1] = col.G
			matrix[x][y][2] = col.B
		}
	}

	return matrix
}

func shrink(img *image.Image, height int) *image.RGBA {
	matrix := pixelMatrix(img)

}
