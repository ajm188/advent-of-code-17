package main

import (
	"fmt"
	"strconv"
	"strings"
)

import (
	"github.com/ajm188/advent-of-code-17/util"
)

const (
	INPUT = "ljoxqyyw"
)

type Point struct {
	r, c int
}

func (p Point) Neighbors() []Point {
	return []Point{
		Point{p.r + 1, p.c},
		Point{p.r - 1, p.c},
		Point{p.r, p.c + 1},
		Point{p.r, p.c - 1},
	}
}

func (p Point) InRange(grid []string) bool {
	return p.r >= 0 && p.c >= 0 && p.r < len(grid) && p.c < len(grid[p.r])
}

func findRegions(grid []string) map[int][]Point {
	pointMap := make(map[Point]int, len(grid)*len(grid))
	regionMap := make(map[int][]Point, 0)
	region := 1
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			p := Point{r, c}
			if _, hasRegion := pointMap[p]; hasRegion {
				continue
			}
			if grid[r][c] == '0' {
				pointMap[p] = -1 // space is free, no region
			} else {
				points := discoverRegion(grid, p, pointMap, region)
				regionMap[region] = points
				region++
			}
		}
	}
	return regionMap
}

func discoverRegion(grid []string, point Point, pointMap map[Point]int, region int) []Point {
	points := make([]Point, 0)
	queue := []Point{point}
	for len(queue) > 0 {
		point, queue = queue[0], queue[1:]
		if _, hasRegion := pointMap[point]; hasRegion {
			// point already assigned to region, stop searching
			continue
		}
		if !point.InRange(grid) {
			// point is out of bounds, skip
			continue
		}
		if grid[point.r][point.c] == '0' {
			// region border found, stop searching
			pointMap[point] = -1
			continue
		}
		queue = append(queue, point.Neighbors()...)
		points = append(points, point)
		pointMap[point] = region
	}
	return points
}

func main() {
	ones := 0
	grid := make([]string, 128)
	for i := 0; i < 128; i++ {
		codes := util.AsASCIICodes(fmt.Sprintf("%s-%d", INPUT, i))
		hk := util.NewHashKnot()
		codes = append(codes, []int{17, 31, 73, 47, 23}...)
		for j := 0; j < util.NUM_ROUNDS; j++ {
			hk.RunRound(codes)
		}
		hash := hk.Hash()
		row := ""
		hexRow := ""
		for _, i := range hash {
			hex := fmt.Sprintf("%.2x", i)
			hexRow += hex
			for _, d := range hex {
				var dec int
				switch d {
				case 'a', 'b', 'c', 'd', 'e', 'f':
					dec = int(d-'a') + 10
				default:
					dec, _ = strconv.Atoi(string(d))
				}
				bin := fmt.Sprintf("%.4b", dec)
				row += bin
			}
		}
		grid[i] = row
		ones += strings.Count(row, "1")
	}
	fmt.Println(ones)
	regions := findRegions(grid)
	fmt.Println(len(regions))
}
