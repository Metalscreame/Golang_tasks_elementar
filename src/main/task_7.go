package main

import "errors"
import (
	"math"
	"fmt"
	"strconv"
	"os"
)

/*
7.	Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи, которые удовлетворяют переданному в функцию ограничению:
находятся в указанном диапазоне, либо имеют указанную длину.

Входные параметры: объект ContextFibon с полями min и max, либо с полем length
Выход: массив чисел
 */

const (
	LENGTH = 1
	LIMITS = 2
)

type ContextFibon struct {
	min    int
	max    int
	length int
}

func taskSevenMain() {
	var array []int
	var c ContextFibon
	fmt.Print("\nChoose (length - type 1, limits - type 2): ")
	var firstInput, secondInput int
	var err error
	var bufferString string
	_, err = fmt.Scanf("%d", &firstInput)
	if err != nil {
		errHandler(err)
		os.Exit(1)
	}

	if firstInput == LENGTH {
		fmt.Print("Enter length: ")
		_, err = fmt.Scanf("%s", &bufferString)
		if err != nil {
			errHandler(err)
			os.Exit(1)
		}

		//if float check
		firstInput, err = strconv.Atoi(bufferString)
		if err != nil || firstInput <= 0 {
			errHandler(errors.New(ERROR_SIGNED))
			os.Exit(1)
		}
		c.length = firstInput

	} else if firstInput == LIMITS {
		fmt.Print("Enter min point: ")
		_, err = fmt.Scanf("%s", &bufferString)
		if err != nil { // syntax with errors checking
			errHandler(err)
			os.Exit(1)
		}

		firstInput, err = strconv.Atoi(bufferString)
		if err != nil || firstInput <= 0 {
			errHandler(errors.New(ERROR_SIGNED))
			os.Exit(1)
		}

		fmt.Print("Enter max point: ")
		fmt.Scanf("%s", &bufferString)

		secondInput, err = strconv.Atoi(bufferString)
		if err != nil || secondInput <= 0 {
			errHandler(errors.New("Wrong number"))
			os.Exit(1)
		}

		c.min = firstInput
		c.max = secondInput
	} else {
		errHandler(errors.New("Wrong firstInput, you can choose only beetwen 1 or 2"))
		os.Exit(1)
	}

	array, err = fibon(c)
	if err != nil {
		errHandler(err)
		os.Exit(1)
	}

	fmt.Printf(" The numbers are: \n")
	for _, num := range array {
		fmt.Print(" ", num)
	}
	fmt.Printf("")
}


func fibon(c ContextFibon) (res []int, err error) {
	if c.length != 0 {
		if c.length < 2 {
			return nil, errors.New("Length can't be lower than 2")
		}
		res := make([]int, c.length)
		res[0] = 0
		res[1] = 1
		for i := 2; i < c.length; i++ {
			res[i] = res[i-1] + res[i-2]
		}
		return res, nil
	} else if c.max != 0 { // if it goes here, then its mode 2

		res := make([]int, 0)
		//var fibon [] int64

		var current, previous, prePrevious int
		for i := 0; ; i++ {
			if i == 0 {
				current=0
			}else if i == 1{
				previous=1
				current=1
			}else {
				current = previous + prePrevious
				prePrevious = previous
				previous = current
			}
			if current >= c.max {
				break
			}
			if current >= c.min {
				res = append(res, current)
			}
		}
		if len(res) == 0 {
			return res, errors.New("There are 0 numbers at this limits")
		}
		return res, nil

	} else if c.min == c.max {
		return nil, errors.New("Start can't be the same as max")
	} else {
		return nil, errors.New(ERROR_WRONG_INPUT)
	}
}