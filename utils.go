package text2img

import (
	"fmt"
	"image/color"
)

func Hex(scol string) (color.RGBA, error) {
	format := "#%02x%02x%02x"
	factor := uint8(1)
	if len(scol) == 4 {
		format = "#%1x%1x%1x"
		factor = uint8(17)
	}

	var r, g, b uint8
	n, err := fmt.Sscanf(scol, format, &r, &g, &b)
	if err != nil {
		return color.RGBA{}, err
	}
	if n != 3 {
		return color.RGBA{}, fmt.Errorf("color: %v is not a hex-color", scol)
	}
	return color.RGBA{r * factor, g * factor, b * factor, 255}, nil
}

func must(c color.RGBA, e error) color.RGBA {
	return c
}
