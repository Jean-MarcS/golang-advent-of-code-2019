package main

import (
	"fmt"
)

// Both following functions are from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + start
	}
	return xs
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func extractParameters(myList []int, index int) (int, int, int, int) {
	instruction := myList[index]

	opcode := instruction % 100
	instruction = instruction / 100
	param1Mode := instruction % 10
	instruction = instruction / 10
	param2Mode := instruction % 10
	instruction = instruction / 10

	param1 := 0
	param2 := 0
	param3 := 0

	if index+1 < len(myList) {
		if param1Mode == 0 {
			if myList[index+1] < len(myList) {
				param1 = myList[myList[index+1]]
			} else {
				param1 = 0
			}
		} else {
			param1 = myList[index+1]
		}
		if index+2 < len(myList) {
			if param2Mode == 0 {
				if myList[index+2] < len(myList) {
					param2 = myList[myList[index+2]]
				} else {
					param2 = 0
				}
			} else {
				param2 = myList[index+2]
			}
			if index+3 < len(myList) {
				param3 = myList[index+3]
			}
		}
	}

	return opcode, param1, param2, param3

}

func runProgram(myList []int, index int, val1 int, val2 int) (returnValue int, indexOut int, finished bool) {

	returnValue = 0
	finished = false

	mustContinue := true
	input1used := !(index == 0)
	total := 0

	for mustContinue {

		opcode, param1, param2, param3 := extractParameters(myList, index)
		switch opcode {
		case 99:
			mustContinue = false
			finished = true
			break
		case 1:
			total = param1 + param2
			myList[myList[index+3]] = total
			index += 4
			break
		case 2:
			total = param1 * param2
			myList[myList[index+3]] = total
			index += 4
			break
		case 3:
			if input1used {
				myList[myList[index+1]] = val2
			} else {
				myList[myList[index+1]] = val1
				input1used = true
			}
			index += 2
			break
		case 4:
			mustContinue = false
			returnValue = param1
			index += 2
			break
		case 5:
			if param1 != 0 {
				index = param2
			} else {
				index += 3
			}
			break
		case 6:
			if param1 == 0 {
				index = param2
			} else {
				index += 3
			}
			break
		case 7:
			if param1 < param2 {
				myList[param3] = 1
			} else {
				myList[param3] = 0
			}
			index += 4
			break
		case 8:
			if param1 == param2 {
				myList[param3] = 1
			} else {
				myList[param3] = 0
			}
			index += 4
			break
		}
	}
	return returnValue, index, finished

}

func main() {

	myList := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 67, 88, 105, 114, 195, 276, 357, 438, 99999, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 1001, 9, 2, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 101, 2, 9, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 4, 9, 1002, 9, 4, 9, 101, 2, 9, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 3, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99}

	// Part 1
	maxThrusterP1 := 0
	bestSettingP1 := 0
	phasesP1 := permutation(rangeSlice(0, 5))
	for _, phase := range phasesP1 {
		currentValue := 0
		for pass := 0; pass < 5; pass++ {
			// Run program with a fresh copy of the instructions
			currentValue, _, _ = runProgram(append([]int(nil), myList...), 0, phase[pass], currentValue)
		}
		if currentValue > maxThrusterP1 {
			maxThrusterP1 = currentValue
			bestSettingP1 = phase[4] + (phase[3] * 10) + (phase[2] * 100) + (phase[1] * 1000) + (phase[0] * 10000)
		}
	}

	fmt.Print("Part 1 result\nMax thruster : ")
	fmt.Println(maxThrusterP1)
	fmt.Print("Phase setting sequence : ")
	fmt.Println(bestSettingP1)

	// Part 2
	maxThrusterP2 := 0
	bestSettingP2 := 0
	var amplifiers [5][]int
	indexes := [5]int{0, 0, 0, 0, 0}
	phasesP2 := permutation(rangeSlice(5, 10))
	for _, phase := range phasesP2 {
		currentValue := 0
		finished := false
		// Init amplifiers
		for i := 0; i < 5; i++ {
			amplifiers[i] = append([]int(nil), myList...)
			indexes[i] = 0
		}
		for pass := 0; !(pass == 5 && finished); pass++ {
			if pass == 5 {
				pass = 0
			}
			// Run program with a fresh copy of the instructions
			currentValue, indexes[pass], finished = runProgram(amplifiers[pass], indexes[pass], phase[pass], currentValue)
			if currentValue > maxThrusterP2 {
				maxThrusterP2 = currentValue
				bestSettingP2 = phase[4] + (phase[3] * 10) + (phase[2] * 100) + (phase[1] * 1000) + (phase[0] * 10000)
			}
		}
	}

	fmt.Print("Part 2 result\nMax thruster : ")
	fmt.Println(maxThrusterP2)
	fmt.Print("Phase setting sequence : ")
	fmt.Println(bestSettingP2)
}
