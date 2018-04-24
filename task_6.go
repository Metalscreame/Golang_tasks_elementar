package main

import (
	"fmt"
	"math"
	"encoding/json"
)

/*
Вывести через запятую ряд длиной n, состоящий из натуральных чисел, квадрат которых не меньше заданного m.

Входные параметры: длина и значение минимального квадрата
Выход: строка с рядом чисел

 */

type ErrorResponse struct {
	Status string
	Reason string
}

func main() {

	var length, m int
	var errorString string
	fmt.Print("Type length \"n\" : ")
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("\n")
		fmt.Scanln(&errorString)
		main()
	}
	fmt.Print("Type sqrt \"m\" : ")
	_, err2 := fmt.Scanf("%d", &m)
	if err2 != nil {
		fmt.Printf(err2.Error())
		fmt.Printf("\n")
		fmt.Scanln(&errorString)
		main()
	}

	startPointFl := math.Ceil(math.Sqrt(float64(m)))
	fmt.Printf("The sqrt of m is %f", math.Sqrt(float64(m)))
	startPoint := int(startPointFl)
	fmt.Printf("\nNumbers cant be lower than %d", startPoint)

	//for example, if the number is 100 then ceil ceil will return 10 which is sqrt of 100
	if int(math.Sqrt(float64(m))) == startPoint {
		startPoint++
	}

	fmt.Printf("\nThe result is: \n")
	for i := 0; i < int(length)-1; i++ {
		if startPoint == m {
			fmt.Printf("\n\nNumbers cant be printed anymore, because numbers has reached m\n")
			break
		}
		fmt.Print(startPoint)
		fmt.Print(",")
		startPoint++
	}
	fmt.Print(startPoint)
}
func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

