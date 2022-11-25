package structs

type Rectangle struct {
	Width  float64
	Length float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Length)
}

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Length
}
