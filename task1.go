package main

import (
	"fmt"
	"math"
)

func Pembulatan(angka float64) {
	fmt.Println("---- Pembulatan! ----")

	temp := math.Round(angka)
	// fmt.Printf("Result Pembulatan: %.1f", temp)
	fmt.Println(*&angka)
	fmt.Println(temp)
}