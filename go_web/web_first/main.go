package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数

	// for k, v := range r.URL.Query() {
	// 	fmt.Println("key:", k, ", value:", v[0])
	// }

	// for k, v := range r.PostForm {
	// 	fmt.Println("key:", k, ", value:", v[0])
	// }
	t := template.New("index")

	t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")

	data := map[string]string{
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)

	// fmt.Fprintf(w, "这是一个开始")
}

func myWebv2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")

	data := map[string]string{
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)

}

func main() {
	http.HandleFunc("/", myWeb)
	http.HandleFunc("/index", myWebv2)

	fmt.Println("服务器即将开启, 访问地址 http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务器开启操作", err)
	}
}
