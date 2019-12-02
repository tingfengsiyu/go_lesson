package main

//函数闭包
import (
	"fmt"
	"strings"
	"time"
)

func recusive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 10 {
		return //退出
	}
	recusive(n + 1)
}

func Adder() func(d int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix

		}
		return name
	}
}

func main() {
	//recusive(0)
	/*
		f := Adder()
		fmt.Println(f(1))
		fmt.Println(f(100))
		fmt.Println(f(1000))
		//1
		//101
		//1101

	*/
	f1 := makeSuffix(".tmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))
	//test.tmp
	//pic.tmp
}
