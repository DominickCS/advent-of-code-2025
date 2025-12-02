package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotateLock(combination string, initLockPOS int, passZeroCount int) (int, int) {
	direction := combination[0]
	convertedCombo, _ := strconv.Atoi(combination[1:])

	passZeroCount += (1 * (convertedCombo / 100))

	if direction == 'L' {
		// fmt.Println("CURRENT POS:", initLockPOS, "ROTATING LEFT:", convertedCombo, "TICKS")
		newLockPOS := initLockPOS - (convertedCombo % 100)

		if newLockPOS <= 0 {
			if initLockPOS != 0 {
				passZeroCount += 1
				newLockPOS = newLockPOS % 100 // IF BELOW RANGE
				newLockPOS += 100             // IF NEGATIVE
				newLockPOS = newLockPOS % 100 // IF AT 0 or SIMILAR
			}
		}

		newLockPOS = newLockPOS % 100 // IF BELOW RANGE
		newLockPOS += 100             // IF NEGATIVE
		newLockPOS = newLockPOS % 100 // IF AT 0 or SIMILAR

		return newLockPOS, passZeroCount

	} else if direction == 'R' {
		newLockPOS := initLockPOS + (convertedCombo % 100)

		if newLockPOS >= 100 {
			if initLockPOS != 0 {
				passZeroCount += 1
			}

			newLockPOS = newLockPOS % 100

		}

		return newLockPOS, passZeroCount
	}

	return 0, 0
}

func main() {
	var combinationList []string
	initLockPOS := 50
	zeroCount := 0
	passZeroCount := 0

	// Open combination file
	f, err := os.Open("combination.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		combinationList = append(combinationList, scanner.Text())
	}

	for i := 0; i < len(combinationList); i++ {
		initLockPOS, passZeroCount = rotateLock(combinationList[i], initLockPOS, passZeroCount)

		if initLockPOS == 0 {
			zeroCount += 1
		}
	}

	fmt.Println("zeroCount: ", zeroCount)
	fmt.Println("Passed Zero", passZeroCount, "Times.")
}
