package main

import (
	"fmt"
	"errors"
)

/*

Проверить является ли число или его часть палиндромом. Например, число 4 не является палиндромом,
но является палиндромом его часть 3443. Числа меньшие, чем 10 палиндромом не считать.

Входные параметры: число
Выход: извлечённый из числа палиндром либо 0, если извлечение не удалось.


 */

const LEAST_POSSIBLE_NUM_FOR_PALINDROME = 10
const THERE_IS_NO_PALINDROME_MSG = "0"
const MESSAGE_ENTER_PALINDROME  =  "a number to find if there is a palindrome or not"

func taskFourMain() {
	printPalindromeResults(isPalindrome(palindromeInput()))
}

func palindromeInput() string {
	var input int
	buffer := getValidScannedStringValue(MESSAGE_ENTER_PALINDROME)
	input = getValidIntegerFromString(buffer)

	if input < LEAST_POSSIBLE_NUM_FOR_PALINDROME {
		simpleErrorsChecker(errors.New(ERROR_SIGNED))
	}
	return buffer //after checking convert it back to str to work
}

func printPalindromeResults(palindromeIsHere bool, res []string) {
	if !palindromeIsHere {
		fmt.Println(THERE_IS_NO_PALINDROME_MSG)
	} else {
		for _, el := range res {
			fmt.Print(el) // better o concat and print once
		}
	}
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

		if equalsFlag {
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