package main

import (
	"fmt"
	"strconv"
	"errors"
	"os"
)

/*

Проверить является ли число или его часть палиндромом. Например, число 4 не является палиндромом,
но является палиндромом его часть 3443. Числа меньшие, чем 10 палиндромом не считать.

Входные параметры: число
Выход: извлечённый из числа палиндром либо 0, если извлечение не удалось.

// поискать точку старта где i==i+1
потом от центра начать двигаться вбока 

 */

func taskFourMain() {
	var input int
	fmt.Print("Enter a number to if its a polindrome or not: ")

	var err error
	var buffer string
	_, err = fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err, "")

	//if float check
	input, err = strconv.Atoi(buffer)
	if err != nil || input <= 9 {
		errHandler(errors.New(ERROR_SIGNED))
		os.Exit(1)
	}

	srtNumber := strconv.Itoa(input)
	flag, res := isPalindrome(srtNumber)
	if flag == false {
		fmt.Println("0")
	} else {
		for _, el := range res {
			fmt.Print(el)
		}
	}
}

func isPalindrome(input string) (detectionFlag bool, res []string) {
	var equalsFlag bool

	var equalsIndex_i, equalsIndex_j, continueIndex int

	for continueIndex = 0; continueIndex < len(input); continueIndex++ {
		equalsFlag = false

		for i := continueIndex; i < len(input); i++ {
			for j := i + 1; j < len(input); j++ { //finds for equal number in the input
				if input[j] == input[i] {
					equalsFlag = true
					equalsIndex_j = j
					break
				}
			}
			if equalsFlag {
				equalsIndex_i = i
				break
			}
		}

		j := equalsIndex_j
		for i := equalsIndex_i; i <= equalsIndex_j; i++ { // if found checks if its a polindrome from i to j
			if input[i] == input[j] {
				detectionFlag = true
			} else {
				detectionFlag = false
				break
			}
			j--
		}
		if detectionFlag { // if the flag is still true - its a polindrome
			for i := equalsIndex_i; i < equalsIndex_j+1; i++ {
				res = append(res, string(input[i]))
			}
			break
		}
	}
	return detectionFlag, res

}
