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

 var m,length int

func taskSixMain() {
	taskSixInput()
	printTaskSixResult(getStartPoint())
}

func taskSixInput()  {
	var bufferString string

	fmt.Print("Type length \"n\" : ")
	_, err := fmt.Scanf("%s", &bufferString)
	simpleErrorsChecker(err)

	//if float check
	length, err = strconv.Atoi(bufferString)
	int32InputChecker(length, err)

	fmt.Print("Type sqrt \"m\" : ")
	_, err = fmt.Scanf("%s", &bufferString)
	simpleErrorsChecker(err)

	m, err = strconv.Atoi(bufferString)
	int32InputChecker(m, err)

	if m <= 2 {
		simpleErrorsChecker(errors.New("There are no numbers sqrt of which isnt less them m"))
	}
}

func getStartPoint() int {
	sqrtOfM :=getSqrt(m)
	fmt.Printf("The sqrt of m is %f", sqrtOfM)

	startPoint := int(math.Ceil(sqrtOfM))
	fmt.Printf("\nNumbers cant be lower than %d", startPoint)

	//for example, if the number is 100 then ceil ceil will return 10 which is sqrt of 100
	if int(sqrtOfM) == startPoint {
		startPoint++
	}
	return startPoint
}

func printTaskSixResult(startPoint int)  {
	fmt.Printf("\nThe result is: \n")
	for i := 0; i < int(length); i++ {
		fmt.Print(startPoint)
		fmt.Print(" ")
		startPoint++
	}
}

func getSqrt(num int) (float64)  {
	return math.Sqrt(float64(m))
}