package sorter

// InsertionSort sorts the given array according to insertion sort algoritm
func InsertionSort(arr []int) []int {
	var key int
	var j int
	var resultArr []int

	for i := 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			arr[j+1] = key
		}
	}

	resultArr = arr
	return resultArr
}
