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
	b, err := IsSorted(arr)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !b {
		fmt.Println("Array is not sorted"
		return
	}
}
