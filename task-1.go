package main

import "math"

func Pembulatan(N float64) float64 {
	return math.Round(N * 10)/10
}