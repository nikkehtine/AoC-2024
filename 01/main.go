// https://adventofcode.com/2024/day/1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("01/input_test")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatalf("Error reading input: %s", err)
	}
}
