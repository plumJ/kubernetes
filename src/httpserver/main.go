package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func main() {
	// 设置多路复用处理函数
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", health)
	http.HandleFunc("/test", health)

	// 设置监听端口
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println("Error Listening 80 port", err.Error())
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	// Header
	w.WriteHeader(200)

	// ResponseBody
	_, _ = w.Write([]byte("<h1>Health</h1>"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	// 读取request header
	headers := r.Header
	for header, param := range headers {
		// 重写response header
		var temp string
		for _, i := range param {
			temp = i
			w.Header().Set("client_"+header, temp)
		}
	}

	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("VERSION", runtime.Version())

	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	httpCode := 200
	w.WriteHeader(httpCode)

	// Log
	log := fmt.Sprintf("[ROOT] IP:%s, HTTP_CODE:%d \n", r.RemoteAddr, httpCode)
	_, _ = io.WriteString(os.Stdout, log)

	// ResponseBody
	_, _ = w.Write([]byte("<h1>HOME</h1>"))
}