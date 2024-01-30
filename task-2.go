package main

import (
	"fmt"
)

type deretBilangan struct {
	limit int
}



func (d deretBilangan) ganjil() (string, int) {
	data := "ganjil: "
	count := 0
	for i := 1; i <= d.limit; i++ {
		if i % 2 != 0 {
			data += fmt.Sprintf("%v ", i)
			count += 1
		}
	}

	return data, count
}




func (d deretBilangan) genap() (string, int) {
	data := "genap: "
	count := 0
	for i := 1; i <= d.limit; i++ {
		if i % 2 == 0 {
			data += fmt.Sprintf("%v ", i)
			count += 1
		}
	}

	return data, count
}




func (d deretBilangan) prima() (string, int) {
	data := "prima: "
	count := 0
	for i := 2; i <= d.limit; i++ {

		faktor := []int{}
		for j := 1; j <= i; j++{
			if i % j == 0{
				faktor = append(faktor, j)
			}
		}

		if len(faktor) == 2{
			data += fmt.Sprintf("%v ", i)
			count += 1
		}

	}

		return data, count
	}




func (d deretBilangan) fibonacci() (string, int) {
	data := make([]int, 10)

	data[0], data[1] = 0, 1

	for i := 2; data[i - 1] + data[i - 2] <= d.limit; i++ {
		data[i] = data[i - 1] + data[i -2]
	}
	 
	results := "fibonacci: "
	count := 0
	for _, v := range data {
		results += fmt.Sprintf("%v ", v)
		count += 1
	}

	return results, count
}



func execute() {
	deret := deretBilangan{40}
	
	ganjil, countGanjil := deret.ganjil()
	fmt.Println(ganjil, "banyak data :", countGanjil)

	genap, countGenap := deret.genap()
	fmt.Println(genap, "banyak data :", countGenap)

	prima, countPrima := deret.prima()
	fmt.Println(prima, "banyak data :", countPrima)
	
	fibonacci, countFibonacci := deret.fibonacci()
	fmt.Println(fibonacci, "banyak data :", countFibonacci)
}