package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"encoding/json"
	//"strconv"
	"strconv"
)

/*
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано шестизначное число, и сумма первых трёх цифр равна сумме последних трёх, то этот билет считается счастливым.
2. Сложный — если сумма чётных цифр билета равна сумме нечётных цифр билета, то билет считается счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их большее количество на заданном интервале.

Входные параметры: объект context с полями min и max
Выход: объект с информацией о победившем методе и количестве счастливых билетов для каждого способа подсчёта.

 */

type ErrorResponse struct {
	Status string
	Reason string
}

type Context struct {
	Min int
	Max int
}

func main() {
	var context Context

	fmt.Print("Enter context, example:  {\"min\":320123,\"max\":320320}\n")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	err := json.Unmarshal([]byte(line), &context)
	if err != nil {
		errHandler(err)
		return
	}

	if context.Max <= 0 || context.Min <= 0 {
		errHandler(errors.New("Cant be 0 or signed"))
		return
	}
	if len(strconv.Itoa(context.Max)) < 6 || len(strconv.Itoa(context.Min)) < 6 {
		errHandler(errors.New("Cant have less than 6 numbers"))
		return
	} else if context.Min > context.Max {
		errHandler(errors.New("Min cant be higher than max"))
		return
	}

	//same length check
	if len(strconv.Itoa(int(context.Max))) != len(strconv.Itoa(int(context.Min))) {
		errHandler(errors.New("Must be the same lengths"))
		return
	}

	countEasy := easyWay(context.Min, context.Max)
	fmt.Println("Easy count : ", countEasy)
	countHard := hardWay(context.Min, context.Max)
	fmt.Println("Hard count : ", countHard)

}

func easyWay(min, max int) (count int) {
	//1. Простой — если на билете напечатано шестизначное число, и сумма первых трёх цифр равна сумме последних трёх, то этот билет считается счастливым.
	var isLucky bool
	//length:=len(strconv.Itoa(int(min)))
	for i := min; i <= max; i++ {
		isLucky = false
		digits := make([]int, 0)

		//getting digits
		currentNumber := i
		for currentNumber > 0 {
			digit := currentNumber % 10
			digits = append(digits, digit)
			currentNumber /= 10
		}

		var firstThreeSum, secondThreeSum int

		for j := 0; j < len(digits)/2; j++ {
			firstThreeSum += digits[j]
		}

		if len(digits)%2 == 0 {
			for j := len(digits) / 2; j < len(digits); j++ {
				secondThreeSum += digits[j]
			}
		} else {
			for j := (len(digits) / 2) + 1; j < len(digits); j++ {
				secondThreeSum += digits[j]
			}
		}

		if firstThreeSum == secondThreeSum {
			isLucky = true
		}
		if isLucky {
			count++
		}

	}
	return count
}

func hardWay(min, max int) (count int) {
	//2. Сложный — если сумма чётных цифр билета равна сумме нечётных цифр билета, то билет считается счастливым.
	var isLucky bool
	for i := min; i <= max; i++ {
		isLucky = false
		digits := make([]int, 0)
		var evenSum, oddSum int

		//getting digits
		currentNumber := i
		for currentNumber > 0 {
			digit := currentNumber % 10
			digits = append(digits, digit)
			currentNumber /= 10
		}
		for _, element := range digits {
			if element%2 == 0 {
				evenSum += element
			} else {
				oddSum += element
			}
		}
		if evenSum == oddSum {
			isLucky = true
		}
		if isLucky {
			count++
		}
	}
	return count
}

func errHandler(err error) {
	errResponse := &ErrorResponse{Status: "failed", Reason: err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}
