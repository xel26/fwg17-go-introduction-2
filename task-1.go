package main

import "math"

func Pembulatan(N float64) float64 {
	persepuluh := N * 10											// 1.  memisahkan angka persepuluh dengan angka di belakangnya. 4.37 => 43.7
	round := math.Round(persepuluh)									// 2. membulatkan angka di belakang persepuluh. 43.7 => 44
	presisi := (round * 10)/100										// 3. mengembalikan lagi menjadi angka desimal. 44 => 440 => 4.4
	return presisi													// 4. menampilkan dengan format dua angka di belakang koma. 4.4 => 4.40
}