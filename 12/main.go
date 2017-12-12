package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Program struct {
	PID   string
	Pipes []string
}

func readPrograms(lines []string) map[string]*Program {
	programs := make(map[string]*Program, len(lines))
	for _, line := range lines {
		fields := strings.FieldsFunc(
			line,
			func(r rune) bool { return r == ' ' || r == ',' },
		)
		pid := fields[0]
		pipes := fields[2:]
		programs[pid] = &Program{
			PID:   pid,
			Pipes: pipes,
		}
	}
	return programs
}

func programGroup(pid string, programs map[string]*Program) []string {
	visited := map[string]bool{
		pid: true,
	}
	group := []string{pid}
	queue := []string{pid}
	var nextPID string
	for len(queue) > 0 {
		nextPID, queue = queue[0], queue[1:]
		visited[nextPID] = true
		for _, pipedPID := range programs[nextPID].Pipes {
			if _, seen := visited[pipedPID]; !seen {
				queue = append(queue, pipedPID)
				group = append(group, pipedPID)
			}
		}
	}
	return group
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	programs := readPrograms(
		strings.FieldsFunc(
			string(input),
			func(r rune) bool { return r == '\n' },
		),
	)
	fmt.Println(len(programGroup("0", programs)))

	pids := make([]string, 0, len(programs))
	for pid, _ := range programs {
		pids = append(pids, pid)
	}
	numGroups := 0
	grouped := map[string]bool{}
	for _, pid := range pids {
		if _, inGroup := grouped[pid]; inGroup {
			continue
		}

		numGroups++
		for _, groupPID := range programGroup(pid, programs) {
			grouped[groupPID] = true
		}
	}
	fmt.Println(numGroups)
}
