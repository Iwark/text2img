package main

import (
	"flag"
	"image/jpeg"
	"os"

	"github.com/Iwark/text2img"
)

var fontPath = flag.String("fontpath", "", "path to the font")
var backgroundImagePath = flag.String("bgimg", "", "path to the background image")
var output = flag.String("output", "image.jpg", "path to the output image")
var text = flag.String("text", "", "text to draw")

func main() {
	flag.Parse()
	d, err := text2img.NewDrawer(text2img.Params{
		FontPath:            *fontPath,
		BackgroundImagePath: *backgroundImagePath,
	})
	if err != nil {
		panic(err.Error())
	}
	img, err := d.Draw(*text)
	if err != nil {
		panic(err.Error())
	}
	file, err := os.Create(*output)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		panic(err.Error())
	}
}
