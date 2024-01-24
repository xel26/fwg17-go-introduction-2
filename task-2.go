package main

import (
	"fmt"
)

type deretBilangan struct {
	limit int
}



func (d deretBilangan) ganjil() string {
	data := "ganjil: "
	for i := 1; i <= d.limit; i++ {
		if i % 2 != 0 {
			data += fmt.Sprintf("%v ", i)
		}
	}

	return data
}




func (d deretBilangan) genap() string {
	data := "genap: "

	for i := 1; i <= d.limit; i++ {
		if i % 2 == 0 {
			data += fmt.Sprintf("%v, ", i)
		}
	}

	return data
}




func (d deretBilangan) prima() string {
	data := "prima: "
	for i := 2; i <= d.limit; i++ {

		faktor := []int{}
		for j := 1; j <= i; j++{
			if i % j == 0{
				faktor = append(faktor, j)
			}
		}

		if len(faktor) == 2{
			data += fmt.Sprintf("%v ,", i)
		}

	}

		return data
	}




func (d deretBilangan) fibonacci() string {
	data := make([]int, 10)
	data[0], data[1] = 0, 1

	for i := 2; data[i - 1] + data[i - 2] <= d.limit; i++ {
		data[i] = data[i - 1] + data[i -2]
	}
	 
	results := "fibonacci: "
	for _, v := range data {
		results += fmt.Sprintf("%v, ", v)
	}

	return results
}



func execute() {
	deret := deretBilangan{40}
	
	fmt.Println(deret.ganjil())
	fmt.Println(deret.genap())
	fmt.Println(deret.prima())
	fmt.Println(deret.fibonacci())

}