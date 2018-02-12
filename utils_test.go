package text2img

import (
	"testing"
)

func TestHex(t *testing.T) {
	c, err := Hex("#f00")
	if err != nil {
		t.Error(err.Error())
	}
	if c.R != 255 {
		t.Errorf("R must be 255, got %#v", uint8(c.R))
	}

	c, err = Hex("#00ff08")
	if err != nil {
		t.Error(err.Error())
	}
	if c.G != 255 || c.B != 8 {
		t.Errorf("G and B must be 255, got G: %v, B: %v", uint8(c.G), uint8(c.B))
	}
}
