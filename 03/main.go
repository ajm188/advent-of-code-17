package main

import (
	"fmt"
	"math"
)

const (
	INPUT = 368078
)

type Direction uint

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type Point struct{ x, y int }
type Grid map[Point]int

func turn(direction Direction) Direction {
	return (direction + 1) % 4
}

func move(p Point, direction Direction) Point {
	switch {
	case direction == UP:
		return Point{p.x, p.y + 1}
	case direction == DOWN:
		return Point{p.x, p.y - 1}
	case direction == LEFT:
		return Point{p.x - 1, p.y}
	case direction == RIGHT:
		return Point{p.x + 1, p.y}
	}

	// this is an error case, but let's make the compiler happy
	return p
}

func initialGrid() Grid {
	return map[Point]int{
		Point{0, 0}: 1,
	}
}

func neighbors(p Point) []Point {
	return []Point{
		Point{p.x, p.y + 1},
		Point{p.x, p.y - 1},
		Point{p.x + 1, p.y},
		Point{p.x - 1, p.y},
	}
}

func plus1(p Point, g Grid) int {
	max := 0
	for _, n := range neighbors(p) {
		v, ok := g[n]
		if ok && v > max {
			max = v
		}
	}

	return max + 1
}

func sideLength(ring int) int {
	return 2*ring + 1
}

func iterate(iterator func(Point, Grid) int, stop func(int) bool) (Point, Grid) {
	p := Point{0, 1}
	direction := UP
	grid := initialGrid()
	ring := 1
	movedPerSide := 2

	val := 1
	for {
		if movedPerSide == sideLength(ring) {
			direction = turn(direction)
			movedPerSide = 1
			if direction == UP {
				ring++
			}
		}

		val = iterator(p, grid)
		if stop(val) {
			break
		}

		grid[p] = val
		p = move(p, direction)
		movedPerSide++
	}
	return p, grid
}

func manhattan(p Point) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func main() {
	p, _ := iterate(plus1, func(v int) bool { return v >= INPUT })
	fmt.Println(manhattan(p))
}
