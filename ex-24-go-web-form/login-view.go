package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url傳遞的參數，對於POST則解析響應包的主體（request body）
	//注意:如果沒有調用ParseForm方法，下面無法獲取表單的數據
	fmt.Println(r.Form) //這些信息是輸出到服務器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Go Web!") //這個寫入到w的是輸出到客戶端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //獲取請求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//請求的是登陸數據，那麼執行登陸的邏輯判斷
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //設置訪問的路由
	http.HandleFunc("/login", login)         //設置訪問的路由
	err := http.ListenAndServe(":9090", nil) //設置監聽的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
