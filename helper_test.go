package sorter

import "testing"

// TestIsSorted tests the arrays if its sorted or not
func TestIsSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		isSorted bool
	}{
		{[]int{0}, true},
		{[]int{0, 0}, true},
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 3, 4}, true},
		{[]int{0, 1, 2, 3, 4}, true},
		{[]int{0, 1, 1}, true},
		{[]int{0, 1, 2, 1}, false},
		{[]int{0, -1}, false},
		{[]int{0, 1, 2, 3, 4, 3}, false},
	}
	for _, test := range tests {
		// if we dont capture variable here, then its not guaranteed all the array
		// values will be tested
		test := test // capture range variable
		t.Run("", func(t *testing.T) {
			t.Parallel()
			// note: never ignore errors
			// here we only check isSorted value
			isSorted, _ := IsSorted(test.arr)
			if isSorted != test.isSorted {
				t.Fatalf("result of isSorted %v should be %v\n", test.arr, test.isSorted)
			}
		})
	}
}

func TestGenerateArray(t *testing.T) {
	tests := []struct {
		item int
		err  error
	}{
		{-1, ErrArrayNoLength},
		{0, ErrArrayNoLength},
		{1, nil},
		{2, nil},
		{3, nil},
		{10, nil},
		{110, nil},
	}
	for _, test := range tests {
		// if we dont capture variable here, then its not guaranteed all the array
		// values will be tested
		test := test // capture range variable
		t.Run("", func(t *testing.T) {
			t.Parallel()
			// never ignore errors
			// here we only check isSorted value
			arr, err := GenerateArray(test.item)
			if err != test.err {
				t.Fatalf("error should be %v, but got: %v", err, test.err)
			}
			if len(arr) > 0 && len(arr) != test.item {
				t.Fatalf("length of array should be %d, but got: %d", len(arr), test.item)
			}
		})
	}
}
