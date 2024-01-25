package main

import (
	"fmt"
	"math"
)

// func roundFloat(val float64, precision uint) float64 {
//     ratio := math.Pow(10, float64(precision))
//     return math.Round(val*ratio) / ratio
// }

func Pembulatan(angka float64) {
	
	temp := math.Round(angka * 100) / 100
	fmt.Printf("Result Pembulatan: %.1f", temp)
	// fmt.Println(*&angka)
	// fmt.Println(temp)
}
