// https://adventofcode.com/2024/day/1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in, err := os.Open("01/input_test")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	defer in.Close()

	out, err := os.Create("01/output")
	if err != nil {
		log.Fatalf("Error creating output file: %s", err)
	}
	defer out.Close()

	s := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		nums := strings.Fields(line)
		numParts := len(nums)
		if numParts != 2 {
			log.Fatalf("Error parsing: %d entries instead of 2", numParts)
		}
		_, err := fmt.Fprintf(w, "%s\n", nums[0])
		if err != nil {
			log.Fatalf("Error writing to file: %s", err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

}
