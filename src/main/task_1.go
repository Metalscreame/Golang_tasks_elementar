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



func main1() {
	var symbolToPrint string
	var widt, length int
	var buffer string

	var err error
	fmt.Print("Enter one symbol to print and press Enter: ")
	_, err = fmt.Scanf("%s", &symbolToPrint)
	if err != nil {
		errHandler(err)
	}else if len(symbolToPrint)>1{
		err:= errors.New(ERROR_ONE_SYMBOL)
		errHandler(err)
		return
	}

	fmt.Printf("Enter the width and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	if err != nil {
		errHandler(err)
		return
	}
	//if float check
	widt,err = strconv.Atoi(buffer)
	if err != nil || widt<=0{
		errHandler(errors.New(ERROR_SIGNED))
		return
	}

	fmt.Printf("Enter the length and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	if err != nil {
		errHandler(err)
		return
	}

	length,err = strconv.Atoi(buffer)
	if err != nil || length<=0{
		errHandler(errors.New(ERROR_SIGNED))
		return
	}

	fmt.Printf("\n\n")

	for i := 1; i < widt+1; i++ {
		if i%2 == 0 {
			fmt.Print(" ")
		}
		for j := 0; j < length; j++ {
			fmt.Print(symbolToPrint)
		}
		fmt.Print("\n")
	}
}

