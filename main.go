package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	imgFileName := os.Args[1]
	pxlType, _ := strconv.Atoi(os.Args[2])

	imgFile, err := os.Open(imgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	x0, y0 := img.Bounds().Min.X, img.Bounds().Min.Y
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	koffX, koffY := width/200+1, height/50

	fmt.Println(koffX, "*", koffY)

	levels := []string{" ", ".", ",", ";", "!", "v", "l", "L", "F", "E", "$"}
	if pxlType == 1 {
		levels = []string{" ", ".", ",", ";", "!", "v", "l", "L", "F", "E", "$"}
	} else if pxlType == 2 {
		levels = []string{" ", "░", "▒", "▓", "█"}
	}

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
