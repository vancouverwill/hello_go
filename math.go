package hello_go

import (
	"fmt"
	"math"
)

func sum(x, y int) int {
	return x + y
}

func Sqrt(x float64) float64 {

	var z float64 = x
	for i := 0; i < 10; i++ {
		fmt.Println(i, z)
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

//func Sqrt(x float64) float64 {

//	var z float64 = x
//	for i := 0; i < 10; i++ {
//		fmt.Println(i, z)
//		z = z - ((math.Pow(x, 2) - x) / (2 * z))
//	}
//	return z
//}

func exampleMaths() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Pow(2, 3))

	fmt.Println(Sqrt(3))
	//	fmt.Println(Sqrt(9))

	fmt.Println(sum(5, 14))
}
