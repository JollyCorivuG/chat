package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
	"go_web/chat/utils"
	"net"
	"os"
)

type UserProcess struct {
	// 暂时不需要字段
}

func (this *UserProcess) Login(userId int, userName string, userPwd string) (err error) {
	// 连接服务器
	c, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	// 延时关闭
	defer c.Close()

	// 准备通过c发送信息给服务端
	var mes model.Message
	mes.Type = model.LoginMesType

	// 创建一个LoginMes结构体
	var loginmes model.LoginMes
	loginmes.UserId = userId
	loginmes.UserName = userName
	loginmes.UserPwd = userPwd

	// 将loginmes序列化
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)

	// 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送登录消息时错误", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg err =", err)
		return
	}

	var loginResMes model.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		CurUser.C = c
		CurUser.UserName = loginmes.UserName
		CurUser.UserStatus = model.UserOnline
		// fmt.Println("登录成功")
		// 显示当前在线用户列表
		for _, v := range loginResMes.UsersId {
			if v == loginmes.UserId {
				continue
			}
			fmt.Printf("在线用户id: %v\n", v)
			user := &model.User{
				UserId:     v,
				UserStatus: model.UserOnline,
			}
			OnlineUsers[v] = user
		}
		// 这里我们还需要在客户端启动一个协程
		// 该协程保持和服务端的通讯，如果服务器有数据推送
		// 则接收并显示
		// 显示我们的登录菜单
		go SeverProcessMes(c)
		for {
			ShowMenu()
		}
	} else {
		fmt.Printf("loginResMes.Error: %v\n", loginResMes.Error)
		os.Exit(0)
	}

	return
}

func (this *UserProcess) Register(userId int, userName string, userPwd string) (err error) {
	c, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer c.Close()

	// 通过c发送信息给服务器
	var mes model.Message
	mes.Type = model.RegisterMesType
	// 创建一个LoginMes结构体
	loginmes := model.RegisterMes{}
	loginmes.User.UserId = userId
	loginmes.User.UserName = userName
	loginmes.User.UserPwd = userPwd

	// 将loginmes序列化
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册消息时错误", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg err =", err)
		return
	}

	var registerResMes model.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你可以去登录了！")
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
