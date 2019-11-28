package main

import (
	"fmt"
	"time"
)

func main()  {
	for i :=0;i<100;i++{
		go test_goroute(i)
	}
	time.Sleep(time.Second)
	fmt.Print(11)
}
