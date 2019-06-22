package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb)

	fmt.Println("服务器即将开启, 访问地址 http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务器开启操作", err)
	}
}
