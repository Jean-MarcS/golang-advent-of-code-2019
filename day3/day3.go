package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type wirePoint struct {
	row int
	col int
}

type wireSection struct {
	p         wirePoint
	direction int // 0: vertical, 1: horizontal
}

func getSection(segment string, startRow int, startCol int) wireSection {

	var ws wireSection

	// Get direction and distance
	sDirection := segment[0:1]
	distance, _ := strconv.Atoi(segment[1:])

	switch sDirection {
	case "U":
		ws.direction = 0
		ws.p.row = startRow + distance
		ws.p.col = startCol
	case "D":
		ws.direction = 0
		ws.p.row = startRow - distance
		ws.p.col = startCol
	case "L":
		ws.direction = 1
		ws.p.row = startRow
		ws.p.col = startCol - distance
	case "R":
		ws.direction = 1
		ws.p.row = startRow
		ws.p.col = startCol + distance
	}

	return ws
}

func between(val int, val1 int, val2 int) bool { // Is val between val1 and val2
	if val1 > val2 {
		if val >= val2 && val <= val1 {
			return true
		}
	} else {
		if val <= val2 && val >= val1 {
			return true
		}
	}
	return false
}

func getIntersection(w1_1 wireSection, w2_1 wireSection, w1_2 wireSection, w2_2 wireSection) (int, int) {
	if w1_2.direction == 0 {
		if between(w1_1.p.col, w2_1.p.col, w2_2.p.col) {
			if between(w2_1.p.row, w1_1.p.row, w1_2.p.row) {
				/*fmt.Println("Trouvé 1")
				fmt.Println(w1_1)
				fmt.Println(w2_1)
				fmt.Println(w1_2)
				fmt.Println(w2_2)
				fmt.Print(w2_1.p.row)
				fmt.Print(" - ")
				fmt.Println(w1_1.p.col)*/
				return w2_1.p.row, w1_1.p.col
			}
		}
	} else {
		if between(w2_2.p.col, w1_1.p.col, w1_2.p.col) {
			if between(w1_1.p.row, w2_1.p.row, w2_2.p.row) {
				/*fmt.Println("Trouvé 2")
				fmt.Println(w1_1)
				fmt.Println(w2_1)
				fmt.Println(w1_2)
				fmt.Println(w2_2)
				fmt.Print(w1_1.p.row)
				fmt.Print(" - ")
				fmt.Println(w2_1.p.col)*/
				return w1_1.p.row, w2_1.p.col
			}
		}
	}
	return 0, 0
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("day3.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	const totalSize = 301

	scanner := bufio.NewScanner(file)
	//	var tissu [1000][1000]int
	var wires [2][totalSize + 1]wireSection
	row := 0
	col := 0
	indexWire := 0
	wires[0][0].direction = 0
	wires[0][0].p.col = 0
	wires[0][0].p.row = 0
	wires[1][0].direction = 0
	wires[1][0].p.col = 0
	wires[1][0].p.row = 0

	// Parsing file
	for scanner.Scan() {
		row = 0
		col = 0
		index := 1
		myString := scanner.Text()
		pos := strings.Index(myString, ",")
		for pos != -1 {
			wires[indexWire][index] = getSection(myString[0:pos], row, col)
			myString = myString[pos+1:]
			row = wires[indexWire][index].p.row
			col = wires[indexWire][index].p.col
			pos = strings.Index(myString, ",")
			index++
		}
		// Last chunk
		wires[indexWire][index] = getSection(myString[0:], row, col)
		row = wires[indexWire][index].p.row
		col = wires[indexWire][index].p.col

		indexWire++
	}

	fmt.Println(wires[0])
	fmt.Println(wires[1])

	// Check cross points
	minDistance := 999999 // random high value
	distance := 0
	w1Length := 0
	w2Length := 0
	minWireLength := 999999 // random high value
	wireLength := 0

	for indexW1 := 1; indexW1 <= totalSize; indexW1++ {
		w2Length = 0
		// Check each segment of wire 1 with all segments of wire 2
		for indexW2 := 1; indexW2 <= totalSize; indexW2++ {
			fmt.Print("w2Length : ")
			fmt.Println(w2Length)
			// Check if diferent direction
			if wires[0][indexW1].direction != wires[1][indexW2].direction {
				row, col = getIntersection(wires[0][indexW1-1], wires[1][indexW2-1], wires[0][indexW1], wires[1][indexW2])
				distance = Abs(row) + Abs(col)
				//fmt.Println(distance)

				if distance != 0 {
					wireLength = w1Length + w2Length
					// Add distance between last point and intersection
					wireLength += Abs(wires[0][indexW1-1].p.col-col) + Abs(wires[0][indexW1-1].p.row-row)
					wireLength += Abs(wires[1][indexW2-1].p.col-col) + Abs(wires[1][indexW2-1].p.row-row)
					fmt.Print("Length : ")
					fmt.Print(w1Length)
					fmt.Print(" - ")
					fmt.Print(w2Length)
					fmt.Print(" - ")
					fmt.Println(wireLength)
					if wireLength < minWireLength {
						minWireLength = wireLength
					}
					if distance < minDistance {
						if row != 0 && col != 0 {
							fmt.Print("Distance : ")
							fmt.Print(row)
							fmt.Print(" - ")
							fmt.Print(col)
							fmt.Print(" - ")
							fmt.Println(distance)
						}

						minDistance = distance
					}

				}

			}
			w2Length += Abs(wires[1][indexW2].p.col-wires[1][indexW2-1].p.col) + Abs(wires[1][indexW2].p.row-wires[1][indexW2-1].p.row)
		}
		// Add distance
		w1Length += Abs(wires[0][indexW1].p.col-wires[0][indexW1-1].p.col) + Abs(wires[0][indexW1].p.row-wires[0][indexW1-1].p.row)

	}
	/*fmt.Println(goodCol)
	fmt.Println(goodRow)*/
	fmt.Println(minDistance)
	fmt.Println(minWireLength)
}
