package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getFuel(masse int) int {
	val := int(math.Floor(float64(masse)/3.0)) - 2
	if val < 0 {
		return 0
	} else {
		return val + getFuel(val)
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
		val, _ := strconv.Atoi(scanner.Text())
		total = total + getFuel(val)

	}

	fmt.Print("Result : ")
	fmt.Println(strconv.Itoa(total))

}
