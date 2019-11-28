package main

import (
	"fmt"
	"time"
)

const (
	Female = 1
	Man    = 2
)

func main() {
	second := time.Now().Unix()
	//print(second)
	if second%Female == 0 {
		fmt.Println(Female)
	} else {
		fmt.Println(Man)
	}

}
