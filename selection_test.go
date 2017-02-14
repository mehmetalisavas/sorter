package sorter

import (
	"fmt"
	"testing"
)

// TestSelectionSort generates unsorted array and sorts this array.
// Then checks if sorted of not
func TestSelectionSort(t *testing.T) {
	generatedArr, err := GenerateArray(8)
	if err != nil {
		t.Fatal(err.Error())
	}

	arr := SelectionSort(generatedArr)
	b, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !b {
		fmt.Println("Array is not sorted")
		return
	}
}

// TestSelectionSortParallel tests the selections sort algorithm
func TestSelectionSortParallel(t *testing.T) {
	tests := []struct {
		arraySize int
		isSorted  bool
		err       error
	}{
		{0, false, ErrArrayNoLength},
		{1, true, nil},
		{2, true, nil},
		{3, true, nil},
		{4, true, nil},
		{10, true, nil},
	}

	for _, test := range tests {
		test := test // capture variable
		t.Run("", func(t *testing.T) {
			t.Parallel()
			arr, err := GenerateArray(test.arraySize)
			if err != test.err {
				t.Fatalf("error should be: %v, but got: %v", test.err, err)
			}

			if len(arr) != test.arraySize {
				t.Fatalf("array length should equal: %v, but got: %v", test.arraySize, len(arr))
			}

			sorted, err := IsSorted(SelectionSort(arr))
			if err != nil {
				t.Fatalf("error should be nil, but got:%v", err)
			}
			if !sorted {
				t.Fatalf("array should be sorted:%v", sorted)
			}
		})
	}
}
