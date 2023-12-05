package internal

import (
	"github.com/rushdown-cforbes/aoc/util"
	"log"
	"os"
)

type Puzzle struct {
	util.Puzzle
}

var (
	Sample  = &Puzzle{}
	Sample2 = &Puzzle{}
	Puzzle1 = &Puzzle{}
	//Puzzle2 = &Puzzle{InputPath: puzzle2Path}
	puzzles = []*Puzzle{Sample, Puzzle1, Sample2}
)

func init() {
	Sample.SetPath(samplePath)
	Puzzle1.SetPath(puzzle1Path)
	Sample2.SetPath(sample2Path)

	errs := 0
	for _, p := range puzzles {
		fileB, err := os.ReadFile(p.InputPath)
		if err != nil {
			errs++
			continue
		}
		p.InputData = fileB
	}
	if errs == len(puzzles) {
		log.Fatal("failed to find input files, the base path was", basePath)
	}
}
