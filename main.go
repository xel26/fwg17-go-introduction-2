package main

import "fmt"




func main() {
	// task 1
	params := [5]float64{4.37, 4.32, 4.324, 4.58, 4.1889}

	for _, v := range params{
		result := Pembulatan(v)
		fmt.Printf("%.2f\n", result)
	}



	//task 2
	execute()
}