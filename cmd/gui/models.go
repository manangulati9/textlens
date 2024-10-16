package gui

type Rect struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

// Width of a rect.
func (r *Rect) Width() int {
	return r.Right - r.Left + 1
}

// Height of rect.
func (r *Rect) Height() int {
	return r.Bottom - r.Top + 1
}

// Expose rect for usage with QT as {left, top, width, height}.
func (r *Rect) Geometry() [4]int {
	return [4]int{r.Left, r.Top, r.Width(), r.Height()}
}

// Width and Height or a rect.
func (r *Rect) Size() [2]int {
	return [2]int{r.Width(), r.Height()}
}

// Create an integer-scaled copy of the Rect.
func (r *Rect) Scale(factor float32) *Rect {
	return &Rect{
		int(float32(r.Left) * factor),
		int(float32(r.Top) * factor),
		int(float32(r.Right) * factor),
		int(float32(r.Bottom) * factor),
	}
}
