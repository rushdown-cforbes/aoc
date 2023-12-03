package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileB, err := os.ReadFile("2023/2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games := make(map[int][]map[string]int)

	lines := strings.Split(string(fileB), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		id, data := parts[0], parts[1]
		id = id[5:]
		rounds := strings.Split(data, ";")

		idNum, _ := strconv.Atoi(id)
		games[idNum] = make([]map[string]int, len(rounds))

		log.Println("game ", id, " data ", data)
		for roundIndex, round := range rounds {
			colorMap := make(map[string]int)
			log.Println("roll", round)
			for _, roll := range strings.Split(round, ",") {
				rollData := strings.Split(strings.TrimSpace(roll), " ")
				numStr, color := rollData[0], strings.ToLower(rollData[1])
				num, _ := strconv.Atoi(numStr)
				if _, ok := colorMap[color]; ok {
					colorMap[color] += num
				} else {
					colorMap[color] = num
				}
			}

			games[idNum][roundIndex] = colorMap
		}
	}

	//firstQuestion := map[string]int{"red": 12, "green": 13, "blue": 14}
	firstQuestion := map[string]int{"red": 12, "green": 13, "blue": 14}
	totalIds := 0
	fewestRequired := make(map[int]map[string]int)
	for gameId, game := range games {
		fewestRequired[gameId] = make(map[string]int)
		valid := true
		for roundIndex, round := range game {
			for color, max := range firstQuestion {
				roundColorValue := round[color]
				if roundIndex == 0 {
					fewestRequired[gameId][color] = roundColorValue
				} else if roundColorValue > fewestRequired[gameId][color] {
					fewestRequired[gameId][color] = roundColorValue
				}
				if roundColorValue > max {
					valid = false
				}
			}
		}
		if valid {
			log.Println("valid game ", gameId)
			totalIds += gameId
		}
	}
	log.Println("total valid ids", totalIds)
	powersSum := 0
	for gameId, fewest := range fewestRequired {
		log.Println(fewest)
		powerSet := 0
		for _, power := range fewest {
			if powerSet == 0 {
				powerSet = power
			} else {
				powerSet *= power
			}
		}
		log.Println("power set for game", gameId, "is", powerSet)
		powersSum += powerSet
	}
	log.Println("sum of power sets is", powersSum)
}
