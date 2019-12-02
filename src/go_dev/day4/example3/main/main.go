package main

import (
	"fmt"
	"sort"
)

/*
import  "fmt"

//数组
func test1(){
	var a [10]int
	a[0] = 100
	a[2] = 200

	for i :=0;i < len(a);i++{
	fmt.Println(a[i])
	}
	for index,val :=range a{
	fmt.Printf("a[%d]=%d\n",index,val)
}
}
func test2(){
	var a [10]int
	b :=a
	b[0] = 100
	fmt.Println(a)
}

func test3(err [5]int){
	err[0] =1000
}
var age2 = [...] int {1,2}
var a3 = [...] int {1:100,2:300}

func main(){
	test2() //值未变 [0 0 0 0 0 0 0 0 0 0]
	var a [5]int
	test3(a)
	fmt.Println(a) //[0 0 0 0 0]
	fmt.Println(age2,a3)
}


//切片

func testSlice(){
	var slice [] int
	var arr [5]int =[...]int{1,2,3,4,5}


	slice = arr[2:3] //包含2 不包含3
	slice = arr[:]
	fmt.Println(slice[:]) //包含所有
	fmt.Println(slice[1:])	//第一个元素到末尾
	fmt.Println(slice[:3])	//到第三个元素
	fmt.Println("sss",slice[:len(arr)-1])	//去掉最后一个元素
	fmt.Println(slice)

}

func main(){
	testSlice()
}
*/
//切片拷贝

func testCopy() {
	var a []int = []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 1)
	copy(b, a)
	fmt.Println(b) //[0]
}

func testStrings() {
	var a = [...]string{"abc", "efg", "a"}
	sort.Strings(a[:])
	fmt.Println(a)
}
func testFloat() {
	var a = [...]float64{2.3, 3.2, 4.2, 0.8}
	sort.Float64s(a[:])
	fmt.Println(a)
}
func testIntSearch() {
	var a = [...]int{1, 2, 3, 4, 5}
	index := sort.SearchInts(a[:], 4)
	fmt.Println(index)
}

func main() {
	testStrings()
	testFloat()
	testIntSearch()
	//testCopy()
}

/*
[a abc efg]
[0.8 2.3 3.2 4.2]
3
*/
