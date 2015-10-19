package sorter

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	generatedArr, err := GenerateArray(8)
	if err != nil {
		t.Fatal(err.Error())
	}

	arr := MergeSort(generatedArr)
	b, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !b {
		fmt.Println("Array is not sorted")
		return
	}
}
