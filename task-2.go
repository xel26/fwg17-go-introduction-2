package main

import "fmt"

type deretBilangan struct {
	limit int
}



var deret deretBilangan = deretBilangan{
	limit: 40,
}



func (d deretBilangan) ganjil() {
	for i := 1; i <= d.limit; i++ {
		if i % 2 != 0 {
			fmt.Println("bilangan ganjil", i)
		}
	}
}




func (d deretBilangan) genap() {
	for i := 1; i <= d.limit; i++ {
		if i % 2 == 0 {
			fmt.Println("bilangan genap", i)
		}
	}
}




func (d deretBilangan) prima() {
	for i := 1; i <= d.limit; i++ {
		if i % 2 == 0 {
			fmt.Println("bilangan prima", i)
		}
	}
}



func execute() {
	// deret := deretBilangan{40}
	deret.ganjil()
	deret.genap()
	// deret.prima()
	// deret.fibonacci()

}