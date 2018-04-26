package main

import (
	"math"
	"fmt"
	"sort"
	"bufio"
	"os"
	"errors"
	"encoding/json"
	"regexp"
	"strconv"
)

/*
Вывести треугольники в порядке убывания их площади.

Входные параметры: массив объектов треугольник
Выход: упорядоченный массив имён треугольников
Примечание:
•	Расчёт площади треугольника должен производится по формуле Герона.
•	Каждый треугольник определяется именами вершин и длинами его сторон.
•	Приложение должно обрабатывать ввод чисел с плавающей точкой:

{
vertices: ‘ABC’,
a: 10,
b: 20,
c: 22.36
}

 */

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
	Sqrt     float64
}

//formula gerona
func getSquare(t Triangle) (float64, error) {
	p := 0.5 * (t.A + t.B + t.C)
	res := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if res == math.NaN() || res == 0 {
		return res, errors.New("Numbers are wrong for a triangle")
	}
	return res, nil
}

func taskThreeMain() {
	var err error

	fmt.Printf("Enter three triangle objects\n")
	fmt.Print("For example : \n {\"vertices\": \"ABC\",\"a\": 10,\"b\": 20,\"c\": 22.36}\n")

	trianglesArray, err := getTriangles()
	simpleErrorChecker(err, "")

	if trianglesArray[0].Vertices == trianglesArray[1].Vertices ||
		trianglesArray[0].Vertices == trianglesArray[2].Vertices ||
		trianglesArray[1].Vertices == trianglesArray[2].Vertices {
		errHandler(errors.New("Names cant be the same"))
		os.Exit(1)
	}

	//sorting
	sort.Slice(trianglesArray, func(i, j int) bool {
		return trianglesArray[i].Sqrt < trianglesArray[j].Sqrt
	})

	for _, el := range trianglesArray {
		fmt.Print(el.Vertices + " " + strconv.FormatFloat(el.Sqrt, 'E', -1, 64))
		fmt.Print(" ")
	}
}

func getTriangles() ([]Triangle, error) {
	var err error
	var regex = regexp.MustCompile(`^[A-C]{3}$`)
	triangleSlice := make([]Triangle, 0)
	for i := 0; i < 3; i++ {
		var t Triangle
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		err = json.Unmarshal([]byte(line), &t)
		if err != nil {
			return nil, err
		}

		if regex.MatchString(t.Vertices) != true {
			return nil, errors.New("Wrong name")
		}

		if t.A <= 0 || t.B <= 0 || t.C <= 0 {
			return nil, errors.New(ERROR_SIGNED)
		}

		if !checkNameAndNumbers(t) {
			return nil, errors.New("Names dont fit actual values")
		}

		t.Sqrt, err = getSquare(t) //calculate every's triangle square
		if err != nil {
			//numbers may be not valid for a triangle, wrong sizes
			return nil, err
		}
		triangleSlice = append(triangleSlice, t)
	}

	return triangleSlice, nil
}

func checkNameAndNumbers(t Triangle) bool {
	r := []rune(t.Vertices)

	if string(r[2]) == "A" && t.A > t.B && t.A > t.C {
		if string(r[0]) == "B" && t.B < t.C {
			return true
		} else if string(r[0]) == "C" && t.C < t.B {
			return true
		} else {
			return false
		}
	}

	if string(r[2]) == "B" && t.B > t.A && t.A > t.C {
		if string(r[0]) == "C" && t.C < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.C {
			return true
		} else {
			return false
		}
	}

	if string(r[2]) == "C" && t.C > t.B && t.C > t.A {
		if string(r[0]) == "B" && t.B < t.A {
			return true
		} else if string(r[0]) == "A" && t.A < t.B {
			return true
		} else {
			return false
		}
	}
	//wont be reached
	return false

}
