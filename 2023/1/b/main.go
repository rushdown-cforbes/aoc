package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	fileB, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	var firstDigit, lastDigit string
	for _, v := range fileB {
		c := string(v)
		if c == "\n" {
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
		}
	}
	log.Println(total)
}
