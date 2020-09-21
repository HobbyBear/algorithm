package base

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	index := BinarySearch(arr, 3)
	fmt.Println(index)
}

func TestBinarySearchLoop(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	index := BinarySearchLoop(arr, 3)
	fmt.Println(index)
}

func BenchmarkBinarySearch(b *testing.B) {
	cap := 100000000
	arr := make([]int,0,cap)
	for i := 0; i < cap; i++ {
		arr = append(arr, i)
	}
	b.ResetTimer()
	for i :=0; i < b.N; i++{
		BinarySearch(arr,100)
	}
}

func BenchmarkBinarySearchLoop(b *testing.B) {
	cap := 100000000
	arr := make([]int,0,cap)
	for i := 0; i < cap; i++ {
		arr = append(arr, i)
	}
	b.ResetTimer()
	for i :=0; i < b.N; i++{
		BinarySearchLoop(arr,100)
	}
}
