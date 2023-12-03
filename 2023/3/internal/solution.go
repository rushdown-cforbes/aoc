package internal

import (
	"log"
	"strconv"
	"strings"
)

var wholeNumbers = make([]*wholeNumberInfo, 0)

func (p *Puzzle) DoWork() (puzzle1Solution, puzzle2Solution int) {
	rows := strings.Split(string(p.InputData), "\n")
	symbolCoords := make(map[int]map[int]struct{})
	numberCoords := make(map[int]map[int]*wholeNumberInfo)
	wholeNumbers = make([]*wholeNumberInfo, 0)

	gears := make([]*gearInfo, 0)

	for y, row := range rows {
		chars := []rune(row)
		numInfo := newWholeNumberInfo(y)

		for x, cell := range chars {
			if cell == '.' {
				// build the whole number
				if numInfo.tryBuild() {
					wholeNumbers = append(wholeNumbers, numInfo)
					numInfo = newWholeNumberInfo(y)
				}

				continue
			}

			if symbolCoords[x] == nil {
				symbolCoords[x] = make(map[int]struct{})
			}
			if numberCoords[x] == nil {
				numberCoords[x] = make(map[int]*wholeNumberInfo)
			}

			if num, err := strconv.Atoi(string(cell)); err == nil {
				numberCoords[x][y] = &wholeNumberInfo{
					buildStr: "",
					isValid:  false,
					value:    num,
				}
				numInfo.buildStr += string(cell)
				numInfo.addX(x)
			} else {
				symbolCoords[x][y] = struct{}{}
				gears = append(gears, newGear(x, y))
				// build the whole number
				if numInfo.tryBuild() {
					wholeNumbers = append(wholeNumbers, numInfo)
					numInfo = newWholeNumberInfo(y)
				}
			}
		}
		// before moving on to the next row, build whatever number was started
		if numInfo.tryBuild() {
			wholeNumbers = append(wholeNumbers, numInfo)
			// No need for new - will be replaced in next loop.
		}
	}

	for x, yMap := range numberCoords {
		for y := range yMap {
			// Fucking gross, but functional.
			// Do something with range +-1 on x and y
			left := x - 1
			right := x + 1
			above := y - 1
			below := y + 1
			if _, ok := symbolCoords[left][y]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[right][y]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[x][above]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[x][below]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[left][above]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[left][below]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[right][above]; ok {
				tryValidateWholeNumberAt(x, y)
			} else if _, ok := symbolCoords[right][below]; ok {
				tryValidateWholeNumberAt(x, y)
			}
		}
	}

	log.Println(len(gears))
	for _, gear := range gears {
		for _, adjCoord := range gear.location.buildAdjacent() {
			for _, numInfo := range wholeNumbers {
				if numInfo.value == 35 {
					log.Println(gear.location)
				}
				for _, digitX := range numInfo.x {
					if adjCoord.x == digitX && adjCoord.y == numInfo.y {
						if gear.location.x == 3 && gear.location.y == 1 && numInfo.value == 35 {
							log.Println("found adj number", numInfo)
						}
						if !gear.hasPart(numInfo) {
							gear.adjacentParts = append(gear.adjacentParts, numInfo)
						}
						break
					}
				}
			}
		}
	}

	sumRatio := 0
	for _, gear := range gears {
		if gear.IsValid() {
			log.Println("gear valid parts", gear.adjacentParts[0].value, gear.adjacentParts[1].value)
			sumRatio += gear.adjacentParts[0].value * gear.adjacentParts[1].value
		}
	}

	sum := 0
	for _, num := range wholeNumbers {
		if num.isValid {
			sum += num.value
		}
	}
	return sum, sumRatio
}

type wholeNumberInfo struct {
	buildStr string
	y        int
	x        []int
	isValid  bool
	value    int
}

func (w *wholeNumberInfo) tryBuild() bool {
	if w.buildStr == "" {
		return false
	}

	wholeNum, err := strconv.Atoi(w.buildStr)
	if err != nil {
		log.Println("error parsing whole number", err)
		return false
	}

	w.value = wholeNum

	return true
}

func (w *wholeNumberInfo) addX(x int) {
	w.x = append(w.x, x)
}

func (w *wholeNumberInfo) contains(x, y int) bool {
	if w.y != y {
		return false
	}

	for _, wx := range w.x {
		if wx == x {
			return true
		}
	}

	return false
}

func newWholeNumberInfo(y int) *wholeNumberInfo {
	return &wholeNumberInfo{
		y: y,
		x: make([]int, 0),
	}
}

func tryValidateWholeNumberAt(x int, y int) {
	for _, info := range wholeNumbers {
		if info.contains(x, y) {
			info.isValid = true
		}
	}
}

type coordinate struct {
	x int
	y int
}

func (c coordinate) buildAdjacent() []coordinate {
	left := c.x - 1
	right := c.x + 1
	above := c.y - 1
	below := c.y + 1

	return []coordinate{
		{x: left, y: c.y},
		{x: right, y: c.y},
		{x: c.x, y: above},
		{x: c.x, y: below},
		{x: left, y: above},
		{x: right, y: above},
		{x: left, y: below},
		{x: right, y: below},
	}
}

func newGear(x, y int) *gearInfo {
	return &gearInfo{
		location:      coordinate{x, y},
		adjacentParts: make([]*wholeNumberInfo, 0),
	}
}

type gearInfo struct {
	location      coordinate
	adjacentParts []*wholeNumberInfo
}

func (gi *gearInfo) IsValid() bool {
	return len(gi.adjacentParts) == 2
}

func (gi *gearInfo) hasPart(info *wholeNumberInfo) bool {
	for _, part := range gi.adjacentParts {
		if part.y != info.y || len(part.x) != len(info.x) {
			continue
		}
		for i, x := range info.x {
			if x == part.x[i] {
				return true
			} else {
				continue
			}
		}
	}
	return false
}
