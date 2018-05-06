package main

import (
	"fmt"
	"errors"
	"strconv"
)

/*
Вывести шахматную доску с заданными размерами высоты и ширины, по принципу:
*  *  *  *  *  *
  *  *  *  *  *  *
*  *  *  *  *  *
  *  *  *  *  *  *

Входные параметры: длина, ширина, символ для отображения.
Выход: строка с представлением шахматной доски

 */

func taskOneMain() {
	printBoard(taskOneInput())
}

func taskOneInput() (symbolToPrint string, widt, length int) {
	var buffer string

	fmt.Print("Enter one symbol to print and press Enter: ")
	_, err := fmt.Scanf("%s", &symbolToPrint)
	simpleErrorsChecker(err)

	if len(symbolToPrint) > 1 {
		simpleErrorsChecker(errors.New(ERROR_ONE_SYMBOL))
	}

	fmt.Printf("Enter the length and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err)

	length, err = strconv.Atoi(buffer)
	int32InputChecker(length, err)

	fmt.Printf("Enter the width and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err)

	widt, err = strconv.Atoi(buffer)
	int32InputChecker(widt, err)
	return
}

func printBoard(symbolToPrint string, width, length int) {
	fmt.Printf("\n\n")
	for i := 1; i < width+1; i++ { //starts with 1 to print symbol first
		if i%2 == 0 {
			fmt.Print(" ")
		}
		for j := 0; j < length; j++ {
			if j%2 == 0 {
				fmt.Print(symbolToPrint)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
