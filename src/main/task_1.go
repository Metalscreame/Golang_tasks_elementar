package main

import (
	"fmt"
	"errors"
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

const SPACE_SYMBOL = " "
const MESSAGE_LENGTH = "the length"
const MESSAGE_WIDTH  = "the width"
const MESSAGE_SYMBOL  = "the width"

func taskOneMain() {
	printBoard(taskOneInput())
}

func taskOneInput() (symbolToPrint string, width, length int) {
	symbolToPrint = getValidScannedStringValue(MESSAGE_SYMBOL)
	validateSymbolToPrint(symbolToPrint)
	length = getValidIntegerFromString(getValidScannedStringValue(MESSAGE_LENGTH))
	width = getValidIntegerFromString(getValidScannedStringValue(MESSAGE_WIDTH))
	return
}

func printBoard(symbolToPrint string, width, length int) {
	fmt.Printf("\n\n")
	for i := 1; i < width+1; i++ { //starts with 1 to print symbol first
		if i%2 == 0 {
			fmt.Print(SPACE_SYMBOL)
		}
		for j := 0; j < length; j++ {
			if j%2 == 0 {
				fmt.Print(symbolToPrint)
			} else {
				fmt.Print(SPACE_SYMBOL)
			}
		}
		fmt.Println()
	}
}

func validateSymbolToPrint(s string) {
	if len(s) > 1 {//длина в рунах
		simpleErrorsChecker(errors.New(ERROR_ONE_SYMBOL))
	}
}