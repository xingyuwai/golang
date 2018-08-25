package main

import "image"

func main() {
	m := MyImage{image.Point{3, 4}, image.Point{303, 204}}
	//Show(Pic)
	ShowImage(m)
}
