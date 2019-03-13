package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)

	println("max value:=", ret)

	dowhile(10)
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = GetYangHuiTriangleNextLine(nums)
		//println(nums)
		fmt.Println(nums)
	}
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func dowhile(times int) {
	println("for start")
	for a := 0; a < times; a++ {
		println("value=", a)
	}
	println("for end")
}

func GetYangHuiTriangleNextLine(intArr []int) []int {
	var out []int
	var i int
	arrLen := len(intArr)
	out = append(out, 1)
	if 0 == arrLen {
		return out
	}

	for i = 0; i < arrLen-1; i++ {
		out = append(out, intArr[i]+intArr[i+1])
	}
	out = append(out, 1)
	return out
}
