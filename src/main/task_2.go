package main

/*
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить, можно ли один конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей точкой.

Входные параметры: объекты конверт1 и конверт2
Выход: номер конверта, если вложение возможно, 0 – если вложение невозможно.

 */

import (
	"fmt"
)

func taskTwoMain() {
	a, b, c, d := inputConverts()
	printConvertsResult(a, b, c, d )
}

func convertsCheck(A, B, C, D float64) bool {

	if A < B {
		temp := A
		A = B
		B = temp
	}
	if C < D {
		temp := C
		C = D
		D = temp
	}

	if A >= C && B >= D {
		return true
	}
	if C*C+D*D > A*A+B*B {
		return false
	}
	if A*A >= C*C+D*D {
		return B >= C
	}

	return (A*B > 2*C*D) && ((C*C+D*D-B*B)*(C*C+D*D-A*A) <= A*A*B*B-4*A*B*C*D+4*C*C*D*D)
}

func inputConverts() (a, b, c, d float64) {

	fmt.Print("Enter a: ")
	_, err := fmt.Scanf("%f", &a)
	float64InputChecker(a,err)

	fmt.Print("Enter b: ")
	_, err = fmt.Scanf("%f", &b)
	float64InputChecker(b,err)

	fmt.Print("Enter c: ")
	_, err = fmt.Scanf("%f", &c)
	float64InputChecker(c,err)

	fmt.Print("Enter d: ")
	_, err = fmt.Scanf("%f", &d)
	float64InputChecker(d,err)

	return a, b, c, d
}

func printConvertsResult(a,b,c,d float64){
	result := convertsCheck(a, b, c, d)
	if !result {
		result = convertsCheck(c, d, a, b)
		if !result {
			fmt.Println("0")
		} else {
			fmt.Println("AB -> CD")
		}
	} else {
		fmt.Println("AB <- CD")
	}
}