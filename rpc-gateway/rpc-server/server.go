package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// StartRPCServer RPC服务阻塞监听
func StartRPCServer() {
	err := rpc.Register(new(RpcFun))
	if err != nil {
		log.Println(err.Error())
		return
	}
	// json rpc 服务端
	lis, err := net.Listen("tcp", "127.0.0.1:1001")
	if err != nil {
		log.Println(err)
		return
	}
	//服务端等待请求
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		// 并发处理客户端请求
		go jsonrpc.ServeConn(conn)
	}
}

type RpcFun string

func (*RpcFun) Echo(args *Args, reply *Reply) error {
	reply.Resp = []byte(fmt.Sprintf("RPC server accept %+v from client", args))
	return nil
}

type Args struct {
	Method string `json:"method"`
	Arg    string `json:"args"`
	Body   []byte `json:"body"`
}

type Reply struct {
	Resp []byte `json:"resp"`
}
