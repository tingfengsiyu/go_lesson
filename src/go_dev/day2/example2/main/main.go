package main

import (
	"fmt"
	a"go_dev/day2/example2/add"		// 别名
	//-"go_dev/day2/example2/add"
)

func  main()  {
	//fmt.Println("Name=",add.Name)
	fmt.Println("Name=",a.Name)	//引用别名
	fmt.Print("Age=",a.Age)
}
