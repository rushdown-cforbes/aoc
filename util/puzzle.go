package util

import "log"

const (
	puzzle1Name = "Puzzle1"
	puzzle2Name = "Puzzle2"
	sampleName  = "Sample"
)

type Puzzle struct {
	Name      string
	InputPath string
	InputData []byte
}

func (p *Puzzle) SetPath(path string) {
	p.InputPath = path
}

func (p *Puzzle) DoWork() (puzzle1Solution, puzzle2Solution int) {
	log.Println(p.Name, "main function not provided. Should be overridden in the package for each day")
	return -1, -1
}

func (p *Puzzle) Solve() int {
	s1, s2 := p.DoWork()

	switch p.Name {
	case puzzle1Name:
		return s1
	case puzzle2Name:
		return s2
	}
	return -1
}
