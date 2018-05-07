package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Вывести через запятую ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше заданного m.
Входные параметры: длина и значение минимального квадрата
Выход: строка с рядом чисел
 */

type TaskSixContext struct {
	m      int
	length int
}

var context TaskSixContext

func taskSixMain() {
	taskSixInput()
	printTaskSixResult(getStartPoint())
}

func taskSixInput() {
	var bufferString string

	fmt.Print("Type length \"n\" : ")
	_, err := fmt.Scanf("%s", &bufferString)
	simpleErrorsChecker(err)

	//if float check
	context.length, err = strconv.Atoi(bufferString)
	int32InputChecker(context.length, err)

	fmt.Print("Type sqrt \"m\" : ")
	_, err = fmt.Scanf("%s", &bufferString)
	simpleErrorsChecker(err)

	context.m, err = strconv.Atoi(bufferString)
	int32InputChecker(context.m, err)
}

func getStartPoint() (startPoint int) {
	fmt.Println()
	sqrtOfM := getSqrt(context.m)
	fmt.Printf("The sqrt of m is %f", sqrtOfM)

	startPoint = int(math.Ceil(sqrtOfM))
	fmt.Printf("\nNumbers cant be lower than %d", startPoint)

	//for example, if the number is 100 then ceil ceil will return 10 which is sqrt of 100
	if int(sqrtOfM) == startPoint {
		startPoint++
	}
	return
}

func printTaskSixResult(startPoint int) {
	fmt.Printf("\nThe result is: \n")
	for i := 0; i < int(context.length); i++ {
		fmt.Print(startPoint)
		fmt.Print(" ")
		startPoint++
	}
}

func getSqrt(num int) (float64) {
	return math.Sqrt(float64(num))
}
