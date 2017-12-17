package main

import (
	"fmt"
)

const (
	INPUT = 343
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

func main() {
	lock := NewSpinLock(INPUT)
	for i := 0; i <= 2017; i++ {
		lock = lock.Step()
	}

	for i, v := range lock.vals {
		if v == 2017 {
			if i == len(lock.vals)-1 {
				i = -1
			}
			fmt.Println(lock.vals[i+1])
			break
		}
	}
}
