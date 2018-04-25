package main

import (
	"fmt"
	"math"
	"strconv"
	"errors"
)

/*
Вывести через запятую ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше заданного m.

Входные параметры: длина и значение минимального квадрата
Выход: строка с рядом чисел

 */


func main6() {

	var length, m int
	var bufferString string
	var err error
	fmt.Print("Type length \"n\" : ")
	_, err = fmt.Scanf("%s", &bufferString)
	if err != nil {
		errHandler(err)
		return
	}

	//if float check
	length,err= strconv.Atoi(bufferString)
	if err != nil || length<=0{
		errHandler(errors.New("Must be unsigned integer,cant be 0"))
		return
	}

	fmt.Print("Type sqrt \"m\" : ")
	_, err = fmt.Scanf("%s", &bufferString)
	if err != nil {
		errHandler(err)
		return
	}
	//number check
	m,err= strconv.Atoi(bufferString)
	if err != nil || m<=0{
		errHandler(errors.New("Must be unsigned integer and cant be 0"))
		return
	}

	if m <=2{
		errHandler(errors.New("There are no numbers sqrt of which isnt less them m"))
		return
	}


	startPointFl := math.Ceil(math.Sqrt(float64(m)))
	fmt.Printf("The sqrt of m is %f", math.Sqrt(float64(m)))
	startPoint := int(startPointFl)
	fmt.Printf("\nNumbers cant be lower than %d", startPoint)

	//for example, if the number is 100 then ceil ceil will return 10 which is sqrt of 100
	if int(math.Sqrt(float64(m))) == startPoint {
		startPoint++
	}

	fmt.Printf("\nThe result is: \n")
	for i := 0; i < int(length)-1; i++ {
		if startPoint == m {
			fmt.Printf("\n\nNumbers cant be printed anymore, because numbers has reached m\n")
			break
		}
		fmt.Print(startPoint)
		fmt.Print(",")
		startPoint++
	}
	fmt.Print(startPoint)
}

