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

{"vertices": "ABC","a": 10,"b": 20,"c": 22.36}
{"vertices": "CBA","a": 16,"b": 15,"c": 6}
{"vertices": "BAC","a": 10,"b": 9,"c": 16}
 */

const NUMBER_OF_TRIANBLES  = 3

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
	Sqrt     float64
}


func taskThreeMain() {

	trianglesArray, err := getTriangles()
	simpleErrorsChecker(err)

	sort.Slice(trianglesArray, func(i, j int) bool {
		return trianglesArray[i].Sqrt < trianglesArray[j].Sqrt
	})
	printResultTriangles(trianglesArray)
}

//gerons formula
func (t * Triangle) calculateSquare() {
	p := 0.5 * (t.A + t.B + t.C)
	t.Sqrt = math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if t.Sqrt == math.NaN() || t.Sqrt == 0 {
		simpleErrorsChecker(errors.New("Numbers are wrong for a triangle"))
	}
}

func getTriangles() ([]Triangle, error) {
	var err error
	var regex = regexp.MustCompile(`^[A-C]{3}$`)
	triangleSlice := make([]Triangle, 0)

	fmt.Printf("Enter three triangle objects\n")
	fmt.Print("For example : \n{\"vertices\": \"ABC\",\"a\": 10,\"b\": 20,\"c\": 22.36}\n")

	for i := 0; i < NUMBER_OF_TRIANBLES; i++ {
		var t Triangle
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		err = json.Unmarshal([]byte(line), &t)
		simpleErrorsChecker(err)

		if regex.MatchString(t.Vertices) != true {
			simpleErrorsChecker(errors.New("Wrong name"))
		}

		if t.A <= 0 || t.B <= 0 || t.C <= 0 {
			simpleErrorsChecker(errors.New(ERROR_SIGNED))
		}

		if !checkNameAndNumbers(t) {
			simpleErrorsChecker(errors.New("Names dont fit actual values"))
		}

		t.calculateSquare() //calculate every's triangle square
		triangleSlice = append(triangleSlice, t)
	}

	if triangleSlice[0].Vertices == triangleSlice[1].Vertices ||
		triangleSlice[0].Vertices == triangleSlice[2].Vertices ||
		triangleSlice[1].Vertices == triangleSlice[2].Vertices {
		return nil, errors.New("Names cant be the same")
	}else{
		return triangleSlice, nil
	}
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
	//must not be reached
	return false
}

func printResultTriangles(trianglesArray []Triangle)  {
	for _, el := range trianglesArray {
		fmt.Printf(el.Vertices + " " + floatToString(el.Sqrt))
		fmt.Print(" ")
	}
}


