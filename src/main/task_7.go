package main

import "errors"
import (
	"math"
	"fmt"
	"strconv"
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

func fibon(c ContextFibon) (res []int64, err error) {
	if c.length != 0 {
		if c.length < 2 {
			return nil, errors.New("Length can't be lower than 2")
		}
		res := make([]int64, c.length)
		res[0] = 0
		res[1] = 1
		for i := 2; i < c.length; i++ {
			res[i] = res[i-1] + res[i-2]
		}
		return res, nil
	} else if c.max != 0 { // min can be 0 but max - cant
		res := make([]int64, 0) // init zero size, it will be realocated while append
		from := math.Ceil(math.Log((float64(c.min)-0.5)*math.Sqrt(5)) / math.Log(math.Phi))
		to := math.Floor(math.Log((float64(c.max)+0.5)*math.Sqrt(5)) / math.Log(math.Phi))

		a := bineFunc(from)
		i := from + 1
		b := bineFunc(i)

		for i < to {
			res = append(res, int64(b))
			b += a
			a = b - a
			i++
		}
		return res, nil

	} else if c.min == c.max {
		return nil, errors.New("Start can't be the same as max")
	} else {
		return nil, errors.New(ERROR_WRONG_INPUT)
	}
}

func bineFunc(n float64) float64 {
	phi := 1 + math.Sqrt(5)
	return math.Round((math.Pow(phi, n) - math.Pow(-phi, -n)) / (2*phi - 1))
}

func main7() {
	var array []int64
	var c ContextFibon
	fmt.Print("\nChoose (length - type 1, limits - type 2): ")
	var firstInput, secondInput int
	var err error
	var bufferString string
	_, err = fmt.Scanf("%d", &firstInput)
	if err != nil {
		errHandler(err)
		return
	}

	if firstInput == LENGTH {
		fmt.Print("Enter length: ")
		_, err = fmt.Scanf("%s", &bufferString)
		if err != nil {
			errHandler(err)
			return
		}

		//if float check
		firstInput,err = strconv.Atoi(bufferString)
		if err != nil || firstInput <=0{
			errHandler(errors.New(ERROR_SIGNED))
			return
		}
		c.length = firstInput

	} else if firstInput == LIMITS {
		fmt.Print("Enter min point: ")
		_, err = fmt.Scanf("%s", &bufferString)
		if err != nil { // syntax with errors checking
			errHandler(err)
			return
		}

		firstInput,err = strconv.Atoi(bufferString)
		if err != nil || firstInput <=0{
			errHandler(errors.New(ERROR_SIGNED))
			return
		}

		fmt.Print("Enter max point: ")
		fmt.Scanf("%s", &bufferString)

		secondInput,err = strconv.Atoi(bufferString)
		if err != nil|| secondInput <=0{
			errHandler(errors.New("Wrong number"))
			return
		}

		c.min = firstInput
		c.max = secondInput
	} else {
		errHandler(errors.New("Wrong firstInput, you can choose only beetwen 1 or 2"))
		return
	}

	array, err = fibon(c)
	if err != nil {
		errHandler(err)
		return
	}

	fmt.Printf(" The numbers are: \n")
	for _, num := range array {
		fmt.Print(" ", num)
	}
	fmt.Printf("")
}

