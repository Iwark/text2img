package text2img

import (
	"image/color"
	"math/rand"
	"time"
)

// Color contains a good conbination of backgroundColor and textColor
type Color struct {
	BackgroundColor color.RGBA
	TextColor       color.RGBA
}

var colors []Color

func init() {
	rand.Seed(time.Now().UnixNano())
	g333 := must(Hex("#333"))
	fff := must(Hex("#fff"))
	colors = []Color{
		Color{must(Hex("#003d47")), fff},
		Color{must(Hex("#128277")), fff},
		Color{must(Hex("#d24136")), fff},
		Color{must(Hex("#eb8a3e")), fff},
		Color{must(Hex("#ebb582")), fff},
		Color{must(Hex("#785a46")), fff},
		Color{must(Hex("#bc6d4f")), fff},
		Color{must(Hex("#1e1f26")), fff},
		Color{must(Hex("#283655")), fff},
		Color{must(Hex("#4d648d")), fff},
		Color{must(Hex("#265c00")), fff},
		Color{must(Hex("#faaf08")), fff},
		Color{must(Hex("#fa812f")), fff},
		Color{must(Hex("#fa4032")), fff},
		Color{must(Hex("#6c5f5b")), fff},
		Color{must(Hex("#cdab81")), fff},
		Color{must(Hex("#4f4a45")), fff},
		Color{must(Hex("#04202c")), fff},
		Color{must(Hex("#304040")), fff},
		Color{must(Hex("#5b7065")), fff},
		Color{must(Hex("#1e0000")), fff},
		Color{must(Hex("#500805")), fff},
		Color{must(Hex("#9d331f")), fff},
		Color{must(Hex("#68a225")), fff},
		Color{must(Hex("#fdffff")), g333},
		Color{must(Hex("#2c4a52")), fff},
		Color{must(Hex("#537072")), fff},
		Color{must(Hex("#8e9b97")), fff},
		Color{must(Hex("#f4ebdb")), g333},
		Color{must(Hex("#d8412f")), fff},
		Color{must(Hex("#fe7a47")), fff},
		Color{must(Hex("#fcfdfe")), g333},
		Color{must(Hex("#867666")), fff},
		Color{must(Hex("#e1b80d")), fff},
		Color{must(Hex("#003b46")), fff},
		Color{must(Hex("#07575b")), fff},
		Color{must(Hex("#66a5ad")), fff},
		Color{must(Hex("#af6c59")), fff},
		Color{must(Hex("#e68f71")), fff},
		Color{must(Hex("#021c1e")), fff},
		Color{must(Hex("#004445")), fff},
		Color{must(Hex("#2c7873")), fff},
		Color{must(Hex("#6fb98f")), fff},
		Color{must(Hex("#434343")), fff},
		Color{must(Hex("#767676")), fff},
		Color{must(Hex("#c16707")), fff},
		Color{must(Hex("#f08d16")), fff},
		Color{must(Hex("#77262a")), fff},
		Color{must(Hex("#9e2d29")), fff},
		Color{must(Hex("#c35d44")), fff},
		Color{must(Hex("#202d35")), fff},
		Color{must(Hex("#0e3c54")), fff},
		Color{must(Hex("#2a677c")), fff},
		Color{must(Hex("#4f3538")), fff},
		Color{must(Hex("#66443b")), fff},
		Color{must(Hex("#c29f83")), fff},
		Color{must(Hex("#210e3b")), fff},
		Color{must(Hex("#4b194c")), fff},
		Color{must(Hex("#872b76")), fff},
	}
}

// PickColor picks a color
func PickColor() Color {
	return colors[rand.Intn(len(colors))]
}
