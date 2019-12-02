package main

import (
	"errors"
	"fmt"
	//"time"
)

func s() {
	var a []int
	a = append(a, 10, 100)
	fmt.Println(a)
	a = append(a, a...) //... 展开数组元素
	fmt.Println(a)      //[10 100 10 100]

}
func initConfig() (err error) {
	return errors.New("Init config failed!!!")
}
func test() {
	//匿名函数
	/*
		defer func(){

			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

	*/
	err := initConfig()
	if err != nil {
		panic(err)
	}
	/*
		b := 0
		a := 100/b
		fmt.Println(a)
		return

	*/
}

func make1() {
	s1 := new([]int)
	fmt.Println(s1)
	//&[]

	s2 := make([]int, 100)
	fmt.Println(s2) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	*s1 = make([]int, 5)
	(*s1)[0] = 100
	s2[0] = 100
	fmt.Println(s1) //&[100 0 0 0 0]
	return
}

func main() {
	/*s()
	var i int
	fmt.Println(i)
	j  := new(int)
	fmt.Println(j)
	for {
		//无限循环
		test()
		time.Sleep(time.Second)
	}

	*/
	make1()
	/*
		&[]
		[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
		 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	*/

}
