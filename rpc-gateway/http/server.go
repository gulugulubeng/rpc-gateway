package http

import (
	"context"
	"net/http"
	"rpc-gateway/gw"
)

var (
	server http.Server
)

// StartHttpServer 异步开启HTTP服务
func StartHttpServer(addr string) error {
	// 将HTTP转换为RPC
	http.Handle("/", new(gw.Handler))
	// 初始化HTTP服务
	server = http.Server{
		Addr: addr,
	}
	// HTTP服务开始阻塞监听
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// ShutdownHttpServer 关闭HTTP服务
func ShutdownHttpServer() error {
	return server.Shutdown(context.Background())
}
