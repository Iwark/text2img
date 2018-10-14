package text2img

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// Drawer is the main interface for this package
type Drawer interface {
	Draw(string) (*image.RGBA, error)
	SetColors(color.RGBA, color.RGBA)
	SetFontPath(string) error
	SetFontSize(float64)
	SetSize(int, int)
}

// Params is parameters for NewDrawer function
type Params struct {
	Width               int
	Height              int
	FontPath            string
	BackgroundImagePath string
	FontSize            float64
	BackgroundColor     color.RGBA
	TextColor           color.RGBA
}

// NewDrawer returns Drawer interface
func NewDrawer(params Params) (Drawer, error) {
	d := &drawer{}
	if params.FontPath != "" {
		err := d.SetFontPath(params.FontPath)
		if err != nil {
			return d, err
		}
	}
	if params.BackgroundImagePath != "" {
		err := d.SetBackgroundImage(params.BackgroundImagePath)
		if err != nil {
			return d, err
		}
		d.SetSize(d.BackgroundImage.Bounds().Size().X, d.BackgroundImage.Bounds().Size().Y)
	} else {
		d.SetSize(params.Width, params.Height)
	}

	d.SetColors(params.TextColor, params.BackgroundColor)
	d.SetFontSize(params.FontSize)

	return d, nil
}

type drawer struct {
	BackgroundColor *image.Uniform
	BackgroundImage image.Image
	Font            *truetype.Font
	FontSize        float64
	Height          int
	TextColor       *image.Uniform
	Width           int

	autoFontSize bool
}

// Draw returns the image of a text
func (d *drawer) Draw(text string) (img *image.RGBA, err error) {
	if d.BackgroundImage != nil {
		imgRect := image.Rectangle{image.Pt(0, 0), d.BackgroundImage.Bounds().Size()}
		img = image.NewRGBA(imgRect)
		draw.Draw(img, img.Bounds(), d.BackgroundImage, image.ZP, draw.Src)
	} else {
		img = image.NewRGBA(image.Rect(0, 0, d.Width, d.Height))
		draw.Draw(img, img.Bounds(), d.BackgroundColor, image.ZP, draw.Src)
	}
	if d.autoFontSize {
		d.FontSize = d.calcFontSize(text)
	}
	textWidth := d.calcTextWidth(d.FontSize, text)

	if d.Font != nil {
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(d.Font)
		c.SetFontSize(d.FontSize)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(d.TextColor)
		c.SetHinting(font.HintingNone)

		textHeight := int(c.PointToFixed(d.FontSize) >> 6)
		pt := freetype.Pt((d.Width-textWidth)/2, (d.Height+textHeight)/2)
		_, err = c.DrawString(text, pt)
		return
	}
	err = errors.New("Font must be specified")
	// point := fixed.Point26_6{640, 960}
	// fd := &font.Drawer{
	// 	Dst:  img,
	// 	Src:  d.TextColor,
	// 	Face: basicfont.Face7x13,
	// 	Dot:  point,
	// }
	// fd.DrawString(text)
	return
}

// SetBackgroundImage sets the specific background image
func (d *drawer) SetBackgroundImage(imagePath string) (err error) {
	src, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return
	}
	d.BackgroundImage = img
	return
}

// SetColors sets the textColor and the backgroundColor
func (d *drawer) SetColors(textColor, backgroundColor color.RGBA) {
	r1, g1, b1, a1 := backgroundColor.RGBA()
	r2, g2, b2, a2 := textColor.RGBA()
	if r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2 {
		color := PickColor()
		d.TextColor = image.NewUniform(color.TextColor)
		d.BackgroundColor = image.NewUniform(color.BackgroundColor)
		return
	}
	d.TextColor = image.NewUniform(textColor)
	d.BackgroundColor = image.NewUniform(backgroundColor)
}

// SetColors sets the font
func (d *drawer) SetFontPath(fontPath string) (err error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return
	}
	d.Font = f
	return
}

// SetColors sets the fontSize
func (d *drawer) SetFontSize(fontSize float64) {
	if fontSize > 0 {
		d.autoFontSize = false
		d.FontSize = fontSize
		return
	}
	d.autoFontSize = true
}

// SetColors sets the size
func (d *drawer) SetSize(width, height int) {
	if width <= 0 {
		d.Width = 1200
	} else {
		d.Width = width
	}
	if height <= 0 {
		d.Height = 630
	} else {
		d.Height = height
	}
}

func (d *drawer) calcFontSize(text string) (fontSize float64) {
	const padding = 4
	fontSizes := []float64{128, 64, 48, 32, 24, 18, 16, 14, 12}
	for _, fontSize = range fontSizes {
		textWidth := d.calcTextWidth(fontSize, text)
		if textWidth < d.Width {
			return
		}
	}
	return
}

func (d *drawer) calcTextWidth(fontSize float64, text string) (textWidth int) {
	var face font.Face
	if d.Font != nil {
		opts := truetype.Options{}
		opts.Size = fontSize
		face = truetype.NewFace(d.Font, &opts)
	} else {
		face = basicfont.Face7x13
	}
	for _, x := range text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			return
		}
		textWidth += int(float64(awidth) / 64)
	}
	return
}
