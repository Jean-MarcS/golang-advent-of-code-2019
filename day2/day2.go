package main

import (
	"fmt"
)

func computeList(myList [121]int, val1 int, val2 int) int {
	continuer := true
	index := 0
	total := 0
	myList[1] = val1
	myList[2] = val2
	for continuer {
		opcode := myList[index]
		switch opcode {
		case 99:
			continuer = false
			break
		case 1:
			total = myList[myList[index+1]] + myList[myList[index+2]]
			myList[myList[index+3]] = total
			break
		case 2:
			total = myList[myList[index+1]] * myList[myList[index+2]]
			myList[myList[index+3]] = total
			break
		}
		index += 4
	}

	return myList[0]
}

func main() {

	myList := [121]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 10, 23, 2, 13, 23, 27, 1, 5, 27, 31, 2, 6, 31, 35, 1, 6, 35, 39, 2, 39, 9, 43, 1, 5, 43, 47, 1, 13, 47, 51, 1, 10, 51, 55, 2, 55, 10, 59, 2, 10, 59, 63, 1, 9, 63, 67, 2, 67, 13, 71, 1, 71, 6, 75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 9, 87, 1, 6, 87, 91, 2, 91, 6, 95, 1, 95, 6, 99, 2, 99, 13, 103, 1, 6, 103, 107, 1, 2, 107, 111, 1, 111, 9, 0, 99, 2, 14, 0, 0}

	fmt.Print("Part 1 result : ")
	fmt.Println(computeList(myList, 12, 2))

	for noun := 99; noun >= 0; noun-- {
		for verb := 99; verb >= 0; verb-- {
			if computeList(myList, noun, verb) == 19690720 {
				fmt.Print("Part 2 result : ")
				fmt.Println((noun * 100) + verb)
				// Change conditions to exit
				noun = -1
				break
			}

		}
	}
}
