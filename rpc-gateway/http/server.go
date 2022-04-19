package http

import (
	"context"
	"flag"
	"log"
	"net/http"
	"rpc-gateway/gw"
)

var (
	server   http.Server
	httpAddr string
)

// 命令行参数HTTPAddr=localhost:80设置HTTP监听端口
func init() {
	flag.StringVar(&httpAddr, "HTTPAddr", "localhost:80", "HTTP监听地址")
}

// StartHttpServer 开启HTTP服务
func StartHttpServer() {
	// 将HTTP转换为RPC
	http.Handle("/", new(gw.Handler))
	// 初始化HTTP服务
	server = http.Server{
		Addr: httpAddr,
	}
	// 异步开启HTTP服务开始阻塞监听
	go func() {
		log.Printf("[StartHttpServer]HTTP server addr %s", httpAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("[StartHttpServer]HTTP server listen and serve return %+v", err)
		}
	}()
}

// ShutdownHttpServer 关闭HTTP服务
func ShutdownHttpServer() error {
	return server.Shutdown(context.Background())
}
