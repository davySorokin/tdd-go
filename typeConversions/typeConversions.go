package main

import "math"

/*func converting(x, y int, a float64, b uint64) (c, v int, f float64, z uint64) {
	return
}*/

func converting(a, b int) (int, int, float64, uint64) {
	var (
		aa, bb = a, b
		cc     = math.Sqrt(float64(aa*aa + bb*bb))
		dd     = uint64(cc)
	)
	return aa, bb, cc, dd
}

func main() {
	converting(3, 4)
}
