package main


/*
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить, можно ли один конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей точкой.

Входные параметры: объекты конверт1 и конверт2
Выход: номер конверта, если вложение возможно, 0 – если вложение невозможно.

 */

import (
	"fmt"
	"errors"
	"math"
)

func main2() {
	var a,b,c,d float64

	fmt.Print("Enter a: ")
	_,err:=fmt.Scanf("%f",&a)
	if err!= nil || a<=0{
		errHandler(errors.New(ERROR_WRONG_INPUT))
		return
	}

	fmt.Println()
	fmt.Print("Enter b: ")
	_,err=fmt.Scanf("%f",&b)
	if err!= nil || a<=0{
		errHandler(errors.New(ERROR_WRONG_INPUT))
		return
	}

	fmt.Println()
	fmt.Print("Enter c: ")
	_,err=fmt.Scanf("%f",&c)
	if err!= nil || a<=0{
		errHandler(errors.New(ERROR_WRONG_INPUT))
		return
	}

	fmt.Println()
	fmt.Print("Enter d: ")
	_,err=fmt.Scanf("%f",&d)
	if err!= nil || a<=0{
		errHandler(errors.New(ERROR_WRONG_INPUT))
		return
	}

	if math.Sqrt(a)+math.Sqrt(b) <= math.Sqrt(c)+math.Sqrt(d)  && a<c && b<d{
		fmt.Println("AB -> CD")
	} else if math.Sqrt(a)+math.Sqrt(b) >= math.Sqrt(c)+math.Sqrt(d) && a>c && b>d{
		fmt.Println("AB <- CD")
	}else {
		fmt.Println("0")
	}

}
