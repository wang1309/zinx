package main

import (
	"zinx/znet"
)


func main() {
	//创建一个server句柄
	s := znet.NewServer("zinx")

	//开启服务
	s.Serve()
}
