package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

const lenPath = len("/view/")

//正则表达式 验证标题
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

//html 模板
var templates = make(map[string]*template.Template)

var err error

//页面的数据，到时候加入模板中，有标题和消息体
type Page struct {
	Title string
	Body  []byte
}

//初始化 edit 与 view的模板
func init() {
	for _, tmpl := range []string{"edit", "view"} {
		//template.Must(template.ParseFiles(tmpl + ".html")) 把模板文件转换为 *template.Template 类型的对象，
		//即对应的 edit.html 和 view.html 模板进行解析
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}

func main() {
	//根据不同的路径做对应的处理
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	//监听 8080端口
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//返回一个闭包，用makeHandler 将传入的 Handler进行前置处理，即验证URL是否符合要求
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//获取标题，并验证
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		//实际调用传入的 handler进行处理
		fn(w, r, title)
	}
}

//
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//读取文本文件返回 Page
	p, err := load(title)
	if err != nil { // 没有改页面，重定向到 edit进行上传
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//读取文本文件返回 Page，
	p, err := load(title)
	if err != nil {
		//如歌没有改文件，则创建一个Page，设置好标题
		p = &Page{Title: title}
	}
	//渲染edit页面进行编辑
	renderTemplate(w, "edit", p)
}

//保存eidt编辑的信息
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//获取到body，然后包装成Page，并且save
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	//将body保存到 title.txt中
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//渲染页面
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// file created with read-write permissions for the current user only
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//文件是 title.txt格式的，返回Page
func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
