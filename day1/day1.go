package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calculeFuel(masse int) int {
	val := int(math.Floor(float64(masse)/3.0)) - 2
	if val < 0 {
		return 0
	} else {
		return val + calculeFuel(val)
	}
}

func main() {
	file, err := os.Open("day1.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		maString := scanner.Text()
		//fmt.Println(maString)

		val, _ := strconv.Atoi(maString)
		total = total + calculeFuel(val)

	}

	fmt.Println("Terminé !")
	fmt.Println(strconv.Itoa(total))

	/*if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}*/

}
