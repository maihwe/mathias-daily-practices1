package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}
func masin() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(hexToDecimal("FF"))
	fmt.Println(hexToDecimal("FA"))
}
