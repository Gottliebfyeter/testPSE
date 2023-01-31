package main

import "fmt"

func main() {
	var size, num int
	fmt.Printf("Enter size array: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)

	for i := 0; i < size; i++ {
		//fmt.Print(i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Printf("Enter number: ")
	fmt.Scanln(&num)
	res := blueOcean(arr, num)
	fmt.Println(res)
}

func blueOcean(arr []int, num int) []int {
	index := 0

	for _, i := range arr {
		if i != num {
			arr[index] = i
			index++
		}
	}
	return arr[:index]
}
