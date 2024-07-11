package cmd

type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vector3) Scale(other float64) Vector3 {
	return Vector3{v.X * other, v.Y * other, v.Z * other}
}

func (v Vector3) Lerp(other Vector3, amount float64) Vector3 {
	return Vector3{
		v.X + (other.X-v.X)*amount,
		v.Y + (other.Y-v.Y)*amount,
		v.Z + (other.Z-v.Z)*amount,
	}
}
