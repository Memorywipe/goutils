package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: %s [file]", os.Args[0])
	} else {
		fmt.Println(sumFile(os.Args[1]))
	}
}

func sumFile(path string) float64 {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	sum := 0.0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

		if err != nil  {
			fmt.Fprintln(os.Stderr, err)
		} else {
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return sum
}
