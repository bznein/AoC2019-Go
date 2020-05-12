package main

import (
	"fmt"
	intcode "github.com/bznein/AoC2019Go/intCode"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	input, err := ioutil.ReadAll(f)

	check(err)

	values := make([]int, 0)

	for _, i := range strings.Split(string(input), ",") {
		j, err := strconv.Atoi(strings.TrimSuffix(i, "\n"))
		check(err)
		values = append(values, j)
	}

	ex := intcode.IntCodeCompiler{Values: append([]int(nil), values...)}
	ex.SetInput(1)
	ex.Run()

	v, _ := ex.Output().Get()

	fmt.Println("Part 1:", v)

	ex = intcode.IntCodeCompiler{Values: append([]int(nil), values...)}
	ex.SetInput(5)
	ex.Run()

	v, _ = ex.Output().Get()

	fmt.Println("Part 2:", v)

}
