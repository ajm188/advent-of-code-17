package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y, Z int
}
type Position Coordinate
type Velocity Coordinate
type Acceleration Coordinate

type Particle struct {
	Position
	Velocity
	Acceleration
}

func readTuple(tuple string) (x, y, z int, err error) {
	err = nil
	x, y, z = 0, 0, 0
	values := strings.FieldsFunc(tuple, func(r rune) bool { return r == ',' })

	x, err = strconv.Atoi(values[0])
	y, err = strconv.Atoi(values[1])
	z, err = strconv.Atoi(values[2])
	return
}

func readParticle(line string) (*Particle, error) {
	fields := strings.FieldsFunc(line, func(r rune) bool { return r == '=' })

	posTuple := fields[1]
	posX, posY, posZ, err := readTuple(posTuple[1 : len(posTuple)-4])
	if err != nil {
		return nil, err
	}

	velTuple := fields[2]
	velX, velY, velZ, err := readTuple(velTuple[1 : len(velTuple)-4])
	if err != nil {
		return nil, err
	}

	accTuple := fields[3]
	accX, accY, accZ, err := readTuple(accTuple[1 : len(accTuple)-1])
	if err != nil {
		return nil, err
	}

	particle := &Particle{
		Position:     Position{posX, posY, posZ},
		Velocity:     Velocity{velX, velY, velZ},
		Acceleration: Acceleration{accX, accY, accZ},
	}
	return particle, nil
}

func readParticles(lines []string) (particles []*Particle, err error) {
	particles = make([]*Particle, len(lines))
	err = nil

	var particle *Particle
	for i, line := range lines {
		particle, err = readParticle(line)
		if err != nil {
			break
		}
		particles[i] = particle
	}
	return
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	particles, err := readParticles(lines)
	if err != nil {
		panic(err)
	}

	acc := math.Inf(1)
	accBuckets := map[float64][]int{}
	for i, particle := range particles {
		manhattan := math.Abs(float64(particle.Acceleration.X)) + math.Abs(float64(particle.Acceleration.Y)) + math.Abs(float64(particle.Acceleration.Z))
		if bucket, ok := accBuckets[manhattan]; ok {
			accBuckets[manhattan] = append(bucket, i)
		} else {
			accBuckets[manhattan] = []int{i}
		}
		if manhattan < acc {
			acc = manhattan
		}

	}

	closest := 0
	vel := math.Inf(1)
	for _, i := range accBuckets[acc] {
		particle := particles[i]
		manhattan := math.Abs(float64(particle.Velocity.X)) + math.Abs(float64(particle.Velocity.Y)) + math.Abs(float64(particle.Velocity.Z))
		if manhattan < vel {
			vel = manhattan
			closest = i
		}
	}

	fmt.Println(closest)
}
