package add
var Name string = "hello world"
var Age int = 22
//小写age 无法被其他包文件引用，因为是私有变量

func init() {
	//初始化函数，执行在main之前
	Name = "test"
	Age = 11
}

