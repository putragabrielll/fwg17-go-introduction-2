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
	return temp_1.genap
}

func Cetak(angka int){
	var temp_genap int
	var temp_ganjil int
	temp_genap = 0
	temp_ganjil = 0
	for i := 0; i <= angka; i++ {
		if i % 2 == 0 {
			// fmt.Println(i)
			temp_genap+=i
		} else if i % 2 == 1 {
			// fmt.Println(i)
			temp_ganjil+=i
		}
	}

	var data bilangan = bilangan{
		prima: 1,
		ganjil: temp_ganjil,
		genap: temp_genap,
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