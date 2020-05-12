package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

type Grid struct {
	m   map[Point]int
	mux sync.Mutex
}

func (g *Grid) Inc(p Point, ip int) {
	g.mux.Lock()
	defer g.mux.Unlock()
	g.m[p] += ip
}

func followWire(input string, grid *Grid, ip int, wg *sync.WaitGroup, inc bool) {
	defer wg.Done()
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
			grid.Inc(p, ip)
			if inc {
				ip++
			}
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
	grid := Grid{m: map[Point]int{}}

	var gridp2 [2]Grid
	gridp2[0] = Grid{m: map[Point]int{}}
	gridp2[1] = Grid{m: map[Point]int{}}

	var wg sync.WaitGroup
	i := 1
	for scanner.Scan() {
		v := scanner.Text()
		check(err)
		wg.Add(2)
		go followWire(v, &grid, i, &wg, false)
		go followWire(v, &gridp2[i-1], 1, &wg, true)
		i++
	}

	wg.Wait()

	var minK = Point{math.MaxInt32, math.MaxInt32}
	for k, v := range grid.m {
		if v == 3 {
			val := k.Abs()
			if val < minK.Abs() {
				minK = k
			}

		}
	}
	fmt.Println("Part 1: ", minK.Abs())

	minV := math.MaxInt32
	for k, v := range gridp2[0].m {
		v2, found := gridp2[1].m[k]
		if found && v+v2 < minV {
			minV = v + v2
		}
	}
	fmt.Println("Part 2: ", minV)

}
