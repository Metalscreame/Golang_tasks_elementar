package main

import "fmt"

/*
Вывести шахматную доску с заданными размерами высоты и ширины, по принципу:
*  *  *  *  *  *
  *  *  *  *  *  *
*  *  *  *  *  *
  *  *  *  *  *  *

Входные параметры: длина, ширина, символ для отображения.
Выход: строка с представлением шахматной доски

 */

func main() {
	var symbolToPrint string
	var widt, length int

	fmt.Printf("Enter the symbol to print and press Enter: ")
	_, err := fmt.Scanf("%s", &symbolToPrint)
	if err != nil {
		typeErrHandler(err)
	}

	fmt.Printf("Enter the width and press Enter: ")
	_, err2 := fmt.Scanf("%d", &widt)
	if err2 != nil {
		typeErrHandler(err2)
	}
	fmt.Printf("Enter the length and press Enter: ")
	_, err3 := fmt.Scanf("%d", &length)
	if err3 != nil {
		typeErrHandler(err3)
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

func typeErrHandler(err error) {
	var errorString string
	fmt.Printf(err.Error())
	fmt.Printf("\n")
	fmt.Scanln(&errorString) // for clearing the buffer
	main()
}
