package main


/*
Есть два конверта со сторонами (a,b) и (c,d). Требуется определить, можно ли один конверт вложить в другой. Программа должна обрабатывать ввод чисел с плавающей точкой.

Входные параметры: объекты конверт1 и конверт2
Выход: номер конверта, если вложение возможно, 0 – если вложение невозможно.

 */


 //вопрос. к чему тут числа  с плавающей и что с чем сравнивать?

import (
	"fmt"
	"encoding/json"
)

const CONVERT_1  ="конверт1"
const CONVERT_2  ="конверт2"

func main() {
	var input string

	fmt.Scanln(input)

	if input == CONVERT_1{

	}else if input == CONVERT_2{

	}else {
		fmt.Printf("Wrong input")
		main()
	}


}
func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

type ErrorResponse struct {
	Status string
	Reason string
}