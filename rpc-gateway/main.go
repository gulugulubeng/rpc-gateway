package main

import (
	"fmt"
	"os"
	"os/signal"
	"rpc-gateway/http"
	"rpc-gateway/rpc"
	"syscall"
)

func main() {
	go rpc.StartRpcServer()
	go http.StartHttpServer(":8080")

	// 主进程阻塞
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	fmt.Printf("System stop by signal:%+v", <-signals)
	http.ShutdownHttpServer()
	rpc.ShutdownRpcServer()
}
