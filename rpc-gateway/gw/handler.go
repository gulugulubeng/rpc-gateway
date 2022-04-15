package gw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"rpc-gateway/rpc"
)

type Handler struct{}

func (gw Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 获得URL参数
	args := req.URL.Query()
	// 判断URL参数是否存在目标集群 以及目标集群是否存在
	clu := args.Get("cluster")
	if clu == "" || rpc.RpcClientMap[clu] == nil {
		rw.Write([]byte("目标集群不存在"))
		return
	}
	// 判断rpc请求函数
	if args.Get("rpcMethod") == "" {
		rw.Write([]byte("RPC函数无效"))
		return
	}
	// 读取body数据
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("读取body错误: %+v", err)))
		return
	}
	// RPC请求
	var reply Reply
	arg := &Args{
		Method: req.Method,
		Arg:    args.Encode(),
		Body:   bytes,
	}
	err = rpc.RpcClientMap[clu].Call(args.Get("rpcMethod"), arg, &reply)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(fmt.Sprintf("调用RPC错误: %+v", err)))
	}
	rw.Write(reply.Resp)
}

type Args struct {
	Method string `json:"method"`
	Arg    string `json:"args"`
	Body   []byte `json:"body"`
}

type Reply struct {
	Resp []byte `json:"resp"`
}
