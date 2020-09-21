package base

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{3, 5, 2, 1}

	arr = InsertSort(arr)

	fmt.Println(arr)
}
