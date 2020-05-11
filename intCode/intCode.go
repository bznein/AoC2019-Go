package AoC2019Go

type IntCodeCompiler struct {
	Values []int
	Reg    int
}

func (i *IntCodeCompiler) Run() {
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

func (c IntCodeCompiler) Peek(pos int) int {
	//TODO add bound check
	return c.Values[pos]
}

func (c *IntCodeCompiler) Set(pos, val int) {
	//TODO bound checks
	c.Values[pos] = val
}
