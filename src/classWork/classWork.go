package classWork

import (
"fmt"
"sort"
"math/rand"
)

func classWork1() {
	//fmt.Print(findAvarage(1,2,3,4))

	slice := make([]int, 13)

	fmt.Println("The slice is:")
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(12)
		fmt.Print(" ", slice[i])
	}

	fmt.Println()
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})

	fmt.Println()
	for _, el := range slice {
		fmt.Print(" ",el)
	}
	fmt.Println()

	//var temp int

	tempSlice := make([]int, len(slice))

	j := len(tempSlice) - 1
	k := 0
	for i := 0; i < len(slice); i++ {
		if i%2 != 0 {
			j--
			tempSlice[j+1] = slice[i]
		} else {
			tempSlice[k] = slice[i]
			k++
		}
	}

	for _, el := range tempSlice {
		fmt.Print(" ",el)
	}
}

