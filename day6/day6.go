package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func satellitesNumber(satellites map[string][]string, planet string, deep int) int {

	total := deep
	// Check if planet has other planets orbiting around
	if len(satellites[planet]) > 0 {
		for _, p := range satellites[planet] {
			total += satellitesNumber(satellites, p, deep+1)
		}
	}

	return total
}

func main() {
	file, err := os.Open("day6.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	satellites := make(map[string][]string)

	for scanner.Scan() {
		orbit := scanner.Text()
		pos := strings.Index(orbit, ")")
		if pos != -1 {
			planetA := orbit[0:pos]
			planetB := orbit[pos+1:]

			// Test if exists
			_, ok := satellites[planetA]

			if !ok {
				satellites[planetA] = []string{planetB}
			} else {
				satellites[planetA] = append(satellites[planetA], planetB)
			}

		}
	}

	fmt.Print("Part 1 result : ")
	fmt.Println(satellitesNumber(satellites, "COM", 0))
}
