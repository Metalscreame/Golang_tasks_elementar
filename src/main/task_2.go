package main

/*
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить, можно ли один конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей точкой.

Входные параметры: объекты конверт1 и конверт2
Выход: номер конверта, если вложение возможно, 0 – если вложение невозможно.

 */

import (
	"fmt"
	"math"
)

func taskTwoMain() {
	getConvertsResult(inputConverts())
}

func convertsFitCheck(a, b, c, d float64) (bool) {
	if c <= a && d <= b ||
		c >= a && b >= (2*c*d*a+(c*c-d*d)*math.Sqrt(c*c+d*d-a*a))/(c*c+d*d) {
		return true
	} else {
		return false
	}

}
func inputConverts() (a, b, c, d float64) {

	fmt.Println("You must enter 4 sizes of 2 converts. First convert a and b, second - c and d")
	fmt.Print("Enter a: ")
	_, err := fmt.Scanf("%f", &a)
	float64InputChecker(a, err)

	fmt.Print("Enter b: ")
	_, err = fmt.Scanf("%f", &b)
	float64InputChecker(b, err)

	fmt.Print("Enter c: ")
	_, err = fmt.Scanf("%f", &c)
	float64InputChecker(c, err)

	fmt.Print("Enter d: ")
	_, err = fmt.Scanf("%f", &d)
	float64InputChecker(d, err)
	return
}

func getConvertsResult(a, b, c, d float64) {
	result := convertsFitCheck(a, b, c, d)
	if !result {
		result = convertsFitCheck(c, d, a, b)
		if !result {
			fmt.Println("0")
		} else {
			fmt.Println("AB -> CD")
		}
	} else {
		fmt.Println("AB <- CD")
	}
}
