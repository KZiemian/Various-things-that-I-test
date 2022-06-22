func (i Image) ColorModel() color.Model {
	return Color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	image.Rect(0, 0, 20, 20)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{(x+y)/2, (x+y)/2, 255, 255}
}
