package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
	"os"
)

const(
	ERROR_SIGNED = "Numbers cant be signed, or float or 0"
	ERROR_ONE_SYMBOL = "There can be only one symbol"
	ERROR_SAME_LENGTH = "Must be the same lengths"
	ERROR_WRONG_INPUT = "Wrong input"
	ERROR_WRONG_MAIN_INPUT = "Wrong input. Numbers cant be signed, or float or 0 and must be from 1 to 7"
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
		os.Exit(1)
	}

	//if float check
	task,err = strconv.Atoi(buffer)
	if err != nil || task<1 ||task>7{
		errHandler(errors.New(ERROR_WRONG_MAIN_INPUT))
		os.Exit(1)
	}

	switch task {
	case 1:
		taskOneMain()
	case 2:
		taskTwoMain()
	case 3:
		taskThreeMain()
	case 4:
		taskFourMain()
	case 5:
		taskFiveMain()
	case 6:
		taskSixMain()
	case 7:
		taskSevenMain()

	}

}
