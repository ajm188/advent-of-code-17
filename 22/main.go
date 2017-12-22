package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type NodeState int

const (
	CLEAN NodeState = iota
	WEAKENED
	INFECTED
	FLAGGED
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Point struct {
	X, Y int
}

func (self Point) Up() Point {
	return Point{
		self.X,
		self.Y + 1,
	}
}

func (self Point) Down() Point {
	return Point{
		self.X,
		self.Y - 1,
	}
}

func (self Point) Left() Point {
	return Point{
		self.X - 1,
		self.Y,
	}
}

func (self Point) Right() Point {
	return Point{
		self.X + 1,
		self.Y,
	}
}

type Carrier struct {
	Point
	Direction
}

func NewCarrier() *Carrier {
	return &Carrier{
		Point:     Point{0, 0},
		Direction: UP,
	}
}

func (self *Carrier) Infect(grid map[Point]NodeState) (infection bool) {
	state, ok := grid[self.Point]
	if !ok {
		state = CLEAN
	}
	switch state {
	case CLEAN:
		self.TurnLeft()
		grid[self.Point] = INFECTED
	case INFECTED:
		self.TurnRight()
		grid[self.Point] = CLEAN
	}
	infection = grid[self.Point] == INFECTED
	self.Move()
	return
}

func (self *Carrier) TurnLeft() {
	if self.Direction == UP {
		self.Direction = LEFT
	} else {
		self.Direction = self.Direction - 1
	}
}

func (self *Carrier) TurnRight() {
	self.Direction = (self.Direction + 1) % 4
}

func (self *Carrier) Move() {
	var point Point
	switch self.Direction {
	case UP:
		point = self.Point.Up()
	case DOWN:
		point = self.Point.Down()
	case LEFT:
		point = self.Point.Left()
	case RIGHT:
		point = self.Point.Right()
	}
	self.Point = point
}

func constructGrid(matrix [][]byte) map[Point]NodeState {
	grid := map[Point]NodeState{}
	xOff, yOff := len(matrix[0])/2, len(matrix)/2
	for i := yOff; i >= -yOff; i-- {
		for j := -xOff; j <= xOff; j++ {
			point := Point{X: j, Y: i}
			x, y := j+xOff, (-i)+yOff
			state := CLEAN
			if matrix[y][x] == byte('#') {
				state = INFECTED
			}
			grid[point] = state
		}
	}
	return grid
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	grid := constructGrid(bytes.FieldsFunc(input, func(r rune) bool { return r == '\n' }))
	carrier := NewCarrier()
	infections := 0
	for i := 0; i < 10000; i++ {
		if carrier.Infect(grid) {
			infections++
		}
	}
	fmt.Println(infections)
}
