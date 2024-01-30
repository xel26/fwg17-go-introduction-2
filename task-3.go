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

type kubus struct {
	sisi float64
}

func (p kubus) volume() float64 {
	return p.sisi * p.sisi * p.sisi
}

func result() {
	var bangunDatar hitung2d = persegi{6}
	fmt.Println("luas Persegi: ", bangunDatar.luas())
	fmt.Println("keliling Persegi: ", bangunDatar.keliling())

	var bangunRuang hitung3d = kubus{6}
	fmt.Println("volume kubus: ", bangunRuang.volume())
}