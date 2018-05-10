package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
	"os"
)

const (
	ERROR_SIGNED           = "Numbers cant be signed, or float or 0 or string"
	ERROR_ONE_SYMBOL       = "There can be only one symbol"
	ERROR_SAME_LENGTH      = "Must be the same lengths"
	ERROR_WRONG_INPUT      = "Wrong input"
	ERROR_WRONG_MAIN_INPUT = "Wrong input. Numbers cant be signed, or float or 0 and must be from 1 to 7"
	ERROR_CODE             = 1
)

type ErrorResponse struct {
	Status string
	Reason string
}

func main() {
	for true{
		taskDispatcher(mainInput())
		fmt.Println()
	}
}

func mainInput() (taskNum int) {
	var buffer string
	fmt.Print("Choose a task (1-7): ")
	_, err := fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err)

	taskNum, err = strconv.Atoi(buffer)
	if err != nil || taskNum < 1 || taskNum > 7 {
		simpleErrorsChecker(errors.New(ERROR_WRONG_MAIN_INPUT))
	}
	return
}

func taskDispatcher(num int) {
	switch num {
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

func simpleErrorsChecker(err error) {
	if err != nil {
		errHandler(err)
		os.Exit(ERROR_CODE)
	}
}

func float64InputChecker(num float64, err error) {
	if err != nil || num <= 0 {
		errHandler(errors.New(ERROR_WRONG_INPUT))
		os.Exit(ERROR_CODE)
	}
}

func int32InputChecker(num int, err error) {
	if err != nil || num <= 0 {
		errHandler(errors.New(ERROR_SIGNED))
		os.Exit(ERROR_CODE)
	}
}

func errHandler(err error) {
	errResponse := &ErrorResponse{Status: "failed", Reason: err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

func floatToString(num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(num, 'f', 6, 64)
}

func getValidIntegerFromString(str string)(res int){
	res, err := strconv.Atoi(str)
	int32InputChecker(res, err) // if there is error, wont go further
	return
}

func getValidScannedIntegerValue(messageToAsk string ) (res int) {
	var buffer string
	fmt.Printf("Enter %s and press Enter: ", messageToAsk)
	_, err := fmt.Scanf("%s", &buffer)
	simpleErrorsChecker(err)
	res=getValidIntegerFromString(buffer)
	int32InputChecker(res, err)
	return
}

func getValidScannedString(message string)(bufferString string) {
	fmt.Println(message)
	_, err := fmt.Scanf("%s", &bufferString)
	simpleErrorsChecker(err)
	return
}

func getValidScannedStringValue(messageToAsk string) (value string) {
	fmt.Printf("Enter %s and press Enter: ", messageToAsk)
	_, err := fmt.Scanf("%s", &value)
	simpleErrorsChecker(err)
	return
}

func getValidScannedFloatValue(messageToAsk string) (f float64) {
	fmt.Printf("Enter %s and press Enter: ", messageToAsk)
	_, err := fmt.Scanf("%f", &f)
	float64InputChecker(f, err)
	return f
}

func getSeparatedDigitsFromNumber(currentNumber int) (digits []int) {
	for currentNumber > 0 {
		digit := currentNumber % 10
		digits = append(digits, digit)
		currentNumber /= 10
	}
	return
}