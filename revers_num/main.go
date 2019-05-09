package main

import (
	"fmt"
	"strconv"
)

func main() {
	xs := []int{
		123,
		1<<31 - 2,
		1<<31 - 1,
		1 << 31,
		6463847412,
		7463847412,
		8463847412,
		-123,
	}
	for _, x := range xs {
		fmt.Printf("%v revers: %v\n", x, reverse(x))
	}

}
func reverse(x int) int {
	var (
		ret      int
		negative bool
		t        = ""
		a        = 1<<31 - 1
	)
	//fmt.Printf("%#v\n", a)
	if x < 0 {
		negative = true
		x = x * -1
	}
	s := strconv.Itoa(x)
	for i := len(s); i > 0; i-- {
		t += string(s[i-1])
	}
	ret, _ = strconv.Atoi(t)
	if ret > a {
		return 0
	}
	if negative {
		ret = ret * -1
	}
	return ret
}
