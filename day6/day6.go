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

func reachable(satellites map[string][]string, currentPlanet string, testingPlanet string) (bool, int) {

	distance := 1
	found := false
	// Check if planet has other planets orbiting around
	if len(satellites[currentPlanet]) > 0 {
		for _, p := range satellites[currentPlanet] {
			found, distance = reachable(satellites, p, testingPlanet)
			if found {
				distance++
				break
			}
		}
	} else {
		if currentPlanet == testingPlanet {
			distance = 0
			found = true
		}
	}

	return found, distance
}

func bothPossible(satellites map[string][]string, planet string) (bool, int) {
	foundYou, distanceYou := reachable(satellites, planet, "YOU")
	foundSanta, distanceSanta := reachable(satellites, planet, "SAN")

	if foundYou && foundSanta {
		return true, distanceYou + distanceSanta - 2
	} else {
		return false, 0
	}
}

func shortestDistance(satellites map[string][]string, planet string) int {

	found, distance := bothPossible(satellites, planet)
	if found {
		if len(satellites[planet]) > 0 {
			for _, p := range satellites[planet] {
				distanceNew := shortestDistance(satellites, p)
				if distanceNew < distance {
					distance = distanceNew
				}
			}
		}
		return distance
	}
	return 99999 // Dumb high value
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
	fmt.Print("Part 2 result : ")
	fmt.Println(shortestDistance(satellites, "COM"))

}
