package main

//水仙花
import (
	"fmt"
)

/*
func isNumber(n int) bool {
	var i, j, k int
	i = n % 10
	j = (n/10) %10
	k = (n/100)%10
	sum :=  i*i*i + j*j*j + k*k*k
	return sum == n
}
func main(){
	var n ,m int
	fmt.Scanf("%d%d",&n,&m)
	for  i :=n; i <m ;i++ {
	if isNumber(i)	== true {
		fmt.Println(i ," is 水仙花")
	} else {
		fmt.Println(n,"is not 水仙花")
	}
	}
}*/
//> main.exe
//100 999
//153  is 水仙花
//370  is 水仙花
//371  is 水仙花
//407  is 水仙花

//阶乘之和 4！+3！+2！+1！
func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := sum(n)
	fmt.Println(s)
}

//>main.exe
//4	输入
//33
