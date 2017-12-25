package main

import (
	"fmt"
)

type TuringMachine struct {
	tape     map[int]int
	position int
	state    string
}

func NewTuringMachine() *TuringMachine {
	return &TuringMachine{
		tape:     map[int]int{},
		position: 0,
		state:    "A",
	}
}

func (self *TuringMachine) Step() {
	switch self.state {
	case "A":
		self.stepA()
	case "B":
		self.stepB()
	case "C":
		self.stepC()
	case "D":
		self.stepD()
	case "E":
		self.stepE()
	case "F":
		self.stepF()
	}
}

func (self *TuringMachine) set() bool {
	v, ok := self.tape[self.position]
	return ok && v == 1
}

func (self *TuringMachine) stepA() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position++
		self.state = "B"
	} else {
		self.tape[self.position] = 1
		self.position--
		self.state = "E"
	}
}

func (self *TuringMachine) stepB() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position++
		self.state = "C"
	} else {
		self.tape[self.position] = 1
		self.position++
		self.state = "F"
	}
}

func (self *TuringMachine) stepC() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position--
		self.state = "D"
	} else {
		self.tape[self.position] = 0
		self.position++
		self.state = "B"
	}
}

func (self *TuringMachine) stepD() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position++
		self.state = "E"
	} else {
		self.tape[self.position] = 0
		self.position--
		self.state = "C"
	}
}

func (self *TuringMachine) stepE() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position--
		self.state = "A"
	} else {
		self.tape[self.position] = 0
		self.position++
		self.state = "D"
	}
}

func (self *TuringMachine) stepF() {
	if !self.set() {
		self.tape[self.position] = 1
		self.position++
		self.state = "A"
	} else {
		self.tape[self.position] = 1
		self.position++
		self.state = "C"
	}
}

func main() {
	tm := NewTuringMachine()
	ITERATIONS := 12459852
	for i := 0; i < ITERATIONS; i++ {
		tm.Step()
	}
	checksum := 0
	for _, v := range tm.tape {
		checksum += v
	}
	fmt.Println(checksum)
}
