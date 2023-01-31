package main

import "fmt"

func main() {
	var size int
	fmt.Printf("Enter size array: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	res := parityOutlier(arr)
	fmt.Println(res)
}

func parityOutlier(arr []int) string {
	var even, odd []int
	var res string
	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	if len(even) == 1 {
		res = fmt.Sprintf("%d (the only even number)", even)
	}
	if len(odd) == 1 {
		res = fmt.Sprintf("%d (the only odd number)", odd)
	}
	if len(even) == 0 {
		res = fmt.Sprintf("false (all odd integer, no outlier)")
	}
	if len(odd) == 0 {
		res = fmt.Sprintf("false (all even integer, no outlier)")
	}
	return res
}
