package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	x, y int
}

func (p *Point) U() {
	p.y--
}
func (p *Point) D() {
	p.y++
}

func (p *Point) L() {
	p.x--
}

func (p *Point) R() {
	p.x++
}

type Grid map[Point]int

func followWire(input string, grid Grid, ip int) {
	p := Point{0, 0}
	for _, s := range strings.Split(input, ",") {
		v, err := strconv.Atoi(s[1:])
		check(err)
		for i := 0; i < v; i++ {
			switch s[0] {
			case 'U':
				p.U()
			case 'D':
				p.D()
			case 'L':
				p.L()
			case 'R':
				p.R()
			default:
				panic("Wrong direction!")
			}
			grid[p] += ip
		}
	}
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Point) Abs() int {
	return Abs(p.x) + Abs(p.y)
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := Grid{}

	i := 1
	for scanner.Scan() {
		v := scanner.Text()
		check(err)
		followWire(v, grid, i)
		i++
	}

	var minK = Point{math.MaxInt32, math.MaxInt32}
	for k, v := range grid {
		if v == 3 {
			val := k.Abs()
			if val < minK.Abs() {
				minK = k
			}

		}
	}
	fmt.Println("Part 1: ", minK.Abs())
}
