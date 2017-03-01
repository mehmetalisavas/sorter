package sorter

import (
	"fmt"
	"testing"
)

// TestInsertionSortManual sorts the random array and should find result array
func TestInsertionSortManual(t *testing.T) {
	arrRandom := []int{7, 2, 5, 4, 6, 9}
	resultArr := []int{2, 4, 5, 6, 7, 9}
	arr := InsertionSort(arrRandom)
	exp := resultArr

	equals(t, exp, arr)
}

// TestInsertionSort sorts the generated array with insertion sort algorithm
func TestInsertionSort(t *testing.T) {
	generatedArr, err := GenerateArray(8)
	if err != nil {
		t.Fatal(err.Error())
	}

	arr := InsertionSort(generatedArr)

	b, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !b {
		fmt.Println("Array is not sorted")
		return
	}
}

// TestRecursiveInsertionSort sorts the generated array with insertion sort algorithm recursively
func TestRecursiveInsertionSort(t *testing.T) {
	generatedArr, err := GenerateArray(8)
	if err != nil {
		t.Fatal(err.Error())
	}

	arr := RecursiveInsertionSort(generatedArr)

	b, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !b {
		fmt.Println("Array is not sorted")
		return
	}
}

// TestInsertionSortParallel tests the insertion sort with paralleling
func TestInsertionSortParallel(t *testing.T) {
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

			sorted, err := IsSorted(InsertionSort(arr))
			if err != nil {
				t.Fatalf("error should be nil, but got:%v", err)
			}
			if !sorted {
				t.Fatalf("array should be sorted:%v", sorted)
			}
		})
	}
}
