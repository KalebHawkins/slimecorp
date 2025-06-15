package engine

type Vector2 struct {
	X, Y float64
}

func (v *Vector2) Add(vec *Vector2) {
	v.X += vec.X
	v.Y += vec.Y
}

func (v *Vector2) Multiply(scaler float64) {
	v.X *= scaler
	v.Y *= scaler
}
