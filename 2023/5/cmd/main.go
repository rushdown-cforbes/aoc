package main

import (
	"github.com/rushdown-cforbes/aoc/2023/5/internal"
	"log"
)

func main() {
	p := internal.Sample
	s1, s2 := p.DoWork()
	log.Println(p.Name, ":", s1)
	log.Println(p.Name, ":", s2)
}
