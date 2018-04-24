package main

import (
	"encoding/json"
	"math"
	"fmt"
	"sort"
	"bufio"
	"os"
	"errors"
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

type ErrorResponse struct {
	Status string
	Reason string
}

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
	Sqrt     float64
}

type BySquare []Triangle

//formula gerona
func getSquare(t Triangle) (float64, error) {
	p := 0.5 * (t.A + t.B + t.C)
	res := math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if res == math.NaN() {
		return res, errors.New("Numbers are wrong for a triangle")
	}
	return res, nil
}

func main() {

	var err error
	fmt.Printf("Enter an triangle object,  (print 'done' to exit adding):")
	fmt.Print("For example {\"vertices\": \"ABC\",\"a\": 10,\"b\": 20,\"c\": 22.36}")

	s := `{"vertices":"ABC","a": 10,"b": 20,"c": 22.36}`

	var t Triangle
	json.Unmarshal([]byte(s), &t)

	trianglesArray := make([]Triangle, 0)

	//parse jsons
	//	var readLn string
	fmt.Printf("\n")
	for true {
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		if line == "done\n" {
			break
		}
		var t Triangle

		err = json.Unmarshal([]byte(line), &t)
		if err != nil {
			//valid json
			errHandler(err)
			break
		}

		t.Sqrt, err = getSquare(t) //calculate every's triangle square
		if err != nil {
			//numbers may be not valid for a triangle, wrong sizes
			errHandler(err)
			break
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

func  errHandler(err error) {
	errResponse := &ErrorResponse{Status:"failed", Reason:err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}