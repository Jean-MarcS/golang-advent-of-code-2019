package main

import (
	"fmt"
)

func main() {

	totalPwdPart1 := 0
	totalPwdPart2 := 0
	for i := 387638; i <= 919123; i++ {
		pairFound := false
		pairs := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		number := i
		cm := number / 100000
		number -= cm * 100000
		dm := number / 10000

		if dm >= cm {
			if cm == dm {
				pairFound = true
				pairs[cm]++
			}
			number -= dm * 10000
			m := number / 1000
			if m >= dm {
				if dm == m {
					pairFound = true
					pairs[dm]++
				}
				number -= m * 1000
				c := number / 100
				if c >= m {
					if m == c {
						pairFound = true
						pairs[m]++
					}
					number -= c * 100
					d := number / 10
					if d >= c {
						if c == d {
							pairFound = true
							pairs[c]++
						}
						u := number - (d * 10)
						if u >= d {
							if d == u {
								pairFound = true
								pairs[d]++
							}

							if pairFound {
								// Add value to total part1
								totalPwdPart1++
								// Check if there's a "good" pair
								pairOK := false
								for j := 0; j < 10; j++ {
									if pairs[j] == 1 {
										pairOK = true
									}
								}
								if pairOK {
									totalPwdPart2++
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(totalPwdPart1)
	fmt.Println(totalPwdPart2)
}
