package main

import "fmt"

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

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

type balok struct {
	panjang float64
	lebar float64
	tinggi float64
}

func (p balok) volume() float64 {
	return p.panjang * p.lebar * p.tinggi
}


type CombinedStruct struct {
	persegi
	balok
}

func result() {
	var persegi hitung = CombinedStruct{persegi: persegi{sisi: 2}}
	var kubus hitung = CombinedStruct{balok: balok{panjang: 3, lebar: 2, tinggi: 1}}

	fmt.Printf("keliling persegi %v ", persegi.keliling())
	fmt.Printf("luas persegi %v ", persegi.luas())
	fmt.Printf("volume balok %v ", kubus.volume())
}