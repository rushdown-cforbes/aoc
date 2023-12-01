package main

import (
	"log"
	"os"
	"strconv"
)

var numbers = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	fileB, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	lines := 0
	var firstDigit, lastDigit string
	for index, v := range fileB {
		c := string(v)
		if c == "\n" {
			lines++
			if lastDigit == "" {
				lastDigit = firstDigit
			}
			num, err := strconv.Atoi(firstDigit + lastDigit)
			if err == nil {
				log.Println(num)
				total += num
				firstDigit = ""
				lastDigit = ""
			}
			continue
		}

		if _, err := strconv.Atoi(c); err == nil {
			if firstDigit == "" {
				firstDigit = c
			} else {
				lastDigit = c
			}
		} else {
			numStr, valid := checkNumber(fileB, index)
			if valid {
				if firstDigit == "" {
					firstDigit = numStr
				} else {
					lastDigit = numStr
				}
			}
		}
	}
	log.Println(total)
	log.Println(lines)
}

func checkNumber(b []byte, i int) (string, bool) {
	for num, numI := range numbers {
		if len(b) < i+len(num) {
			continue
		}
		for letterIndex := range num {
			if b[i+letterIndex] == num[letterIndex] {
				if letterIndex+1 == len(num) {
					return numI, true
				}
			} else {
				break
			}
		}
	}
	return "", false
}
