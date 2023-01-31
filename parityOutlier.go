package main

import "fmt"

func main() {
	var size int
	var even, odd []int
	fmt.Printf("Enter size array: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)

	for i := 0; i < size; i++ {
		//fmt.Print(i)
		fmt.Scanf("%d", &arr[i])
	}

	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	if len(even) == 1 {
		fmt.Println(even)
	}
	if len(odd) == 1 {
		fmt.Println(odd)
	}
	if len(even) != 1 && len(odd) != 1 {
		fmt.Println("false")
	}
}
