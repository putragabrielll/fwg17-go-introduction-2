package main

import "fmt"

// import "fmt"

type bilangan struct {
	prima int
	ganjil int
	genap int
	fibonacci int
}

func (temp_1 bilangan) deretanBilangan() int{
	return temp_1.ganjil + temp_1.genap + temp_1.fibonacci * 0
}

func Cetak(angka int){
	var temp_g int
	temp_g = 0
	for i := 0; i <= angka; i++ {
		fmt.Println(i)
		if i % 2 == 0 {
			
		}
	}

	var data bilangan = bilangan{
		prima: 1,
		ganjil: temp_g,
		genap: 3,
		fibonacci: 12,
	}

	fmt.Println(data.deretanBilangan())



	// MAP CONTOH
	// biodata := map[string]string{
	// 	"name": "gabriel",
	// 	"pendidikan": "UNAI",

	// }

	// fmt.Println((biodata["pendidikan"]))
}