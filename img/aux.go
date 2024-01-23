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

func imageToMatrix(img *image.Image) [][][3]uint8 {
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

func matrixToImage(matrix [][][3]uint8) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, x_max, len(matrix[0])))

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			img.Set(x, y, color.RGBA{matrix[x][y][0], matrix[x][y][1], matrix[x][y][2], 255})
		}
	}
	return img
}

func shrink(img *image.Image, height int) *image.RGBA {
	matrix := imageToMatrix(img)

	x_max := len(matrix)
	y_max := len(matrix[0])

	downscalingFactor := height / (*img).Bounds().Max.Y

	width := (*img).Bounds().Max.X * downscalingFactor

	new_matrix := make([][][3]uint8, width)
	for x := 0; x < width; x++ {
		new_matrix[x] = make([][3]uint8, height)
	}

	mappingMatrix := make([][][2]float64, x_max)
	for x := 0; x < x_max; x++ {
		mappingMatrix[x] = make([][2]float64, y_max)
		for y := 0; y < y_max; y++ {
			mappingMatrix[x][y] = [2]float64{float64(x) * float64(downscalingFactor), float64(y) * float64(downscalingFactor)}
		}
	}

	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {

		}
	}

}
