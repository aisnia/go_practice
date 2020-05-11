package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64 //数
	avg     float64   //平均数
	mid     float64   //中位数
}

//html 的 表单
const form = `<html><body><form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></html></body>`

var pageTop = ""
var pageBottom = ""

const error = `<p class="error">%s</p>`

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(w http.ResponseWriter, req *http.Request) {
	//请求头设置 html媒体类型
	w.Header().Set("Content-Type", "text/html")

	err := req.ParseForm()

	//首页
	fmt.Fprint(w, pageTop, form)

	if err != nil {
		//错误页面
		fmt.Fprintf(w, error, err)
	} else {
		//请求成功  processRequest 处理请求
		if numbers, message, ok := processRequest(req); ok {
			stats := getStats(numbers)
			fmt.Fprint(w, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(w, error, message)
		}
	}
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	var text string
	// request.Form["numbers"] 解析name为numbers的表单，返回 数据，和是否存在
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		//处理如果网页中输入的是中文逗号，切分
		if strings.Contains(slice[0], "&#65292") {
			text = strings.Replace(slice[0], "&#65292;", " ", -1)
		} else {
			text = strings.Replace(slice[0], ",", " ", -1)
		}
		//遍历，取出两边空格
		for _, field := range strings.Fields(text) {
			//转为 float64，放入numbers
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}

	return numbers, "", true
}

//获取结构体，计算平均值，中值等
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	//排序
	sort.Float64s(stats.numbers)
	stats.avg = sum(numbers) / float64(len(numbers))
	stats.mid = median(numbers)
	return
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

//输入的html结果页面
func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.avg, stats.mid)
}
