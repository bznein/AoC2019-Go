package main

import (
	"fmt"
	"math"
)

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func main() {
	low := 367479
	high := 893698

	totP1 := 0
	totP2 := 0
	for i := low; i <= high; i++ {
		hasSame := false
		increasing := true
		larger := false
		final := false
		for d := 6; d > 1; d-- {
			d1 := digit(i, d)
			d2 := digit(i, d-1)
			if d1 > d2 {
				increasing = false
				break
			} else if d1 == d2 {
				hasSame = true
				larger = false
				for d-2 >= 1 && d1 == digit(i, d-2) {
					larger = true
					d--
				}
				if !larger {
					final = true
				}
			}
		}
		if hasSame && increasing {
			totP1++
			if final {
				totP2++
			}
		}
	}

	fmt.Println("Part 1: ", totP1)
	fmt.Println("Part 2: ", totP2)
}
