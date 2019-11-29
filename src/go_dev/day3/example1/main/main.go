package main

import "fmt"

/*
func urlProcess(url string ) string{
	result := strings.HasPrefix(url,"http://")	//判断一个url是否以http://开头，不是就加上
	if !result{
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}
func pathProcess(path string) string {
	result := strings.HasSuffix(path,"/")	//判断一个路径是否以/结尾，不是加上/
	if !result{
		path = fmt.Sprintf("path:/%s",path)
	}
	return path
}
func main(){
	var (
		url string
		path string
	)
	fmt.Scanf("%s%s",&url,&path)	//终端输入

	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Println(url,path)
}
//>main.exe
//10.0.0.1 dd
//http://10.0.0.1 /dd/


func main(){
	str := " hello world  \n"
	result := strings.Replace(str,"world","you",1)
	fmt.Println(result)

	count := strings.Count(str,"1")
	fmt.Println(count)

	result = strings.Repeat(str,3)
	fmt.Println("ss",result)

	result = strings.ToLower(str)
	fmt.Println(result)

	result = strings.ToUpper(str)
	fmt.Println(result)

	result = strings.TrimSpace(str)
	fmt.Println(result)

	result = strings.Trim(str,"hello")
	fmt.Println("ss",result)

	s := strings.Fields(str)	//返回数组
	for i :=0 ; i <len(s);i++{
		fmt.Println(s[i])
	}
	d := strings.Split(str,"1")
	for i :=0 ; i <len(d);i++{
		fmt.Println(d[i])
	}
	str2 := strings.Join(d,"1")
	fmt.Println(str2)

	str2 = strconv.Itoa(1000)
	fmt.Println(str2)

	number,error := strconv.Atoi(str2)
	if error != nil {
		fmt.Println("can not ",error)
	}
	fmt.Println(number,error)

}
*/
/*
func test(){
	time.Sleep(time.Microsecond * 100)
}
func main(){
	now := time.Now()
	fmt.Println("当前时间：",now.Format("2006/01/02 15:04:05"))	//格式必须是"2006/01/02 15:04:05"  结果是 2019/11/29 14:57:56

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d us\n",(end-start)/1000)
	//当前时间： 2019/11/29 15:01:21
	//cost:1024 us
}



func main(){
	var a int = 10
	fmt.Println(&a) //0xc0000100b8

	var p *int
	p = &a
	fmt.Println(*p)		//10
	*p = 100
	fmt.Println(a)		//100
}

func main(){
	var str string
	fmt.Scanf("%s",&str)
	number,err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("convert failed''")
		return
	}
	fmt.Println(number)
}


func main(){
	var a int =10
	switch a{
	case 0,1,2,3:
		print('0')
	default:
		print("ss")
	}
}


func main(){
	str := "hello world 中国"
	for index,val := range str{
		fmt.Printf("index[%d] val[%c] len[%d]\n", index,val, len([]byte(string(val))))

	}
}*/
type op_func func(int, int) int //自定义类型
func add(a, b int) int {
	return a + b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	sum := operator(c, 100, 200)
	fmt.Println(sum) //300
}
