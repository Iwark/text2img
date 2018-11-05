text2img
===
[![GoDoc](https://godoc.org/gopkg.in/Iwark/spreadsheet.v2?status.svg)](https://godoc.org/github.com/Iwark/text2img)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

package `text2img` lets you generate an image from a text. It's especially useful to generate OG images.

![Example](https://i.imgur.com/3MjL1Pg.jpg)

## Installation

### Binary

If you use `text2img` from your command line, install the binary from [releases](https://github.com/Iwark/text2img/releases).

Or use `go get`:

```
go get github.com/Iwark/text2img/cmd/text2img
```

### Go code

If you use `text2img` from your go code, install the package by `go get`.

```
$ go get github.com/Iwark/text2img
```

## Usage

### Binary

When using a random color background image:

```
$ text2img -fontpath="fonts/font.ttf" -output="test.jpg" -text="text2img generates the image from a text"
```

Using a specific image file:

```
$ text2img -fontpath="fonts/font.ttf" -output="test.jpg" -text="text2img generates the image from a text" -bgimg="gophers.jpg"
```

![bgimgEx](https://i.imgur.com/MWNV44f.jpg)

([The Go gopher](https://blog.golang.org/gopher) was designed by [Renée French.](http://reneefrench.blogspot.com/))

### Go code

You can use this package as follows:

```go
package main

import (
  "image/jpeg"
  "os"

  "github.com/Iwark/text2img"
)

func main() {
  path := "fonts/font.ttf"
  d, err := text2img.NewDrawer(text2img.Params{
    FontPath: path,
  })
  checkError(err)

  img, err := d.Draw("text2img generates the image from a text")
  checkError(err)

  file, err := os.Create("test.jpg")
  checkError(err)
  defer file.Close()

  err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
  checkError(err)
}
```

More usage can be found at the [godoc](https://godoc.org/github.com/Iwark/text2img).

## License

This package is released under the MIT License.
