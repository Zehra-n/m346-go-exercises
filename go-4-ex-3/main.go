package main

// TODO: implement the function computeQuadraticFormula
import (
	"fmt"
	"math"
)

func computeDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	D := computeDiscriminant(a, b, c)

	if D > 0 {
		x1 := (-b - math.Sqrt(D)) / 2 * a
		x2 := (-b + math.Sqrt(D)) / 2 * a
		return []float64{x1, x2}
	} else if D == 0 {
		x := (-b / 2 * a)
		return []float64{x}
	} else {
		return []float64{}
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3.0, 4.0, 1.0)) // [-9 -3]
	fmt.Println(computeQuadraticFormula(2.0, 4.0, 2.0)) // [-4]
	fmt.Println(computeQuadraticFormula(3.0, 4.0, 2.0)) // []
}