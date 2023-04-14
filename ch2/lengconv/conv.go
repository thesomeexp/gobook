package lengconv

// 假设 1m = 3.28 ft
func MToF(m Metre) Feet {
	return Feet(m * 3.28)
}
func FToM(f Feet) Metre {
	return Metre(f / 3.28)
}
