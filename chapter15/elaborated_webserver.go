package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

//包 expvar 可以创建（Int，Float 和 String 类型）变量，并将它们发布为公共变量。
//通常作为计数器
var helloRequests = expvar.NewInt("hello-requests")

// webroot web的根路径，root
var webroot = flag.String("root", "/home/user", "web root directory")

// simple flag server
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

//计数器结构体
type Counter struct {
	n int
}

//计数器对象 ctr 有一个 String() 方法，所以它实现了 expvar.Var 接口。这使其可以被发布，
func (ctr *Counter) String() string { return fmt.Sprintf("%d", ctr.n) }

//ServeHTTP 函数使 ctr 成为处理器，因为它的签名正确实现了 http.Handler 接口。
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	//get请求则递增，post请求则是设置值
	case "GET": // increment n
		ctr.n++
	case "POST": // set n to posted value
		buf := new(bytes.Buffer)
		io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// a channel
type Chan chan int

//Chan 也实现了http.Handler 的接口，可以处理请求，就是从通道取出数据并且返回
func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}
func main() {
	flag.Parse()
	//Logger 处理器 监听所有请求
	http.Handle("/", http.HandlerFunc(Logger))

	//hello处理器
	http.Handle("/go/hello", http.HandlerFunc(HelloServer))

	// 直接发布，变成公共变量，并且直接作为Handler接口 处理请求
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)

	//FileServer(root FileSystem) Handler 返回一个处理器，它以 root 作为根，用文件系统的内容响应 HTTP 请求。
	// 要获得操作系统的文件系统，用 http.Dir，例如：
	//http://localhost:12345/go/ggg.html  获取/tmp/ggg.html 的这个文件
	http.Handle("/go/", http.FileServer(http.Dir("/tmp"))) // uses the OS filesystem
	//另一种方式 /go/ 对应webroot
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))

	http.Handle("/flags", http.HandlerFunc(FlagServer))

	//命令行参数
	http.Handle("/args", http.HandlerFunc(ArgServer))

	//Chan 也实现了http.Handler 的接口，可以处理请求，就是从通道取出数据并且返回
	http.Handle("/chan", ChanCreate())

	//显示当前时间
	http.Handle("/date", http.HandlerFunc(DateServer))

	//监听12345 端口
	err := http.ListenAndServe(":12345", nil)
	//发生错误直接 panic
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

//日志，返回404的头部，和oops给浏览器
func Logger(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	//加1
	helloRequests.Add(1)
	io.WriteString(w, "hello, world!\n")
}

func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	//该处理函数使用了 flag 包。VisitAll 函数迭代所有的标签（flag），打印它们的名称、值和默认值（当不同于“值”时）。
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

//该处理函数迭代 os.Args 以打印出所有的命令行参数。如果没有指定则只有程序名称（可执行程序的路径）会被打印出来。即Args[0]
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

//返回一个 chan int，里面有协程 将x写入c通道
func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

// 显示当前的时间
func DateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}

	//调用的是操作系统的 /bin/date 命令
	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}
