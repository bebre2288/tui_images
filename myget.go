package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	fullUrlFile string
	fileName    string = "temp.png"
	tes         image.Image
)

func main() {

	//https://pngimg.com/uploads/parrot/parrot_PNG96579.png
	//
	fullUrlFile = os.Args[1]
	for x := 1; x < 50; x++ {
		doIt("https://pngimg.com/uploads/starfish/starfish_PNG" + strconv.Itoa(x) + ".png")
	}

}

func doIt(url string) {
	resp, err := httpClient().Get(url)
	if err != nil {
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	file, err := os.Create(fileName)
	fmt.Println(err)
	//size, err := io.Copy(file, resp.Body)
	//size, err := io.Copy(tes, resp.Body)
	fmt.Println(size)
	printAscii()
	os.Remove(fileName)
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func printAscii() {
	imgFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
	img, err := png.Decode(imgFile)

	x0, y0 := img.Bounds().Min.X, img.Bounds().Min.Y
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	koffX, koffY := width/200+1, height/50

	levels := []string{" ", "░", "▒", "▓", "█"}

	length := len(levels)
	for y := y0; y < height-koffY; y += koffY {
		for x := x0; x < width-koffX; x += koffX {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / uint8((255 / length))
			if level == uint8(length) {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}
