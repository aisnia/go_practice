//同一个包下面的文件属于同一个工程文件，不要import就可直接使用
//而且 package名都一样，并且最后是与目录名一样，这里没有因为需要 main包运行
//命名 采用域名保证唯一
package main

//导包，有三种导包方式，必须使用到包，否则会报错，
/*import (
	."fmt"
)
这种可以直接省略前面的 包名前缀 如 Println("hello,word")

import (
	f "fmt"
)
这种是设置别名，可以如下使用：f.Println("hello,world")
import (
	_ "包名"
)
这种可以不使用，包里面的函数，但是会执行里面的 init函数*/

import "fmt"

//main函数，程序入口
func main() {
	fmt.Println("hello,world")
}
