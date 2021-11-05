package types

import (
	"math"
)

// Float64KeepCeil Keep the first n decimal places, and round up n+1 digit
func Float64KeepCeil(f64 float64, n uint8) float64 {
	k := math.Pow(10, float64(n))
	f64 *= k
	f64 = math.Ceil(f64)
	f64 /= k
	return f64
}

// Float64KeepFloor Keep the first n digits of the decimal, and round down n+1 digit
func Float64KeepFloor(f64 float64, n uint8) float64 {
	k := math.Pow(10, float64(n))
	f64 *= k
	f64 = math.Floor(f64)
	f64 /= k
	return f64
}

// Float64KeepRound Keep the first n digits of the decimal, round the n+1 digit
func Float64KeepRound(f64 float64, n uint8) float64 {
	k := math.Pow(10, float64(n))
	f64 *= k
	f64 = math.Round(f64)
	f64 /= k
	return f64
}
