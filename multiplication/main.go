package main

import (
	"fmt"
	
)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		restult := i * num
		fmt.Printf("%2d x %2d = %2d\n", i, num, restult)
	}
}
