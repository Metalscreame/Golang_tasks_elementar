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


 */

const LEAST_POSSIBLE_NUM_FOR_PALINDROME = 10
const THERE_IS_NO_PALINDROME_MSG = "0"

func taskFourMain() {
	printPalindromeResults(isPalindrome(palindromeInput()))
}
//io buffer faster concat
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

		if equalsFlag{
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
	}
	return detectionFlag, res
}

func palindromeInput() string {
	var input int
	var buffer string

	fmt.Print("Enter a number to find if its a polindrome or not: ")
	_, err := fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err)

	//if float check
	input, err = strconv.Atoi(buffer)
	if err != nil || input < LEAST_POSSIBLE_NUM_FOR_PALINDROME {
		errHandler(errors.New(ERROR_SIGNED))
		os.Exit(ERROR_CODE)
	}

	return buffer //after checking convert it back to str to work
}

func printPalindromeResults(palindromeIsHere bool, res []string) {
	if !palindromeIsHere {
		fmt.Println(THERE_IS_NO_PALINDROME_MSG)
	} else {
		for _, el := range res {
			fmt.Print(el)// better o concat and print once
		}
	}
}
