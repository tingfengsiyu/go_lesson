package main

//map类型 k-v
import (
	"fmt"
	"sort"
)

func testMap() {
	var a map[string]string
	a = make(map[string]string, 10)
	// 同上两行 var a map[string]string =map[string]string{"key":"value",}
	a["abc"] = "efg"
	a["dd"] = "dd"
	fmt.Println(a)
}
func testMap2() {
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string, 100)
	a["key1"]["key2"] = "value"
	a["key1"]["key3"] = "value1"
	a["key1"]["key4"] = "value4"
	fmt.Println(a)
	//嵌套map map[key1:map[key2:value key3:value1 key4:value4]]

}

func modify(a map[string]map[string]string) {
	//修改值
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["passwd"] = "123456"
	a["zhangsan"]["name"] = "luofeng"
	/*
		if ok{
			val["passwd"] ="123456"
			val["name"] ="luofeng"
		}else{
			a["zhangsan"] = make(map[string]string)
			a["zhangsan"]["passwd"]="123456"
			a["zhangsan"]["name"]="luofeng"
		}*/
	return
}
func testMap3() {
	a := make(map[string]map[string]string, 100)
	modify(a)
	fmt.Println(a)
	//map[zhangsan:map[name:luofeng passwd:123456]]
}
func trans(a map[string]map[string]string) {
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}
}
func testMap4() {
	//map遍历
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string, 100)
	a["key1"]["key2"] = "value"
	a["key1"]["key3"] = "value1"
	a["key1"]["key4"] = "value4"
	a["key2"] = make(map[string]string, 100)
	a["key2"]["key2"] = "value"
	trans(a)
	delete(a, "key1")

	trans(a)
}

func testMaps() {
	var a []map[int]int

	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 100
	fmt.Println(a)
}

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[0] = 1
	a[1] = 2
	a[2] = 11
	a[3] = 12
	a[4] = 13
	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func testMapSort1() {
	//k变成v，v变成k
	var a map[string]int
	var b map[int]string

	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 1
	a["de"] = 2

	for k, v := range a {
		b[v] = k
	}
	fmt.Println(b) //map[1:abc 2:de]

}

func main() {
	testMap()
	testMap2()
	testMap3()
	testMap4()
	testMaps()
	testMapSort()
	testMapSort1()
}
