package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
)

const(
	ERROR_SIGNED = "Numbers cant be signed, or float or 0"
	ERROR_ONE_SYMBOL = "There can be only one symbol"
	ERROR_SAME_LENGTH = "Must be the same lengths"
	ERROR_WRONG_INPUT = "Wrong input"
	ERROR_WRONG_MAIN_INPUT = "Numbers cant be signed, or float or 0 and must be from 1 to 7"
)


type ErrorResponse struct {
	Status string
	Reason string
}


func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

func main()  {
	var buffer string
	var task int
	fmt.Print("Choose a task (1-7): ")
	_, err := fmt.Scanf("%s", &buffer)
	if err != nil {
		errHandler(err)
		return
	}

	//if float check
	task,err = strconv.Atoi(buffer)
	if err != nil || task<1 ||task>7{
		errHandler(errors.New(ERROR_WRONG_MAIN_INPUT))
		return
	}

	switch task {
	case 1:
		main1()
	case 2:
		main2()
	case 3:
		main3()
	case 4:
		main4()
	case 5:
		main5()
	case 6:
		main6()
	case 7:
		main7()

	}

}
