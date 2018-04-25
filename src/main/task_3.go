package main

import (
	"math"
	"fmt"
	"sort"
	"bufio"
	"os"
	"errors"
	"encoding/json"
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
	if res == math.NaN() {
		return res, errors.New("Numbers are wrong for a triangle")
	}
	return res, nil
}

func main3() {

	var err error
	fmt.Printf("Enter three triangle objects\n")
	fmt.Print("For example {\"vertices\": \"ABC\",\"a\": 10,\"b\": 20,\"c\": 22.36}\n")

	trianglesArray := make([]Triangle, 0)

	for i:=0;i<3;i++ {
		var t Triangle
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		err = json.Unmarshal([]byte(line), &t)
		if err != nil {
			errHandler(err)
			return
		}

		if t.A<=0 || t.B<=0 || t.C<=0{
			errHandler(errors.New(ERROR_SIGNED))
			return
		}
		if len(t.Vertices)>3 {
			errHandler(errors.New("Vertices can only exist of 3 symbols"))
			return
		}

		t.Sqrt, err = getSquare(t) //calculate every's triangle square
		if err != nil {
			//numbers may be not valid for a triangle, wrong sizes
			errHandler(err)
			return
		}
		trianglesArray = append(trianglesArray, t)
	}

	//sorting
	sort.Slice(trianglesArray, func(i, j int) bool {
		return trianglesArray[i].Sqrt < trianglesArray[j].Sqrt
	})

	for _, el := range trianglesArray {
		fmt.Print(el.Vertices)
		fmt.Print(" ")
	}
}
