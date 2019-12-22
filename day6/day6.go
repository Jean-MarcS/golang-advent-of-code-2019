package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func orbitsNumber(orbits map[string][]string, planet string, deep int) int {

	total := deep
	// Check if planet has other planets orbiting around
	if len(orbits[planet]) > 0 {
		for _, p := range orbits[planet] {
			total += orbitsNumber(orbits, p, deep+1)
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
	orbits := make(map[string][]string)

	for scanner.Scan() {
		orbit := scanner.Text()
		pos := strings.Index(orbit, ")")
		if pos != -1 {
			planetA := orbit[0:pos]
			planetB := orbit[pos+1:]

			// Test if exists
			_, ok := orbits[planetA]

			if !ok {
				orbits[planetA] = []string{planetB}
			} else {
				orbits[planetA] = append(orbits[planetA], planetB)
			}

		}
	}

	fmt.Print("Part 1 result : ")
	fmt.Println(orbitsNumber(orbits, "COM", 0))
}
