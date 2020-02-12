package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// scan tokens (lines) from the puzzle input
	for scanner.Scan() {
		// instead of lines, we need to parse up integers and commas ("intcodes") made up of opcodes.
		// the opcodes are '1', '2', and '99'.
		// opcode 1 adds together two numbers stored at certain positions, stored in a third position
		// opcode 2 multiplies two integers stored at certain positions, stored in a third position
		// opcode 99 halts the program
		stringcodes := strings.Split(scanner.Text(), ",")
		intcodes := make([]int, len(stringcodes))
		for i, s := range stringcodes {
			intcodes[i], _ = strconv.Atoi(s)
		}

		for x := 0; x <= 99; x++ {
			for y := 0; y <= 99; y++ {
				tmp := make([]int, len(intcodes))
				copy(tmp, intcodes)
				computer := &Computer{intcodes: tmp}
				computer.setInputs(x, y)
				computer.parse()
				if computer.output() == 19690720 {
					log.Println("found!", x, y)
					log.Println("answer:", 100*x+y)
				}
			}
		}
	}
}

// opcode represents the individual instructions to the Intcode computer.
// The first value, 'op', is the type of operation, either add and store, or multiply and store.
// The values 'a' and 'b' are the positions of the intcodes (in the intcode computer) used as parameters to the add or multiply functions.
// The value x is the intcode position where the result of the function will be stored.
type opcode struct {
	op,
	a,
	b,
	x int
}

// Opcode constants map the type of operation to a particular function.
const (
	OpcodeAddAndStore      = 1
	OpcodeMultiplyAndStore = 2
	OpcodeHalt             = 99
)

// Computer represents an Intcode computer used in AOC2019's Day 2 challenge.
//
// It can run intcode programs, which are described as:
// An Intcode program is a list of integers separated by commas (like 1,0,0,3,99). To run one, start by looking at the first integer (called position 0). Here, you will find an opcode - either 1, 2, or 99. The opcode indicates what to do; for example, 99 means that the program is finished and should immediately halt. Encountering an unknown opcode means something went wrong.
type Computer struct {
	intcodes []int
}

// NewComputer whips us up a new Intcode computer, with the specified intcodes, expected as a slice of integers.
func NewComputer(intcodes []int) *Computer {
	return &Computer{intcodes: intcodes}
}

func (c *Computer) run(oc opcode) {
	switch oc.op {
	case OpcodeAddAndStore:
		c.intcodes[oc.x] = c.intcodes[oc.a] + c.intcodes[oc.b]
	case OpcodeMultiplyAndStore:
		c.intcodes[oc.x] = c.intcodes[oc.a] * c.intcodes[oc.b]
	}
}

// The inputs should still be provided to the program by replacing the values at addresses 1 and 2, just like before.
// In this program, the value placed in address 1 is called the noun, and the value placed in address 2 is called the verb.
// Each of the two input values will be between 0 and 99, inclusive.
func (c *Computer) setInputs(noun, verb int) {
	c.intcodes[1] = noun
	c.intcodes[2] = verb
}

func (c *Computer) output() int {
	return c.intcodes[0]
}

func (c *Computer) parse() {
	// read position 0
	// if 1 or 2, read input variable positions 0+1, 0+2 and output position from 0+3 into inputX, inputY, output, advance 4 positions
	// else if 99, return
	// Start looking at the head for the first opcode and potential parameters
	for idx := 0; idx <= len(c.intcodes); {
		switch c.intcodes[idx] {
		case 1, 2:
			opcode := opcode{
				op: c.intcodes[idx],
				a:  c.intcodes[idx+1],
				b:  c.intcodes[idx+2],
				x:  c.intcodes[idx+3],
			}
			// log.Println("loading", opcode)
			// Run the opcode against the intcode storage, potentially manipulating storage or positional values in-place.
			c.run(opcode)
			// As we found a correct opcode (1 or 2), advance the intcode position by the number of values in the instruction.
			idx += 4
		case 99:
			// log.Println("halt")
			return
		}
	}
}
