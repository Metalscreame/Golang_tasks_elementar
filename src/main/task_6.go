package main

import (
	"fmt"
	"math"
)

/*
Вывести через запятую ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше заданного m.
Входные параметры: длина и значение минимального квадрата
Выход: строка с рядом чисел
 */

const MESSAGE_SQRT = "number to get sqrt"

type TaskSixContext struct {
	m      int
	length int
}

var context TaskSixContext

func taskSixMain() {
	taskSixInput()
	fmt.Println()
	printTaskSixResult(getStartPoint())
}

func taskSixInput() {
	context.length = getValidScannedIntegerValue(MESSAGE_LENGTH)
	context.m = getValidScannedIntegerValue(MESSAGE_SQRT)
}

func getStartPoint() (startPoint int) {
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
