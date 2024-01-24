package img

import (
	"image"
	"image/color"
	"math"
)

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

func grayscalePixel(c color.Color) int {
	pixel_color := color.RGBAModel.Convert(c).(color.RGBA)
	r, g, b := pixel_color.R, pixel_color.G, pixel_color.B

	return int(math.Round(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)))
}

func avgPixel(img *image.Image, x, y, w, h int) color.RGBA {
	cnt := 0.0

	meanR, meanG, meanB := 0.0, 0.0, 0.0

	max := (*img).Bounds().Max

	for i := x; i <= x+w && i < max.X; i++ {
		for j := y; j <= y+w && j < max.Y; j++ {
			r, g, b, _ := (*img).At(i, j).RGBA()
			meanR += float64(r)
			meanG += float64(g)
			meanB += float64(b)
			cnt++
		}
	}

	meanR = meanR / cnt * (255 / 4294967295)
	meanG = meanG / cnt * (255 / 4294967295)
	meanB = meanB / cnt * (255 / 4294967295)

	return color.RGBA{uint8(meanR), uint8(meanG), uint8(meanB), 255}
}

func Shrink(img *image.Image, new_height int) *image.RGBA {

	max := (*img).Bounds().Max

	downscalingFactor := float64(new_height) / float64(max.Y)

	new_width := int(math.Round(downscalingFactor * float64(max.X)))

	new_img := image.NewRGBA(image.Rect(0, 0, new_width, new_height))

	h := int(math.Round(float64((*img).Bounds().Max.X)/float64(new_width))) - 1
	w := int(math.Round(1/downscalingFactor)) - 1

	xaux := 0
	yaux := 0

	for x := 0; x < max.X; x += w {
		for y := 0; y < max.Y; y += h {
			new_img.Set(xaux, yaux, avgPixel(img, x, y, w, h))
			yaux++
		}
		xaux++
	}

	return new_img
}
