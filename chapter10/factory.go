package main

import "fmt"

//file 小写设置为私有的 可强制使用工厂方法创建
type File struct {
	fd   int    //文件标识
	name string //文件名
}

func newFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
func main() {
	f1 := newFile(1, "1")
	f2 := newFile(2, "2")
	fmt.Println(f1)
	fmt.Println(f2)

	/*试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，
	但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！ 因为 new( map[string]string) 返回的是一个指向 nil 的指针，
	它尚未被分配内存。所以在使用 map 时要特别谨慎。*/
}
