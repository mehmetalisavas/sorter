package sorter

// MergeSort divides the unsorted list sublists that containing 1 element
// After this process, gets these each of list that has 1 element
// and merge sublists repeatedly to produce new sorted  sublists
// This process will last until there is only 1 sublist remaining
// This remaining list will be sorted list.
//
// In example ; Assume an array has {4,3,6,2,9,8}
// After MergeSort called
// we have 2 sublists that have for left {4,3,6} and for right {2,9,8}
// List will look like this
// 					{4,3,6,2,9,8}
// 			{4,3,6}---------------|-----------{2,9,8}
// 		{4}--------|------{3,6}		  {2}---------|---------{9,8}
// 		{4}		{3}--|--{6}	  {2}   	    {9}---|---{8}
//---- After this section
// Merge will work to bring together the sublists as sorted
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	l := MergeSort(left)
	r := MergeSort(right)
	merge := Merge(l, r)

	return merge
}

func Merge(left, right []int) []int {
	arr := make([]int, len(left)+len(right))
	var i, j, k int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}

	return arr
}
