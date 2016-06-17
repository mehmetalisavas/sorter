package sorter

// SelectionSort
func SelectionSort(arr []int) []int {
	var min int
	var temp int

	for i := 0; i < len(arr)-1; i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp

	}
	return arr
}
