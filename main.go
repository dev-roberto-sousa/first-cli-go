package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read a text from a Reader (such a file)
	scanner := bufio.NewScanner(r)
	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		// Define the scanner split type to words (default is split by lines)
		scanner.Split(bufio.ScanWords)
	}
	// Define a counter
	words_count := 0
	// For every word scanned, increment the counter
	for scanner.Scan() {
		words_count++
	}
	// Return total
	return words_count

}
