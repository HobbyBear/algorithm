package base

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 5, 2, 1}

	arr = BubbleSort(arr)

	fmt.Println(arr)
}
