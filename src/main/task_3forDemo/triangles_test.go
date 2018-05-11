package task_3forDemo

import "testing"

func TestTrianglesSquareSort(t *testing.T) {

	testCases := []struct {
		name             string
		input            []Triangle
		expVerticesNames []string
		needError        bool
	}{
		{
			name: "no error, 3 triangles",
			input: []Triangle{{"ABC", 10, 20, 22.36, 0},
				{"CBA", 16, 15, 6, 0},
				{"BAC", 10, 9, 16, 0}},
			expVerticesNames: []string{"BAC", "CBA", "ABC"},
			needError:        false,
		},
		{
			name: "error, 3 triangles, one side is negative",
			input: []Triangle{{"ABC", 10.0, -20.0, 22.36, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"BAC", 10.0, 9, 22.36, 16}},
			needError: true,
		},
		{
			name: "error, 5 triangles, 3 of them have same names",
			input: []Triangle{{"ABC", 10.0, 20.0, 22.36, 0},
				{"CBA", 16.0, 15.0, 6, 0},
				{"BAC", 16, 15, 6, 0},
				{"BAC", 16, 15, 6, 0},
				{"BAC", 16, 15, 6, 0}},
			needError: true,
		},
		{
			name: "error, 2 triangles, 2nd have wrong triangle sizes",
			input: []Triangle{{"ABC", 10.0, 20.0, 22.36, 0},
				{"BAC", 10.0, 9, 22.36, 0}},
			needError: true,
		},
		{
			name: "error, empty slice",
			input: []Triangle{},
				needError: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			actualSlice, err := TrianglesSquareSort(testCase.input)
			if testCase.needError && err == nil {
				t.Errorf("for returned error: want <nil> but got \"%v\"", err)
			} else if !testCase.needError && err != nil {
				for i := range actualSlice {
					t.Logf("want %v and got %v", testCase.expVerticesNames[i], actualSlice[i])
				}
			}
		})
	}
}
