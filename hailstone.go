package main

import (
	"fmt"
	"log"
	"os"
)

const script = "hailstone"

func main() {
	// Check if any arguments were given
	if len(os.Args[1:]) == 0 {
		logf("Missing argument! Cf. '" + script + " --help' for more detail.")
	}

	// Check if os.Args[1] is a positive integer > 0
	var n uint64
	_, err := fmt.Sscanf(os.Args[1], "%d", &n)
	if err != nil || n <= 0 {
		logf("First argument must be a positive integer > 0!")
	}

	hailstone(n)

	// Print a newline when we are done
	fmt.Println()
}

func hailstone(n uint64) {
	fmt.Printf("%d ", n)

	if n == 1 {
		return
	}

	if n%2 != 0 {
		n = 3*n + 1
		fmt.Printf("%d ", n)
	}

	hailstone(n / 2)
}

func logf(m string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix)
	log.SetPrefix("FATAL ")
	log.Fatalln(m)
}
