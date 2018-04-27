package main

import (
	"fmt"
	"strconv"
	"errors"
	"os"
)

/*

Проверить является ли число или его часть палиндромом. Например, число 1234437 не является палиндромом,
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
	simpleErrorChecker(err, "")

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
	}else {
		for _, el := range res {
			fmt.Print(el)
		}
	}
}

func isPalindrome(input string) (detectionFlag bool, result []string) {
	if len(input)%2 == 0 {
		for i := 0; i < len(input)/2; i++ {
			if input[i] != input[len(input)-i-1] {
				detectionFlag = false
				continue
			} else {
				detectionFlag = true
				result = append(result, string(input[i]))
			}
		}
	} else {
		runes := []rune(input)
		var resultTemp []string
		substringed := string(runes[1:len(input)])
		for i := 0; i < len(substringed)/2; i++ {
			if substringed[i] != substringed[len(substringed)-i-1] {
				detectionFlag = false
				continue
			} else {
				detectionFlag = true
				result = append(result, string(substringed[i]))
			}
		}
		// we'll cut the last char and will see if there is a lonidrome or not
		substringed2 := string(runes[0:len(input)-1])

		for i := 0; i < len(substringed2)/2; i++ {
			if substringed2[i] != substringed2[len(substringed2)-i-1] {
				detectionFlag = false
				continue
			} else {
				detectionFlag = true
				resultTemp = append(resultTemp, string(substringed2[i]))
			}
		}

		if len(result) < len(resultTemp) {
			result = resultTemp
			detectionFlag = true
		} else if len(result) == 0 && len(resultTemp) == 0 {
		} else {
			detectionFlag = true //все проверки сделаны для случая, когда при нечетном числе
			// при втором сабстринге не показался ниодин палиндром и флаг не выставился
		}

	}

	//adding second half of the palondrome to the result string
	i := len(result)
	for i > 0 {
		result = append(result, string(result[i-1]))
		i--
	}
	return detectionFlag, result
}

func findCenter()  {
	
}