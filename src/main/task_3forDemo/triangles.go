package task_3forDemo

import (
	"math"
	"sort"
	"errors"
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

[{
	"vertices": "ABC",
	"a": 10,
	"b": 20,
	"c": 22.36
}, {
	"vertices": "CBA",
	"a": 16,
	"b": 15,
	"c": 6
}, {
	"vertices": "BAC",
	"a": 10,
	"b": 9,
	"c": 16
}]

 */
const ERROR_SIGNED = "Numbers cant be signed, or float or 0 or string"

type Triangle struct {
	Vertices string
	A        float64
	B        float64
	C        float64
	Sqrt     float64
}

func TrianglesSquareSort(trianglesToSortSlice []Triangle) ( []Triangle, error) {
	err := validateTriangles(trianglesToSortSlice)
	if err != nil {
		return trianglesToSortSlice, err
	}

	sort.Slice(trianglesToSortSlice, func(i, j int) bool {
		return trianglesToSortSlice[i].Sqrt < trianglesToSortSlice[j].Sqrt
	})

	return trianglesToSortSlice,nil
}

func validateTriangles(t []Triangle) (err error) {
	if len(t) == 0 || len(t) == 1{
		return errors.New("Size cant be nil")
	}

	for i := 0; i < len(t); i++ {
		err = t[i].validateSingleTriangle()
		if err != nil {
			return err
		}
		err = t[i].calculateSquare()
		if err != nil {
			return err
		}
	}


	return
}

func (t *Triangle) validateSingleTriangle() (err error) {
	var regex = regexp.MustCompile(`^[A-C]{3}$`) //своя реализация в анмаршалинге
	if regex.MatchString(t.Vertices) != true {
		return errors.New("Wrong name")
	}
	if t.A <= 0 || t.B <= 0 || t.C <= 0 {
		return errors.New(ERROR_SIGNED)
	}
	if !t.checkNameAndNumbers() {
		return errors.New("Names dont fit actual values")
	}
	return
}

func (t *Triangle) checkNameAndNumbers() bool {
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
	return false
}

func (t *Triangle) calculateSquare() (err error) {
	p := 0.5 * (t.A + t.B + t.C)
	t.Sqrt = math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
	if math.IsNaN(t.Sqrt) || t.Sqrt == 0 {
		return errors.New("Numbers are wrong for a triangle. Can not calculate square")
	}
	return
}
