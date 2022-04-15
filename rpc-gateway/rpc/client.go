package rpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

var (
	RpcClientMap = map[string]*rpc.Client{}
	clusters     = []string{
		"localhost:1001",
	}
)

func StartRpcServer() error {
	for _, cluster := range clusters {
		client, err := jsonrpc.Dial("tcp", cluster)
		if err != nil {
			return err
		}
		RpcClientMap[cluster] = client
	}
	return nil
}

func ShutdownRpcServer() error {
	for _, client := range RpcClientMap {
		err := client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
