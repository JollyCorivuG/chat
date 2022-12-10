package main

import (
	"fmt"
	"go_web/chat/dao"
	"go_web/chat/server/processor"
	"log"
	"net"
	"time"
)

// 处理和客户端的通讯
func process(c net.Conn) {
	defer c.Close()

	// 创建一个总控
	mainprocess := &processor.MainProcess{
		C: c,
	}
	err := mainprocess.Start()
	if err != nil {
		fmt.Println("服务端与客户端通讯错误")
		return
	}
}

// 需要事先初始化的东西
func init() {
	dao.InitPool("localhost:6379", 16, 9, 300*time.Second)
}

func main() {
	// 提示信息
	fmt.Println("服务器正在8889端口监听....")
	l, err := net.Listen("tcp", "0.0.0.0:8889")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器....")
		c, err2 := l.Accept()
		if err2 != nil {
			fmt.Println("l.Accept err =", err)
		}
		// 一旦连接成功，则启动一个协程和客户端保持通讯
		go process(c)
	}
}
