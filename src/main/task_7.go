package main

import "errors"
import (
	"fmt"
	"strconv"
)

/*
7.	Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи, которые удовлетворяют переданному в функцию ограничению:
находятся в указанном диапазоне, либо имеют указанную длину.

Входные параметры: объект ContextFibonacci с полями min и max, либо с полем length
Выход: массив чисел
 */

const (
	LENGTH          = 1
	LIMITS          = 2
	LENGTH_MODE_KEY = "Length"
	LIMIT_MODE_KEY  = "Limits"
)

type ContextFibonacci struct {
	min    int
	max    int
	length int
	result []int
}

func taskSevenMain() {
	var c ContextFibonacci
	c.fibonInput()
	c.getFibonacNumbers()
	c.printFiboncNums()
}

func (c *ContextFibonacci) fibonInput() {
	var firstInput, secondInput int
	var bufferString string

	fmt.Print("\nChoose (length - type 1, limits - type 2): ")
	_, err := fmt.Scanf("%d", &firstInput)
	simpleErrorsChecker(err)

	if firstInput == LENGTH {
		fmt.Print("Enter length: ")
		_, err = fmt.Scanf("%s", &bufferString)
		simpleErrorsChecker(err)

		//if float check
		firstInput, err = strconv.Atoi(bufferString)
		int32InputChecker(firstInput, err)

		c.length = firstInput

	} else if firstInput == LIMITS {
		fmt.Print("Enter min point: ")
		_, err = fmt.Scanf("%s", &bufferString)
		simpleErrorsChecker(err)

		firstInput, err = strconv.Atoi(bufferString)
		int32InputChecker(firstInput, err)

		fmt.Print("Enter max point: ")
		fmt.Scanf("%s", &bufferString)

		secondInput, err = strconv.Atoi(bufferString)
		int32InputChecker(secondInput, err)

		c.min = firstInput
		c.max = secondInput
	} else {
		simpleErrorsChecker(errors.New("Wrong firstInput, you can choose only beetwen 1 or 2"))
	}
}

func (c *ContextFibonacci) getFibonacNumbers() {
	if c.length != 0 {
		c.calculateFibonacNumbers(LENGTH_MODE_KEY)
	} else if c.max != 0 { // if it goes here, then its mode 2
		c.calculateFibonacNumbers(LIMIT_MODE_KEY)
		if len(c.result) == 0 {
			simpleErrorsChecker(errors.New("There are 0 numbers at this limits"))
		}
	} else if c.min == c.max {
		simpleErrorsChecker(errors.New("Start can't be the same as max"))
	} else {
		simpleErrorsChecker(errors.New(ERROR_WRONG_INPUT))
	}
}

func (c *ContextFibonacci) calculateFibonacNumbers(calculateModeKey string) {
	var current, previous, prePrevious int

	for i := 0; ; i++ {
		if i == 0 {
			current = 0
		} else if i == 1 {
			previous = 1
			current = 1
		} else {
			current = previous + prePrevious
			prePrevious = previous
			previous = current
		}

		if calculateModeKey == LIMIT_MODE_KEY {
			if current > c.max {
				break
			}
			if current > c.min {
				c.result = append(c.result, current)
			}
		} else if calculateModeKey == LENGTH_MODE_KEY {
			if len(strconv.Itoa(current)) < c.length {
			} else if len(strconv.Itoa(current)) > c.length {
				break
			} else {
				c.result = append(c.result, current)
			}
		}
	}
}

func (c *ContextFibonacci) printFiboncNums() {
	fmt.Printf(" The numbers are: \n")
	for _, num := range c.result {
		fmt.Print(" ", num)
	}
}
