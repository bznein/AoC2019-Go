package main

import (
	"fmt"
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

	ex := IntCodeCompiler{Values: values}
	ex.Set(1, 12)
	ex.run()
	fmt.Println("Part 1:", ex.peek(0))
}
