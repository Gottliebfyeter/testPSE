package main

import "fmt"

func main() {
	var size int
	var find string
	fmt.Printf("Enter size array: ")
	fmt.Scanln(&size)
	var arr = make([]string, size)

	for i := 0; i < size; i++ {
		fmt.Scanf("%s", &arr[i])
	}
	fmt.Printf("Enter find word: ")
	fmt.Scanln(&find)
	res := findNeedle(arr, find)
	fmt.Println(res)
}

func findNeedle(arr []string, subArr string) int {

	for i, n := range arr {
		if subArr == n {
			return i
		}
	}
	return len(arr)
}
