package algorithm

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	values := []int{2,3,1}
	sorted, err := Sort(values)
	if nil != err {
		fmt.Println(err)
	}else {
		fmt.Println(sorted)
	}
}
