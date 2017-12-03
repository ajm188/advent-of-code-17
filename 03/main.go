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

func turn(direction Direction) Direction {
	return (direction + 1) % 4
}

func move(x, y int, direction Direction) (int, int) {
	switch {
	case direction == UP:
		return x, y + 1
	case direction == DOWN:
		return x, y - 1
	case direction == LEFT:
		return x - 1, y
	case direction == RIGHT:
		return x + 1, y
	}

	return x, y
}

func ringNumber(n int) (ring int) {
	ring = 0
	for {
		oddRing := (2 * ring) + 1
		ringSize := oddRing * oddRing
		if ringSize >= n {
			return
		}
		ring++
	}
	return
}

func ringSize(r int) int {
	odd := (2 * r) - 1
	return odd * odd
}

func ringStart(ring int) int {
	if ring == 0 {
		return 1
	}
	prevRing := ring - 1
	return (2*prevRing+1)*(2*prevRing+1) + 1
}

func location(val, ring int) (x, y int) {
	x, y = ring, (-ring + 1)
	sideLength := (2 * ring) + 1
	i := ringStart(ring)
	amountPerSide := 2
	direction := UP
	for i < val {
		if amountPerSide == sideLength {
			direction = turn(direction)
			amountPerSide = 1
		}

		x, y = move(x, y, direction)
		i++
		amountPerSide++
	}

	return
}

func dist(n int) int {
	if n == 1 {
		return 0
	}
	x, y := location(n, ringNumber(n))

	dist := math.Abs(float64(x)) + math.Abs(float64(y))
	return int(dist)
}

// For ring N, where the origin is in ring 0:
// - The ring begins at (N, (-N + 1))
// - The ring contains (2N + 1) ^ 2 values
// - Each side of the ring has length (2N + 1)
func main() {
	fmt.Println(dist(INPUT))
}
