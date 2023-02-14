package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func solve1() int {
	lines, _ := readline("input.txt")
	var highest int
	var current int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if current > highest {
				highest = current
			}
			current = 0
		} else {
			number, _ := strconv.Atoi(lines[i])
			current += number
		}
	}
	return highest

}

func solve2() int {
	lines, _ := readline("input.txt")
	var found []int
	var current int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			found = append(found, current)
			current = 0
		} else {
			number, _ := strconv.Atoi(lines[i])
			current += number
		}
	}
	sort.Ints(found)
	var sum int
	for i := 0; i < 3; i++ {
		sum += found[len(found)-1-i]
	}
	return sum
}

func readline(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("cannot open the file")
	}
	defer file.Close() // meilleure utilisation pour fermer les fichiers
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func main() {
	fmt.Println(solve1())
	fmt.Println(solve2())
}
