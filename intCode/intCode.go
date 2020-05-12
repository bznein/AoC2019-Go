package AoC2019Go

import (
	"github.com/markphelps/optional"
	"strconv"
)

type IntCodeCompiler struct {
	Values []int
	Reg    int
	Input  int
	output optional.Int
}

type parameterMode int32

const (
	position  = 0
	immediate = 1
)

func getOpcode(val int) int {
	return val % 100
}

func getMode(val, pos int) parameterMode {
	return parameterMode((val % pos) / (pos / 10))
}

func (i IntCodeCompiler) read(param int, mode parameterMode) int {
	switch mode {
	case position:
		return i.Values[param]
	case immediate:
		return param
	default:
		panic("Wrong mode: " + strconv.Itoa(int(mode)))
	}
}

func (i *IntCodeCompiler) write(param, val int, mode parameterMode) {
	i.Values[param] = val
}

func (i IntCodeCompiler) Output() optional.Int {
	return i.output
}

func (i *IntCodeCompiler) Run() {
	for {
		reg := i.Reg
		instruction := i.Values[reg]
		mode1 := getMode(instruction, 1_000)
		mode2 := getMode(instruction, 10_000)
		mode3 := getMode(instruction, 100_000)
		switch opcode := getOpcode(instruction); opcode {
		case 1:
			v1 := i.read(i.Values[reg+1], mode1)
			v2 := i.read(i.Values[reg+2], mode2)
			i.write(i.Values[reg+3], v1+v2, mode3)
			i.Reg += 4
		case 2:
			v1 := i.read(i.Values[reg+1], mode1)
			v2 := i.read(i.Values[reg+2], mode2)
			i.write(i.Values[reg+3], v1*v2, mode3)
			i.Reg += 4
		case 3:
			i.write(i.Values[reg+1], i.Input, immediate)
			i.Reg += 2
		case 4:
			i.output = optional.NewInt(i.read(i.Values[reg+1], mode1))
			i.Reg += 2
		case 5:
			check := i.read(i.Values[reg+1], mode1)
			val := i.read(i.Values[reg+2], mode2)
			if check != 0 {
				i.Reg = val
			} else {
				i.Reg += 3
			}
		case 6:
			check := i.read(i.Values[reg+1], mode1)
			val := i.read(i.Values[reg+2], mode2)
			if check == 0 {
				i.Reg = val
			} else {
				i.Reg += 3
			}
		case 7:
			v1 := i.read(i.Values[reg+1], mode1)
			v2 := i.read(i.Values[reg+2], mode2)
			var v3 int
			if v1 < v2 {
				v3 = 1
			} else {
				v3 = 0
			}
			i.write(i.Values[reg+3], v3, mode3)
			i.Reg += 4
		case 8:
			v1 := i.read(i.Values[reg+1], mode1)
			v2 := i.read(i.Values[reg+2], mode2)
			var v3 int
			if v1 == v2 {
				v3 = 1
			} else {
				v3 = 0
			}
			i.write(i.Values[reg+3], v3, mode3)
			i.Reg += 4
		case 99:
			return
		default:
			panic("Wrong opcode: " + strconv.Itoa(opcode))
		}

	}
}

func (c *IntCodeCompiler) SetInput(val int) {
	c.Input = val
}

func (c IntCodeCompiler) Peek(pos int) int {
	//TODO add bound check
	return c.Values[pos]
}

func (c *IntCodeCompiler) Set(pos, val int) {
	//TODO bound checks
	c.Values[pos] = val
}
