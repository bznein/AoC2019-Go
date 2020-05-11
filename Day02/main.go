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

type IntCodeCompiler struct {
	Values []int
	Reg    int
}

func (i *IntCodeCompiler) run() {
	for {
		reg := i.Reg
		switch opcode := i.Values[reg]; opcode {
		case 1:
			i.Values[i.Values[reg+3]] = i.Values[i.Values[reg+1]] + i.Values[i.Values[reg+2]]
		case 2:
			i.Values[i.Values[reg+3]] = i.Values[i.Values[reg+1]] * i.Values[i.Values[reg+2]]
		case 99:
			return
		default:
			panic("Wrong opcode!")
		}
		i.Reg += 4
	}
}

func (c IntCodeCompiler) peek(pos int) int {
	//TODO add bound check
	return c.Values[pos]
}

func (c *IntCodeCompiler) Set(pos, val int) {
	//TODO bound checks
	c.Values[pos] = val
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
	ex.Set(1, 12)
	ex.Run()
	fmt.Println("Part 1:", ex.Peek(0))

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			ex = intcode.IntCodeCompiler{Values: append([]int(nil), values...)}
			ex.Set(1, i)
			ex.Set(2, j)
			ex.Run()
			if ex.Peek(0) == 19690720 {
				fmt.Println("Part 2: ", 100*i+j)
				return
			}
		}
	}
}
