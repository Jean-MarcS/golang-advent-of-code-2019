package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func testCorruption(imageLayers string, imageHeight int, imageWidth int) (result int) {
	cpt := 0

	minZero := imageHeight * imageWidth

	for cpt < len(imageLayers) {

		nbZero := 0
		nbOne := 0
		nbTwo := 0
		for h := 0; h < imageHeight; h++ {
			currentH := h * imageWidth
			for w := 0; w < imageWidth; w++ {
				switch imageLayers[currentH+w+cpt] {
				case 48:
					nbZero++
					break
				case 49:
					nbOne++
					break
				case 50:
					nbTwo++
					break

				}
			}
		}
		// Test
		if nbZero < minZero {
			minZero = nbZero
			result = nbOne * nbTwo
		}
		cpt += imageHeight * imageWidth
	}

	return result
}

func afficheImage(imageLayers string, imageHeight int, imageWidth int) {
	finalImage := make([]byte, imageHeight*imageWidth)

	// Init transparent image
	for h := 0; h < imageHeight; h++ {
		currentH := h * imageWidth
		for w := 0; w < imageWidth; w++ {
			finalImage[currentH+w] = 50
		}
	}
	cpt := 0

	for cpt < len(imageLayers) {
		for h := 0; h < imageHeight; h++ {
			currentH := h * imageWidth
			for w := 0; w < imageWidth; w++ {
				// Change only if transparent
				if finalImage[currentH+w] == 50 {
					finalImage[currentH+w] = imageLayers[currentH+w+cpt]
				}
			}
		}
		cpt += imageHeight * imageWidth
	}

	// Print image
	for h := 0; h < imageHeight; h++ {
		currentH := h * imageWidth
		for w := 0; w < imageWidth; w++ {
			switch finalImage[currentH+w] {
			case 48:
				fmt.Print(" ")
				break
			case 49:
				fmt.Print("X")
				break
			case 50:
				fmt.Print(" ")
				break
			}
		}
		fmt.Println()
	}
}

func main() {

	imageLayers := ""
	imageWidth := 25
	imageHeight := 6

	file, err := os.Open("day8.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		imageLayers = scanner.Text()
	}

	// Part 1
	fmt.Print("Part 1 result : ")
	fmt.Println(testCorruption(imageLayers, imageHeight, imageWidth))

	// Part 2
	fmt.Println("Part 2 result : ")
	afficheImage(imageLayers, imageHeight, imageWidth)
}
