package main

import "errors"
import (
	"math"
	"fmt"
	"encoding/json"
	"strconv"
)

/*
7.	Ряд Фибоначчи для диапазона
Вывести все числа Фибоначчи, которые удовлетворяют переданному в функцию ограничению:
находятся в указанном диапазоне, либо имеют указанную длину.

Входные параметры: объект Context с полями min и max, либо с полем length
Выход: массив чисел
 */
type ErrorResponse struct {
	Status string
	Reason string
}

type Context struct {
	min    int
	max    int
	length int
}

func fibon(c Context) (res []int64, err error) {
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
		return nil, errors.New("Input data is wrong")
	}
}

func bineFunc(n float64) float64 {
	phi := 1 + math.Sqrt(5)
	return math.Round((math.Pow(phi, n) - math.Pow(-phi, -n)) / (2*phi - 1))
}

func main() {
	var array []int64
	var c Context
	fmt.Print("\nChoose (length - type 1, limits - type 2): ")
	var i, j int
	var err error
	var bufferString string
	_, err = fmt.Scanf("%d", &i)
	if err != nil {
		errHandler(err)
		return
	}

	if i == 1 {
		fmt.Print("Enter length: ")
		_, err = fmt.Scanf("%s", &bufferString)
		if err != nil {
			errHandler(err)
			return
		}

		//if float check
		i,err = strconv.Atoi(bufferString)
		if err != nil{
			errHandler(errors.New("Wrong number"))
			return
		}
		c.length = i

	} else if i == 2 {
		fmt.Print("Enter min point: ")
		_, err = fmt.Scanf("%d", &i)
		if err != nil { // syntax with errors checking
			errHandler(err)
			return
		}
		fmt.Print("Enter max point: ")
		fmt.Scanf("%d", &j)

		c.min = i
		c.max = j
	} else {
		errHandler(errors.New("Wrong input, u cant type only 1 or 2"))
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
func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

