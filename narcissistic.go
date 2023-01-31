package main

import "fmt"

func main() {
	var num, tempNum, remaind int

	var res int

	fmt.Scan(&num)

	tempNum = num

	for {
		remaind = tempNum % 10
		res += remaind * remaind * remaind
		tempNum /= 10

		if tempNum == 0 {
			break
		}
	}

	if res == num {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
