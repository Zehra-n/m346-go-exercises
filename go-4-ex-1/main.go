package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) (float32, error) {
	if maxPoints < gotPoints {
		return 0.0, fmt.Errorf("scored more points than possible")
	} else if maxPoints <= 0 || gotPoints < 0 {
		return 0.0, fmt.Errorf("negative points aren't possible")
	}
	return gotPoints/maxPoints*5 + 1, nil
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(34.2, 56.0)) // 4.0535717
	fmt.Println(computeGrade(8.6, 12.0))  // 4.5833335
	fmt.Println(computeGrade(-3.0, 15.0)) // 0 negative points aren't possible
	fmt.Println(computeGrade(21.4, 20.0)) // 0 scored more points than possible
}