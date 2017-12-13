package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Firewall struct {
	Layers map[int]*Layer
}

func (f *Firewall) Step() *Firewall {
	layers := make(map[int]*Layer, len(f.Layers))
	for i, layer := range f.Layers {
		layers[i] = layer.Step()
	}
	return &Firewall{
		Layers: layers,
	}
}

func (f *Firewall) Caught(layer int) bool {
	if l, exists := f.Layers[layer]; exists {
		return l.Caught()
	} else {
		return false
	}
}

func (f *Firewall) Severity(layer int) int {
	return f.Layers[layer].Range * layer
}

type Layer struct {
	Range     int
	Scanner   int
	direction int
}

func NewLayer(layerRange int) *Layer {
	return &Layer{
		Range:     layerRange,
		Scanner:   0,
		direction: 1,
	}
}

func (l *Layer) Step() *Layer {
	layer := NewLayer(l.Range)
	layer.direction = l.direction
	layer.Scanner = l.Scanner
	if l.direction > 0 {
		if l.Scanner == (l.Range - 1) {
			layer.direction = -1
			layer.Scanner--
		} else {
			layer.Scanner++
		}
	} else {
		if l.Scanner == 0 {
			layer.direction = 1
			layer.Scanner++
		} else {
			layer.Scanner--
		}
	}
	return layer
}

func (l *Layer) Caught() bool {
	return l.Scanner == 0
}

func makeLayers(lines []string) (map[int]*Layer, int, error) {
	layers := make(map[int]*Layer, len(lines))
	maxLayer := 0
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ':' })
		layerNumber, err := strconv.Atoi(fields[0])
		if err != nil {
			return layers, 0, err
		}

		layerRange, err := strconv.Atoi(fields[1][1:])
		if err != nil {
			return layers, 0, err
		}

		layers[layerNumber] = NewLayer(layerRange)
		maxLayer = layerNumber
	}
	return layers, maxLayer, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	layerLines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	layers, maxLayer, err := makeLayers(layerLines)
	if err != nil {
		panic(err)
	}

	firewall := &Firewall{
		Layers: layers,
	}

	layer, severity := 0, 0
	for layer <= maxLayer {
		if firewall.Caught(layer) {
			severity += firewall.Severity(layer)
		}
		firewall = firewall.Step()
		layer++
	}
	fmt.Println(severity)
}
