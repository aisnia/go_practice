package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		fmt.Fprint(os.Stdout, "%s", buf)
	}
	return
}

//带行数的输出  才有 -n 输出行数      参数是 n   默认 不带false，   描述
var numberFlag = flag.Bool("n", false, "number each line")

func catNum(r *bufio.Reader) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		//如果带有 -n 的话则 会在前面加上 i 作为行号打印，
		if *numberFlag {
			fmt.Fprint(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprint(os.Stdout, "%s", buf)
		}
		return
	}
}

//使用切片作为缓冲
func cat2(f *os.File) {
	//缓冲 512字节
	const NBUF = 512
	var buf [NBUF]byte
	for {
		//每次读满切片，nr为字节数
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			//nw 写入的字节数
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
func main() {
	//参数解析
	flag.Parse()
	//没有参数，则输入什么就输出什么
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdout))
	}

	for i := 0; i < flag.NArg(); i++ {
		//遍历每个参数, 打开文件
		f, err := os.Open(flag.Arg(i))

		if err != nil {
			//错误流，  os.Args[0]是程序的名称，
			fmt.Fprintf(os.Stderr, "%s error reading from %s : %s\n", os.Args[0], flag.Arg(i), err.Error())
		}
		catNum(bufio.NewReader(f))
		//cat2(f)
	}
}
