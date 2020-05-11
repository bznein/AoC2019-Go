package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func massPart2(i int) int {
	f := i/3 - 2
	if f <= 0 {
		return f
	}
	return f + massPart2(f)
}

func massPart1(i int) int {
	return i/3 + 2
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	tot1 := 0
	tot2 := 0
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		check(err)
		tot1 += massPart1(v)
		tot2 += massPart2(v)
	}

	fmt.Println("Part 1: ", tot1)
	fmt.Println("Part 2: ", tot2)

}
