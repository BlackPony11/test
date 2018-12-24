package main

import (
	"fmt"
)

func ReturnOfType(x interface{}) {

	switch x.(type) {
	case bool:
		println("It's bool")
	case int, int16, int32, int64, int8:
		println("It's int")
	case float32, float64:
		println("It's float")
	case string:
		println("It's string")
	case uint, uint16, uint32, uint64, uint8:
		println("It's uint")
	}

}

func nextZn(z, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func Sqrt(x float64) float64 {
	z := nextZn(2, x)
	for zn, delta := z, z; delta > 0.00001; z = zn {
		zn = nextZn(z, x)
		delta = z - zn
	}
	return z
}

func main() {
	ReturnOfType(32)
	ReturnOfType(true)
	fmt.Println(Sqrt(64))

}
