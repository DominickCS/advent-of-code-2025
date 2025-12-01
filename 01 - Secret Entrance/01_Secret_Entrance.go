package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotateLock(combo string, lockPOS int, passedZero int) (int, int) {
	rotationDIR := combo[0]
	comboToInt, _ := strconv.Atoi(combo[1:])

	if rotationDIR == 'L' {
		fmt.Println("Rotating Left: ", comboToInt, "CURRENT POS: ", lockPOS)
		newLockPOS := lockPOS - comboToInt

		if newLockPOS == 0 {
			fmt.Println("HIT 0", passedZero)
			passedZero += 1
		}

		if newLockPOS < 0 {
			if len(combo[1:]) >= 3 {
				passedZero += (1 * (comboToInt / 100))
				fmt.Println("PASSED ZERO", (1 * (comboToInt / 100)), "Times.")
			}
			fmt.Println("PASSED ZERO", passedZero)
			passedZero += 1
			newLockPOS = newLockPOS % 100
			newLockPOS += 100
			newLockPOS = newLockPOS % 100
			return newLockPOS, passedZero
		}
		return newLockPOS, passedZero

	} else if rotationDIR == 'R' {
		fmt.Println("Rotating Right: ", comboToInt, "CURRENT POS: ", lockPOS)
		newLockPOS := lockPOS + comboToInt

		if newLockPOS == 0 {
			fmt.Println("HIT 0", passedZero)
			passedZero += 1
		}

		if newLockPOS >= 100 {
			if len(combo[1:]) >= 3 {
				passedZero += (1 * (comboToInt / 100))
				fmt.Println("PASSED ZERO", (1 * (comboToInt / 100)), "Times.")
			}
			fmt.Println("PASSED ZERO", passedZero)
			passedZero += 1
			newLockPOS = newLockPOS % 100
			return newLockPOS, passedZero
		}
		return newLockPOS, passedZero
	}
	return 0, 0
}

func main() {
	var combos []string
	lock := 50
	zeroCount := 0
	passedZero := 0

	// Open combination file
	f, err := os.Open("combination.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		combos = append(combos, scanner.Text())
	}

	for i := 0; i < len(combos); i++ {
		lock, passedZero = rotateLock(combos[i], lock, passedZero)
		if lock == 0 {
			zeroCount += 1
		}
	}

	fmt.Println("zeroCount: ", zeroCount)
	fmt.Println("passed zero: ", passedZero)
}
