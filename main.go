package main

import (
	"log"
	"flag"
)


func main() {
	day := flag.Int("day", 0, "Day of the puzzle (1-25)")
	part := flag.Int("part", 0, "Part of the puzzle (1 or 2)")
	flag.Parse()

	if *day == 0 || *part == 0 {
		log.Fatal("Please specify both --day and --part")
	}

	solutions[*day][*part]()
}


