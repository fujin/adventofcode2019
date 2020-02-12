package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

// Recursively calculate fuel for given mass.
//
// Fuel required to launch a given module is based on its mass. Specifically, to
// find the fuel required for a module, take its mass, divide by three, round
// down, and subtract 2.
//
// Fuel itself requires fuel just like a module - take its mass, divide by
// three, round down, and subtract 2. However, that fuel also requires fuel, and
// that fuel requires fuel, and so on. Any mass that would require negative fuel
// should instead be treated as if it requires zero fuel; the remaining mass, if
// any, is instead handled by wishing really hard, which has no mass and is
// outside the scope of this calculation.
func fuel(mass float64) float64 {
	total := math.Floor(mass/3) - 2

	if total > 0 && fuel(total) > 0 {
		total += fuel(total)
	}

	return total
}

func main() {
	// Total of how much fuel required by each module (and all required fuel)
	sum := 0.0

	// open the puzzle input
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	// get a buffered scanner
	scanner := bufio.NewScanner(f)

	// scan tokens (lines) from the puzzle input
	for scanner.Scan() {
		if mass, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
			sum += fuel(mass)
		}
	}

	log.Printf("%f", sum)
}
