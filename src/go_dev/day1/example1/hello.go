package main
//每一个go文件就是一个包，引入main

import "fmt"



func add(a int,b int) int{
	var sum int
	sum = a+b
	return sum
}

func main() {
	//入口函数，从main开始执行
	fmt.Println("Hello, World!")
	var c int
	c= add(100,200)
	fmt.Print("add(100,200)------",c)
}