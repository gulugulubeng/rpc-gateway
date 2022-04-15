package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go StartRPCServer()

	// 主进程阻塞
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	fmt.Printf("System stop by signal:%+v", <-signals)
}
