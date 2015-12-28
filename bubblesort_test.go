package sorter

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	generatedArr, err := GenerateArray(8)
	if err != nil {
		t.Fatal(err.Error())
	}

	arr := BubbleSort(generatedArr)
	isSorted, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !isSorted {
		t.Fatal("Array is not sorted")
	}
}
