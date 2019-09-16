package main

import (
	"fmt"
	"io"
	"net"
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
		go func() {
			defer func() {
				conn.Close()
				fmt.Println("===========")
			}()
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data) //每次为什么只读到一个字节？
				fmt.Printf("%v", conn)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("client %s is close!\n", conn.RemoteAddr().String())
						//在这里直接退出goroutine，关闭由defer操作完成
						return
					}

					fmt.Println("读取客户端数据错误:", err.Error())
					break
				}
				fmt.Println("客服端发来数据:", string(data[0:i]))
				conn.Write([]byte("."))
			}
		}()
	}
}
