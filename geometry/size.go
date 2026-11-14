package geometry

type Size struct {
	Width, Height float64
}

func (s Size) Scale(factor float64) Size {
	return Size{Width: s.Width * factor, Height: s.Height * factor}
}

func (s Size) Radius() float64 {
	return s.Width / 2
}
