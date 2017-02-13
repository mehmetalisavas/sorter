package sorter

import "testing"

// TestBubbleSort generates the array with 8 elements randomly
// and checks the BubbleSort is running correctly or not
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

func TestBubbleSortParallel(t *testing.T) {
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

			sorted, err := IsSorted(BubbleSort(arr))
			if err != nil {
				t.Fatalf("error should be nil, but got:%v", err)
			}
			if !sorted {
				t.Fatalf("array should be sorted:%v", sorted)
			}
		})
	}

}
