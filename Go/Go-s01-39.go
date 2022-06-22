package main

import "golang.org/x/tour/pic"

func PicOne(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		picture[x] = make([]uint8, dy)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			picture[x][y] = uint8((x + y) / 2)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
