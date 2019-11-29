package main

import "fmt"

var a = "G"

func n() {
	fmt.Println(a)
}
func m() {
	a := 'A'
	fmt.Println(a) //65 ASCII码字符集
}
func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprint("%c", str[strLen-i-1])
	}
	return result
}
func reverse1(str string) string {
	var result []byte //数组  []类型 字节
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1]) //数组追加
	}
	return string(result)
}
func main() {
	result1 := reverse("aaaaa")
	result := reverse1("ssss")

	fmt.Println(result1)
	fmt.Println(result)
	n()
	m()
	n()

}
