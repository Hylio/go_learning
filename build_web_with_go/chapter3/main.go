package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // 把参数打印到服务器端
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问路由
	err := http.ListenAndServe(":9090", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
