package main

import (
	"fmt"
)

const (
	INPUT         = 343
	FIFTY_MILLION = 50000000
)

type SpinLock struct {
	vals      []int
	spin, pos int
}

func NewSpinLock(spin int) *SpinLock {
	return &SpinLock{
		vals: []int{0},
		spin: spin,
		pos:  0,
	}
}

func (self *SpinLock) Step() *SpinLock {
	insertLoc := ((self.pos + self.spin) % len(self.vals)) + 1
	newVals := make([]int, 0, len(self.vals)+1)

	for i := 0; i < insertLoc; i++ {
		newVals = append(newVals, self.vals[i])
	}
	newVals = append(newVals, len(self.vals))
	newVals = append(newVals, self.vals[insertLoc:]...)

	return &SpinLock{
		vals: newVals,
		spin: self.spin,
		pos:  insertLoc,
	}
}

func (self *SpinLock) ValueAfter(val int) int {
	for i, v := range self.vals {
		if v == 2017 {
			if i == len(self.vals)-1 {
				break
			}
			return self.vals[i+1]
		}
	}
	return self.vals[0]
}

func main() {
	lock := NewSpinLock(INPUT)
	for i := 0; i <= 2017; i++ {
		lock = lock.Step()
	}

	fmt.Println(lock.ValueAfter(2017))

	firstValue := 1
	position := 1
	for i := 2; i < FIFTY_MILLION; i++ {
		position = ((position + INPUT) % i) + 1
		if position == 1 {
			firstValue = i
		}
	}
	fmt.Println(firstValue)
}
