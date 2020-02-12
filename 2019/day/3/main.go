package main

/* --- Day 3: Crossed Wires ---

The gravity assist was successful, and you're well on your way to the Venus refuelling station. During the rush back on Earth, the fuel management system wasn't completely installed, so that's next on the priority list.

Opening the front panel reveals a jumble of wires. Specifically, two wires are connected to a central port and extend outward on a grid. You trace the path each wire takes as it leaves the central port, one wire per line of text (your puzzle input).

The wires twist and turn, but the two wires occasionally cross paths. To fix the circuit, you need to find the intersection point closest to the central port. Because the wires are on a grid, use the Manhattan distance for this measurement. While the wires do technically cross right at the central port where they both start, this point does not count, nor does a wire count as crossing with itself.

For example, if the first wire's path is R8,U5,L5,D3, then starting from the central port (o), it goes right 8, up 5, left 5, and finally down 3:

...........
...........
...........
....+----+.
....|....|.
....|....|.
....|....|.
.........|.
.o-------+.
...........

Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4, and left 4:

...........
.+-----+...
.|.....|...
.|..+--X-+.
.|..|..|.|.
.|.-X--+.|.
.|..|....|.
.|.......|.
.o-------+.
...........

These wires cross at two locations (marked X), but the lower-left one is closer to the central port: its distance is 3 + 3 = 6.

Here are a few more examples:

    R75,D30,R83,U83,L12,D49,R71,U7,L72
    U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
    R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
    U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

What is the Manhattan distance from the central port to the closest intersection?

To begin, get your puzzle input. */

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct{ x, y, dist int }
type coords map[coord]int

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// scan tokens (lines) from the puzzle input
	c1 := make(coords)
	c2 := make(coords)

	// First grid, first line
	if scanner.Scan() {
		c1 = solve(strings.Split(scanner.Text(), ","))
	}

	// You can guess.
	if scanner.Scan() {
		c2 = solve(strings.Split(scanner.Text(), ","))
	}

	log.Println("coords1", c1)
	log.Println("coords2", c2)
}

func solve(paths []string) (c coords) {
	log.Println("paths", paths)

	// Coordinate channel, unbuffered, as one will be emit for each coordinate step.
	ccd := make(chan coord)

	// Quit channel, so when we're done reading, we can kill the other goroutine.
	qc := make(chan struct{}, 1)

	// Start reader with read access to coordinate channel.
	go draw(ccd, qc)

	// Emit coordinates to the unbuffered channel (blocking).
	trace(paths, ccd)

	return
}

func trace(path []string, coords chan<- coord) {
	step := map[string][]int{
		"U": []int{0, 1},
		"D": []int{0, -1},
		"R": []int{1, 0},
		"L": []int{-1, 0},
	}

	var x, y int
	for _, segment := range path {
		direction := string(segment[0])
		dx := step[direction][0]
		dy := step[direction][1]
		distance, _ := strconv.Atoi(segment[1:])
		// log.Println(
		// 	"segment", segment,
		// 	"direction", direction,
		// 	"dx", dx,
		// 	"dy", dy,
		// 	"distance", distance,
		// )

		for n := 0; n <= distance; n++ {
			x += dx
			y += dy
			coords <- coord{x, y, n}
		}
	}

	return
}

func draw(cc chan coord, qc chan struct{}) coords {
	cm := make(coords)
	// Path length is the number of messages received up until this coordinate.
	// In other languages it would be the enumerator index. Perhaps use a slice?

	for {
		select {
		case cd := <-cc:
			if _, ok := cm[cd]; !ok {
				// coordinate does not exist in map
				cm[cd] = cd.dist
			}
			break
		case <-qc:
			return cm
		}
	}

	return cm
}

// ManhattanDistance returns either zero or the distance between two points on a float64 grid of the same length/size.
func ManhattanDistance(a, b []float64) (dist float64) {
	if len(a) != len(b) {
		return 0
	}

	for i := 0; i < len(a); i++ {
		dist += math.Abs(b[i] - a[i])
	}

	return
}
