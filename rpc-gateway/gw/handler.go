package gw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"rpc-gateway/rpc"
)

type Handler struct{}

/**
协议转换HTTP请求转换为RPC请求

在HTTP请求头设置
	RPC服务地址：cluster
	RPC请求函数：rpcMethod

RPC请求设置
	请求参数为：*http.Request
	响应值接受：*http.ResponseWriter
*/
func (gw Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 从HTTP请求头获取RPC服务
	clu := req.Header.Get("cluster")
	if clu == "" || rpc.ClientMap[clu] == nil {
		rw.Write([]byte("目标RPC服务不存在"))
		return
	}
	// 从HTTP请求头获取rpc请求函数
	if req.Header.Get("rpcMethod") == "" {
		rw.Write([]byte("RPC函数无效"))
		return
	}

	// RPC请求
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf("读取Body错误: %+v", err)))
		return
	}
	args := &Args{
		Method:           req.Method,
		URL:              req.URL,
		Proto:            req.Proto,
		Header:           req.Header,
		Body:             bytes,
		ContentLength:    req.ContentLength,
		TransferEncoding: req.TransferEncoding,
		Host:             req.Host,
		Form:             req.Form,
		PostForm:         req.PostForm,
		MultipartForm:    req.MultipartForm,
		Trailer:          req.Trailer,
		RemoteAddr:       req.RemoteAddr,
		RequestURI:       req.RequestURI,
		TLS:              req.TLS,
		Response:         req.Response,
	}
	var reply Reply
	err = rpc.ClientMap[clu].Call(req.Header.Get("rpcMethod"), args, &reply)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(fmt.Sprintf("调用RPC错误: %+v", err)))
		return
	}
	rw.WriteHeader(reply.HttpStatusCode)
	for k, v := range reply.HeaderKeyValues {
		rw.Header().Set(k, v)
	}
	rw.Write(reply.Body)
}
