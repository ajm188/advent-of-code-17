package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Consider the following structure:
//   +--+
//  /n   \
// +    h +
//  \p   /
//   +--+
// We now have a 3-axis systetm, denoted h for horizontal, p for the positive
// (think of the line defined by y = x in a Cartesian system) and n for the
// negative (similarly y = -x in a Cartesian grid).

type Point struct {
	h, p, n int
}

func OriginPoint() *Point {
	return &Point{
		h: 0,
		p: 0,
		n: 0,
	}
}

func (p *Point) Eq(other *Point) bool {
	return p.h == other.h && p.p == other.p && p.n == other.n
}

func (p *Point) Move(direction Direction) *Point {
	switch direction {
	case N:
		return &Point{
			h: p.h,
			p: p.p - 1,
			n: p.n + 1,
		}
	case NE:
		return &Point{
			h: p.h + 1,
			p: p.p - 1,
			n: p.n,
		}
	case SE:
		return &Point{
			h: p.h + 1,
			p: p.p,
			n: p.n - 1,
		}
	case S:
		return &Point{
			h: p.h,
			p: p.p + 1,
			n: p.n - 1,
		}
	case SW:
		return &Point{
			h: p.h - 1,
			p: p.p + 1,
			n: p.n,
		}
	case NW:
		return &Point{
			h: p.h - 1,
			p: p.p,
			n: p.n + 1,
		}
	default:
		return p
	}
}

type Direction uint

const (
	N Direction = iota
	NE
	SE
	S
	SW
	NW
)

func DirectionFromString(str string) (Direction, error) {
	switch str {
	case "n":
		return N, nil
	case "ne":
		return NE, nil
	case "se":
		return SE, nil
	case "s":
		return S, nil
	case "sw":
		return SW, nil
	case "nw":
		return NW, nil
	default:
		return 0, fmt.Errorf("Could not parse %s as a valid direction.", str)
	}
}

func getDirections(directionStrings []string) ([]Direction, error) {
	directions := make([]Direction, len(directionStrings))
	for i, d := range directionStrings {
		dir, err := DirectionFromString(d)
		if err != nil {
			return directions, err
		}
		directions[i] = dir
	}
	return directions, nil
}

func commaSplitter(r rune) bool {
	return r == ','
}

func followDirections(start *Point, directions []Direction) *Point {
	current := start
	for _, d := range directions {
		current = current.Move(d)
	}
	return current
}

type Path struct {
	*Point
	directions []Direction
}

func shortestPath(source, destination *Point) []Direction {
	directions := []Direction{N, NE, SE, S, SW, NW}
	queue := []Path{
		Path{
			Point:      source,
			directions: []Direction{},
		},
	}
	visited := map[Point]bool{
		*source: true,
	}
	next := queue[0]
	for len(queue) > 0 {
		next, queue = queue[0], queue[1:]
		if next.Point.Eq(destination) {
			return next.directions
		}
		for _, d := range directions {
			point := next.Point.Move(d)
			path := Path{
				point,
				append(next.directions, d),
			}
			if _, alreadySeen := visited[*point]; alreadySeen {
				continue
			} else {
				visited[*point] = true
				queue = append(queue, path)
			}
		}
	}

	return []Direction{}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	directions, err := getDirections(strings.FieldsFunc(string(input[0:len(input)-1]), commaSplitter))
	if err != nil {
		panic(err)
	}
	origin := OriginPoint()
	destination := followDirections(origin, directions)
	fmt.Println(destination)
	path := shortestPath(origin, destination)
	fmt.Println(len(path))
}
