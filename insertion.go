package sorter

// InsertionSort sorts the given array according to insertion sort algoritm
func InsertionSort(arr []int) []int {
	var key int
	var j int

	for i := 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}

/////// Insertion Sort Recursively ////////

func RecursiveInsertionSort(arr []int) []int {
	lenArr := len(arr)
	if lenArr < 2 {
		return arr
	}
	newArr := arr[:lenArr-1]
	r := RecursiveInsertionSort(newArr)
	lastItem := arr[lenArr-1]

	ins := Insert(r, lastItem)

	return ins
}

func Insert(arr []int, key int) []int {
	newArr := make([]int, len(arr)+1)
	i := len(arr) - 1
	for i >= 0 && arr[i] > key {
		newArr[i+1] = arr[i]
		i--
	}
	newArr[i+1] = key

	for i >= 0 {
		newArr[i] = arr[i]
		i--
	}

	return newArr
}
