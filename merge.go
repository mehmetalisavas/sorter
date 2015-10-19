package sorter

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
