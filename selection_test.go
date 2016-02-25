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
