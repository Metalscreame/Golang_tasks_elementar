package main

import (
	"fmt"
	"encoding/json"
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
const WRONG_NUMBER  = "Wrong number"

type ErrorResponse struct {
	Status string
	Reason string
}

func main() {
	var symbolToPrint string
	var widt, length int
	var buffer string

	var err error
	fmt.Printf("Enter the symbol to print and press Enter: ")
	_, err = fmt.Scanf("%s", &symbolToPrint)
	if err != nil {
		errHandler(err)
	}else if len(symbolToPrint)>1{
		err:= errors.New("There can be only one symbol")
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
	if err != nil{
		errHandler(errors.New(WRONG_NUMBER))
		return
	}


	fmt.Printf("Enter the length and press Enter: ")
	_, err = fmt.Scanf("%s", &buffer)
	if err != nil {
		errHandler(err)
		return
	}
	//if float check
	length,err = strconv.Atoi(buffer)
	if err != nil{
		errHandler(errors.New(WRONG_NUMBER))
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

func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}
