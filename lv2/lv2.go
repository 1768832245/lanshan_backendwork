package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(S(3))
}

func S(r float64) float64 {
	result := math.Pi * math.Pow(r, 2)
	return result
}
