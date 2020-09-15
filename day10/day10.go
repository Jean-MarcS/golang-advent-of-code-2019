package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func meteoriteCount(inputMeteorites []string, inputWidth int, inputHeight int, X int, Y int) int {
	type vector struct {
		dx    int
		dy    int
		angle float64
	}

	var vectors []vector
	for x := 0; x < inputHeight; x++ {
		for y := 0; y < inputWidth; y++ {
			if inputMeteorites[x][y:y+1] == "#" {
				// Check it's not the meteorite we are testing
				if !(x == X && y == Y) {
					// Determine vector
					dx := x - X
					dy := y - Y
					//angle := math.Acos(float64(dy)/math.Sqrt(float64(dx)*float64(dx)+float64(dy)*float64(dy))) * 180 / math.Pi
					// wrong method, probably due to approximation. Got 246 instead of 247
					// So this one is better
					angle := math.Round(math.Atan2(float64(x-X), float64(y-Y)) * 180 / math.Pi)
					// Check if a colinear vector already exists
					found := false
					for _, currentVector := range vectors {
						if ((currentVector.dx*dy)-(currentVector.dy*dx) == 0) && angle == currentVector.angle {
							found = true
						}
					}
					if !found {
						vectors = append(vectors, vector{dx: dx, dy: dy, angle: angle})
					}
				}
			}
		}
	}
	return len(vectors)
}

func part1(inputMeteorites []string, inputWidth int, inputHeight int) {
	meteoriteSum := make([]int, inputWidth*inputHeight)
	max := 0
	for x := 0; x < inputHeight; x++ {
		for y := 0; y < inputWidth; y++ {
			if inputMeteorites[x][y:y+1] == "#" {
				meteoriteSum[x*inputHeight+y] = meteoriteCount(inputMeteorites, inputWidth, inputHeight, x, y)
				if meteoriteSum[x*inputHeight+y] > max {
					max = meteoriteSum[x*inputHeight+y]
				}
			}
		}
	}
	fmt.Printf("Max value : %d\n", max)
}

func main() {

	inputWidth := 24
	inputHeight := 24
	inputMeteorites := make([]string, inputHeight)

	file, err := os.Open("day10.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		inputMeteorites[i] = scanner.Text()
		i++
	}

	part1(inputMeteorites, inputWidth, inputHeight)

}
