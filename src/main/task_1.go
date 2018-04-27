package main

import (
	"fmt"
	"errors"
	"strconv"
	"os"
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
	var symbolToPrint string
	var widt, length int
	var buffer string
	var err error

	fmt.Print("Enter one symbol to print and press Enter: ")
	_, err = fmt.Scanf("%s", &symbolToPrint)
	simpleErrorChecker(err, "")

	if len(symbolToPrint) > 1 {
		errHandler(errors.New(ERROR_ONE_SYMBOL))
		os.Exit(1)
	}

	fmt.Printf("Enter the length and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	simpleErrorChecker(err, "")

	length, err = strconv.Atoi(buffer)
	if err != nil || length <= 0 {
		errHandler(errors.New(ERROR_SIGNED))
		os.Exit(1)
	}

	fmt.Printf("Enter the width and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	simpleErrorChecker(err, "")

	//if float check
	widt, err = strconv.Atoi(buffer)
	if err != nil || widt <= 0 {
		errHandler(errors.New(ERROR_SIGNED))
		os.Exit(1)
	}

	fmt.Printf("\n\n")
	for i := 1; i < widt+1; i++ {
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
		fmt.Print("\n")
	}
}
