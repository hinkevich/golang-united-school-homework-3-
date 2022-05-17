package homework

import "testing"

func TestTaskOne(t *testing.T) {

	array := []float32{1, 2, 3, 4, 5, 6}
	if average(array) != 3.5 {
		t.Error("is not pass")
	}
}

func TestTaskTwo(t *testing.T) {
	inputArray := []int64{1, 2, 3, 4, 5, 6}
	want := []int64{6, 5, 4, 3, 2, 1}
	checkArray := reverse(inputArray)

	/*if len(want) != len(checkArray) {
		t.Error("different lenght")
	}*/

	for i, v := range want {
		if v != checkArray[i] {
			t.Errorf("want %d but %d", v, checkArray[len(checkArray)-(i+1)])
		}

	}

}

/*
Task 3: Maps
function that returns map values sorted in order of increasing keys.
Input -> {2: "a", 0: "b", 1: "c"}
Output -> ["b", "c", "a"]
Input -> {10: "aa", 0: "bb", 500: "cc"}
Output -> ["bb", "aa", "cc"]
*/
//testData{input: {10: "aa", 0: "bb", 500: "cc"}, want: {"bb", "aa", "cc"}},
type testData struct {
	input map[int]string
	want  []string
}

func TestTaskThree(t *testing.T) {
	tests := []testData{
		{input: map[int]string{2: "a", 0: "b", 1: "c"}, want: []string{"b", "c", "a"}},
		{input: map[int]string{10: "aa", 0: "bb", 500: "cc"}, want: []string{"bb", "aa", "cc"}},
	}
	for _, test := range tests {
		var got = sortMapValues(test.input)

		for i, testValue := range test.want {
			if testValue != got[i] {
				t.Errorf("want: %s // but: %s // position i:%d", testValue, got[i], i)
			}
		}
	}
}
