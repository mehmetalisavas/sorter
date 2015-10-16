package sorter

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestInsertion(t *testing.T) {
	arrRandom := []int{7, 2, 5, 4, 6, 9}
	resultArr := []int{2, 4, 5, 6, 7, 9}
	arr := InsertionSort(arrRandom)
	exp := resultArr
	equals(t, exp, arr)
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.Fail()
	}
}
