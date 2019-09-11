package main

import (
	"net"
	"fmt"
)

/*
tcp-server
通过telnet与本服务端通信
*/
const (
	ip   = ""
	port = 1985
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
		Zone: "",
	})
	if err != nil {
		fmt.Println("监听端口失败！", err.Error())
	}
	fmt.Println("已经初始化连接，等待客户端连接...")
	Server(listen)
}

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close() //? Possible resource leak, defer is call in a for loop
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data) //每次为什么只读到一个字节？
				fmt.Println("客服端发来数据:", string(data[0:i]))
				if err != nil {
					fmt.Println("读取客户端数据错误:", err.Error())
					break
				}
				conn.Write([]byte("finish"))
			}
		}()
	}
}
