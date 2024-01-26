package main

import (
	"fmt"
	"math"
)



func Pembulatan(angka float64) {
	
	temp := math.Round(angka * 100) / 100
	fmt.Printf("Results Pembulatan dari %v menjadi %.1f", angka, temp)
	// fmt.Println(*&angka)
	// fmt.Println(temp)
}
