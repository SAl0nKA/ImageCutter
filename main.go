package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io/ioutil"
	"path"

	"os"
)

func main() {
	f, _ := os.Open("UNO-Front.png")
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("Error while decoding image:", err)
		return
	}

	var images []image.Image

	maxX := 409
	maxY := 4095 / 7

	startx, starty := 0, 0
	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("New card at", startx, starty)
			upLeft := image.Point{}
			lowRight := image.Point{X: maxX, Y: maxY}
			image := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

			for x := 0; x < maxX; x++ {
				for y := 0; y < maxY; y++ {
					image.Set(x, y, img.At(x+startx, y+starty))
				}
			}
			images = append(images, image)
			startx += maxX
		}
		startx = 0
		starty += maxY
	}

	for i, image := range images {
		fmt.Println("Writting card", fmt.Sprintf("%d.png", i))
		g, _ := os.Create(fmt.Sprintf("%d.png", i))
		png.Encode(g, resize.Resize(82, 117, image, resize.Lanczos3))
		g.Close()
	}

	fmt.Println("done")
}

func ScaleImage(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)

	return img, w, h
}

//resizes images as needed
func resizeExisting() {
	fs, err := ioutil.ReadDir("named")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, image := range fs {
		f, err := os.Open(path.Join("named", image.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Writting image", image.Name())
		img, _ := png.Decode(f)
		g, _ := os.Create(path.Join("named", image.Name()))
		png.Encode(g, resize.Resize(82, 117, img, resize.Lanczos3))
		//resize.Bilinear
		f.Close()
		g.Close()
	}
	fmt.Println("done")
}
