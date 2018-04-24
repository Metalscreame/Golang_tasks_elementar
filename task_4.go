package main

import (
	"fmt"
	"strconv"
	"encoding/json"
)
type ErrorResponse struct {
	Status string
	Reason string
}

func isPalindrome(input string) (bool, []string) {
	var flag bool
	var res []string
	if len(input)%2 == 0 {
		for i := 0; i < len(input)/2; i++ {
			if input[i] != input[len(input)-i-1] {
				flag = false
				continue
			} else {
				flag = true
				res = append(res, string(input[i]))
			}
		}
	} else {
		runes := []rune(input)
		var res2 []string
		substringed := string(runes[1:len(input)])
		for i := 0; i < len(substringed)/2; i++ {
			if substringed[i] != substringed[len(substringed)-i-1] {
				flag = false
				continue
			} else {
				flag = true
				res = append(res, string(substringed[i]))
			}
		}
		//if flag is still false, we'll cut the last char and will see if there is a lonidrome or not
		substringed2 := string(runes[0:len(input)-1])
		for i := 0; i < len(substringed2)/2; i++ {
			if substringed2[i] != substringed2[len(substringed2)-i-1] {
				flag = false
				continue
			} else {
				flag = true
				res2 = append(res2, string(substringed2[i]))
			}
		}

		if len(res) < len(res2) {
			res = res2
		} // else do nothing
	}

	i := len(res)
	for i > 0 {
		res = append(res, string(res[i-1]))
		i--
	}

	return flag, res
}

func main() {
	var input int
	fmt.Println("Enter a number to if its a polindrome or not: ")

	var err error
	_, err = fmt.Scanf("%d", &input)
	if err != nil {
		errHandler(err)
		return
	}

	srtNumber := strconv.Itoa(input)
	flag, res := isPalindrome(srtNumber)
	if flag == false {
		fmt.Println("0")
	} else {
		for _, el := range res {
			fmt.Print(el)
		}
	}
}
func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}

