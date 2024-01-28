package main

import "fmt"

// struck
type tempData struct{
	sisi float64
}



// method
func (p tempData) luas() float64{
	// fmt.Println(p.sisi * p.sisi)
	return p.sisi * p.sisi
}
func (p tempData) keliling() float64{
	// fmt.Println(12 * p.sisi)
	return 12 * p.sisi
}
func (p tempData) volume() float64{
	// fmt.Println(p.sisi * p.sisi * p.sisi)
	return p.sisi * p.sisi * p.sisi
}




// interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}




// output to main
func CekKubus(x float64){
	var data hitung = tempData{
		sisi: x,
	}

	fmt.Println("Luas Kubus:", data.luas())
	fmt.Println("Keliling Kubus:", data.keliling())
	fmt.Println("Volume Kubus:",data.volume())
}