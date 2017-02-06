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
		// never ignore err
		// here we only check isSorted value
		test := test // capture range variable
		t.Run("", func(t *testing.T) {
			t.Parallel()
			isSorted, _ := IsSorted(test.arr)
			if isSorted != test.isSorted {
				t.Fatalf("result of isSorted %v should be %v\n", test.arr, test.isSorted)
			}
		})
	}
}
