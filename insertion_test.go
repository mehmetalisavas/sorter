package sorter

import (
	"fmt"
	"testing"
)

func TestInsertionSortManual(t *testing.T) {
	arrRandom := []int{7, 2, 5, 4, 6, 9}
	resultArr := []int{2, 4, 5, 6, 7, 9}
	arr := InsertionSort(arrRandom)
	exp := resultArr
	equals(t, exp, arr)
}

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
