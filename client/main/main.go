package main

import (
	"fmt"
	"go_web/chat/client/processor"
)

var userId int
var userName string
var userPwd string

func main() {
	// 接受用户的选择
	var key int

	for true {
		fmt.Println("------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1 - 3)")
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scan(&userId)
			fmt.Println("请输入用户名")
			fmt.Scan(&userName)
			fmt.Println("请输入用户密码")
			fmt.Scan(&userPwd)
			userprocess := &processor.UserProcess{}
			userprocess.Login(userId, userName, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scan(&userId)
			fmt.Println("请输入用户名")
			fmt.Scan(&userName)
			fmt.Println("请输入用户密码")
			fmt.Scan(&userPwd)
			userprocess := &processor.UserProcess{}
			userprocess.Register(userId, userName, userPwd)
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
