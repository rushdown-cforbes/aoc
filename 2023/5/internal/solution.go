package internal

import (
	"log"
	"strconv"
	"strings"
)

type category struct {
	name    string
	lookups []*lookupInfo
}

type lookupInfo struct {
	source int
	dest   int
	r      int
}

func (p *Puzzle) DoWork() (puzzle1Solution, puzzle2Solution int) {
	var (
		seeds           []int
		categories      []*category
		currentCategory *category
	)

	categoryIndicator := false
	for lineIndex, line := range strings.Split(string(p.InputData), "\n") {
		if lineIndex == 0 {
			log.Println(line)
			line := line[7:]
			for _, numStr := range strings.Split(line, " ") {
				if num, err := strconv.Atoi(numStr); err == nil {
					seeds = append(seeds, num)
				}
			}
		} else if line == "" {
			if currentCategory != nil {
				categories = append(categories, currentCategory)
			}
			categoryIndicator = true
			continue
		} else if categoryIndicator {
			currentCategory = &category{
				name:    line,
				lookups: make([]*lookupInfo, 0),
			}
			categoryIndicator = false
		} else {
			lookup := &lookupInfo{}
			for index, numStr := range strings.Split(line, " ") {
				if num, err := strconv.Atoi(numStr); err == nil {
					switch index {
					case 0:
						lookup.dest = num
					case 1:
						lookup.source = num
					case 2:
						lookup.r = num
					}
				}
			}
			currentCategory.lookups = append(currentCategory.lookups, lookup)
		}
	}

	for _, category := range categories {
		log.Println("category name:", category.name)
		for _, lookup := range category.lookups {
			log.Println("lookup ", lookup.dest, " ", lookup.source, " ", lookup.r)
		}
	}

	return -1, -1
}
