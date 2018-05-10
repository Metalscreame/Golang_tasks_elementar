package main

/*
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить,
можно ли один конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей точкой.

Входные параметры: объекты конверт1 и конверт2
Выход: номер конверта, если вложение возможно, 0 – если вложение невозможно.

 */

import (
	"fmt"
	"math"
)

const (
	MESSAGE_A_SIZE = "the size of a"
	MESSAGE_B_SIZE = "the size of b"
	MESSAGE_C_SIZE = "the size of c"
	MESSAGE_D_SIZE = "the size of d"
	MESSAGE_AB_CD = "AB -> CD"
	MESSAGE_CD_AB = "CD -> AB"
	MESSAGE_DOESNT_FIT = "0"
	MESSAGE_TASK_2_INPUT = "You must enter 4 sizes of 2 converts. First convert a and b, second - c and d"
)



func taskTwoMain() {
	getConvertsResult(inputConverts())
}

func convertsFitCheck(a, b, p, q float64) (bool) {
	if  p <= a && q <= b ||
		p > a && b >= (2*p*q*a+((p*p-q*q)*math.Sqrt(p*p+q*q-a*a)))/(p*p+q*q) {
		return true
	} else {
		return false
	}
}

func inputConverts() (a, b, c, d float64) {
	fmt.Println(MESSAGE_TASK_2_INPUT)
	a= getValidScannedFloatValue(MESSAGE_A_SIZE)
	b= getValidScannedFloatValue(MESSAGE_B_SIZE)
	c= getValidScannedFloatValue(MESSAGE_C_SIZE)
	d= getValidScannedFloatValue(MESSAGE_D_SIZE)
	return
}

func getConvertsResult(a, b, c, d float64) {
	result := convertsFitCheck(a, b, c, d)
	if !result {
		result = convertsFitCheck(c, d, a, b)
		if !result {
			fmt.Println(MESSAGE_DOESNT_FIT)
		} else {
			fmt.Println(MESSAGE_AB_CD)
		}
	} else {
		fmt.Println(MESSAGE_CD_AB)
	}
}
