package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"encoding/json"
	"strconv"
)

/*
Есть 2 способа подсчёта счастливых билетов:
1. Простой — если на билете напечатано шестизначное число, и сумма первых трёх цифр равна сумме последних трёх, то этот
билет считается счастливым.
2. Сложный — если сумма чётных цифр билета равна сумме нечётных цифр билета, то билет считается счастливым.
Определить программно какой вариант подсчёта счастливых билетов даст их большее количество на заданном интервале.

Входные параметры: объект context с полями min и max
Выход: объект с информацией о победившем методе и количестве счастливых билетов для каждого способа подсчёта.

 */

type Context struct {
	Min              int
	Max              int
	CountEasyFormula int
	CountHardFormula int
}

func taskFiveMain() {
	result, _ := json.Marshal(getContext())
	fmt.Printf("%s", result)
}

func getContext() (context Context) {
	context = taskFiveInput()
	context.validateInputTaskFive()
	context.easyWay()
	context.hardWay()
	return
}

func taskFiveInput() (context Context) {

	fmt.Print("Enter context, example:  {\"min\":320123,\"max\":320320}\n")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	err := json.Unmarshal([]byte(line), &context)
	simpleErrorsChecker(err)
	return
}

func (c *Context) easyWay() {
	var isLucky bool
	var firstThreeSum, secondThreeSum int

	for i := c.Min; i <= c.Max; i++ {
		isLucky = false
		firstThreeSum=0
		secondThreeSum=0

		digits := getSeparatedDigitsFromNumber(i)
		for j := 0; j < 3; j++ {
			firstThreeSum += digits[j]
		}
		for j := len(digits) - 3; j < len(digits); j++ {
			secondThreeSum += digits[j]
		}
		if firstThreeSum == secondThreeSum {
			isLucky = true
		}
		if isLucky {
			c.CountEasyFormula++
		}
	}
}

func (c *Context) hardWay() {
	var isLucky bool
	var evenSum, oddSum int
	for i := c.Min; i <= c.Max; i++ {
		isLucky = false
		evenSum=0
		oddSum=0
		digits := getSeparatedDigitsFromNumber(i)

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
			c.CountHardFormula++
		}
	}
}

func (c *Context) validateInputTaskFive() {
	if c.Max <= 0 || c.Min <= 0 {
		simpleErrorsChecker(errors.New(ERROR_SIGNED))
	}
	if len(strconv.Itoa(c.Max)) < 6 || len(strconv.Itoa(c.Min)) < 6 {
		simpleErrorsChecker(errors.New("Cant have less than 6 numbers"))
	} else if c.Min > c.Max {
		simpleErrorsChecker(errors.New("Min cant be higher than max"))
	}
	if len(strconv.Itoa(int(c.Max))) != len(strconv.Itoa(int(c.Min))) {
		simpleErrorsChecker(errors.New(ERROR_SAME_LENGTH))
	}
}


