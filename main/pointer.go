package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a变量的地址是：%x\n", &a)
	fmt.Printf("ip 变量存储的指针地址：%x\n", ip)
	fmt.Printf("ip 变量的值：%d\n", *ip)

	var ptr *int

	switch {
	case ptr != nil:
		fmt.Println("ptr不是空指针")
		fallthrough
	case ptr == nil:
		fmt.Println("ptr是空指针")
		fallthrough
	case ip != nil:
		fmt.Println("ip不是空指针")
	default:
		fmt.Println("ip是空指针")
	}
}
