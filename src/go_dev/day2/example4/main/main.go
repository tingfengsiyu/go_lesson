package main

import "fmt"

//交换内存地址实例
//func swap( a int,b int)  {
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println(a, b)
	return
}

func test() {
	var a = 100
	fmt.Println(a)
}
func main() {
	first := 100
	second := 200
	//swap(first,second)
	swap(&first, &second) //取内存地址
	fmt.Println(&first, &second)
	test()
}
