package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)
func main()  {
	var pipe chan int
	pipe = make(chan int,1)
	go goroute.Add(100,200,pipe)
	sum :=  <- pipe
	fmt.Println("sum=",sum)
	//for i :=0;i<100;i++{
	//	go test_goroute(i)
	//}
	//time.Sleep(time.Second)
}
