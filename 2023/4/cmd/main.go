package main

import (
	"github.com/rushdown-cforbes/aoc/2023/4/internal"
	"log"
)

func main() {
	p := internal.Puzzle1
	s1, s2 := p.DoWork()
	log.Println(p.Name, ":", s1)
	log.Println(p.Name, ":", s2)
}
