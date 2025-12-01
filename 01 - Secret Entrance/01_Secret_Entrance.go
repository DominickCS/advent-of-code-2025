package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotateLock(combo string, lockPOS int) int {
	rotationDIR := combo[0]

	comboToInt, _ := strconv.Atoi(combo[1:])

	if rotationDIR == 'L' {
		newLockPOS := lockPOS - comboToInt
		if newLockPOS < 0 {
			// fmt.Println("NEW LOCK POS", newLockPOS)
			newLockPOS = newLockPOS % 100
			newLockPOS += 100
			newLockPOS = newLockPOS % 100
			return newLockPOS
		}
		return newLockPOS
	} else if rotationDIR == 'R' {
		newLockPOS := lockPOS + comboToInt
		if newLockPOS >= 100 {
			newLockPOS = newLockPOS % 100
			return newLockPOS
		}
		return newLockPOS
	}
	return 0
}

func main() {
	var combos []string
	lock := 50
	zeroCount := 0

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
		lock = rotateLock(combos[i], lock)
		if lock == 0 {
			zeroCount += 1
		}
	}

	fmt.Println("zeroCount: ", zeroCount)
}
