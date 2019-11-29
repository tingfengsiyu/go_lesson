package main

import "fmt"

func main() {
	var a int = 1
	var str = "hello \n"
	var str1 = `
	this is a,
	hhh------1.
	`
	var d bool
	var b byte = 'c'
	fmt.Println("test", str, str1)
	fmt.Printf("%c\n", b) //字节
	fmt.Printf("%d\n", a) //数字
	fmt.Printf("%v", d)   //原样打印
}
