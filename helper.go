package sorter

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

const (
	// LENGTH is used to define number size
	// all created number (x) will be between 0 < x < LENGTH
	LENGTH = 100
)

// These variables are the error vars for error handling
var (
	ErrNotSorted            = errors.New("Array is not sorted")
	ErrArrayNoLength        = errors.New("Array has no length")
	ErrArrayElementsNotSame = errors.New("Array elements are not same")
)

// GenerateArray generates an array that has length with the given number
// Also created numbers are between 0 and 100
func GenerateArray(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrArrayNoLength
	}

	var arr []int
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(LENGTH))
	}

	return arr, nil
}

// IsSorted checks all consecutive array numbers
// if a[i] > a[i+1] then error occurs
func IsSorted(arr []int) (bool, error) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false, ErrNotSorted
		}
	}

	return true, nil
}

// IsArrayElementsSame checks the given arrays are the same or not
func IsArrayElementsSame(givenArray, expArray []int) (bool, error) {
	for _, giv := range givenArray {
		if !isInArray(giv, expArray) {
			return false, ErrArrayElementsNotSame
		}
	}

	return true, nil
}

// isInArray checks the given element is in the given array or not
func isInArray(element int, array []int) bool {
	for _, a := range array {
		if element == a {
			return true
		}
	}

	return false
}

// equals gets 3 parameters
// First, gets testing packages's itself
// Second parameter is the expectation of value
// Third parameter is the actual value
// Basicly, this is helper function for testing
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.Fail()
	}
}
