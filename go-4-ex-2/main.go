package main

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
import (
	"fmt"
	"math"
)

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a float64
	b float64
}

func hypotenuse(shortSides ShortSides) float64 {
	return math.Sqrt(math.Pow(shortSides.a, 2) + math.Pow(shortSides.b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(3.0, 4.0)) // 5
	fmt.Println(computeHypotenuse(4.9, 7.8)) // 9.211405973031479
	fmt.Println(computeHypotenuse(8.3, 4.9)) // 9.638464608017193

	sides1 := ShortSides{
		3.0,
		4.0,
	}
	sides2 := ShortSides{
		4.9,
		7.8,
	}
	sides3 := ShortSides{
		8.3,
		4.9,
	}
	fmt.Println(computeHypotenuse(sides1.a, sides1.b)) // 5
	fmt.Println(computeHypotenuse(sides2.a, sides2.b)) // 9.211405973031479
	fmt.Println(computeHypotenuse(sides3.a, sides3.b)) // 9.638464608017193
}