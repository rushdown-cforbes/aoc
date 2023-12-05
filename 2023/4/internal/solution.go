package internal

import (
	"log"
	"strconv"
	"strings"
)

type Card struct {
	id          int
	winning     []int
	scratched   []int
	copyIndices []int
}

var (
	cardSum        = 0
	processedCards []*Card
)

func (p *Puzzle) DoWork() (puzzle1Solution, puzzle2Solution int) {
	for lineIndex, line := range strings.Split(string(p.InputData), "\n") {
		if line == "" {
			continue
		}
		log.Println(line)
		card := &Card{id: lineIndex + 1}
		processedCards = append(processedCards, card)
		line := line[5:]

		idSeperated := strings.Split(line, ":")
		_, cards := idSeperated[0], idSeperated[1]

		cardValues := strings.Split(cards, "|")
		winning, yours := cardValues[0], cardValues[1]
		for _, winningStr := range strings.Split(winning, " ") {
			if winningStr != "" {
				if num, err := strconv.Atoi(winningStr); err == nil {
					card.winning = append(card.winning, num)
				}
			}

		}
		for _, yourStr := range strings.Split(yours, " ") {
			if yourStr != "" {
				if num, err := strconv.Atoi(yourStr); err == nil {
					card.scratched = append(card.scratched, num)
				}
			}
		}

		matchedCount := 0
		points := 0
		for _, potential := range card.scratched {
			for _, win := range card.winning {
				if potential == win {
					matchedCount++
					if matchedCount == 1 {
						points++
					} else {
						points *= 2
					}
					card.copyIndices = append(card.copyIndices, card.id+matchedCount)
				}
			}
		}
		cardSum += points
	}

	for _, card := range processedCards {
		log.Println("original card ", card.id, card.copyIndices)
	}

	totalCards := len(processedCards)
	for _, card := range processedCards {
		log.Println("processing card", card.id)
		cardTotalCopies := processCardCopies(card.copyIndices, 0)
		totalCards += cardTotalCopies
	}
	return cardSum, totalCards
}

func processCardCopies(ids []int, count int) int {
	var copyIndices []int
	for _, i := range ids {
		if i > len(processedCards) {
			continue
		}
		count++
		card := processedCards[i-1]
		copyIndices = append(copyIndices, card.copyIndices...)
	}
	if len(copyIndices) < 1 {
		return count
	}
	return processCardCopies(copyIndices, count)
}
