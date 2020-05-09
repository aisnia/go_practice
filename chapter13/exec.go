package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */

	env := os.Environ()

	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	//调用或启动外部系统命令和二进制可执行文件 如 ls -l的命令
	//它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
	//这个函数返回被启动进程的 id（pid），或者启动失败返回错误。
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

	//run的方式
	cmd := exec.Command("gedit")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error")
		os.Exit(0)
	}
}
