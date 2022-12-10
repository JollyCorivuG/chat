package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
	"go_web/chat/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("------------恭喜xxx登录成功------------")
	fmt.Println("\t\t 1 显示在线用户列表")
	fmt.Println("\t\t 2 发送信息")
	fmt.Println("\t\t 3 消息列表")
	fmt.Println("\t\t 4 消息列表")
	fmt.Println("\t\t 请选择(1 - 4)")

	var key int
	fmt.Scan(&key)

	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		fmt.Println("在线用户列表如下:")
	case 2:
		fmt.Println("发送信息")
		fmt.Println("你想对大家说什么:")
		var content string
		fmt.Scan(&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入")
	}
}

// 和服务器端保持通讯
func SeverProcessMes(c net.Conn) {
	tf := &utils.Transfer{
		C: c,
	}
	for {
		fmt.Println("客户端正在读取服务器发送的信息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err)
			return
		}

		switch mes.Type {
		case model.NotifyUserStatusMesType:
			var notifyUserStatusMes *model.NotifyUserStatusMes
			// time.Sleep(10 * time.Second)
			err = json.Unmarshal([]byte(mes.Data), notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal err =", err)
			}
			UpdateUserStatus(notifyUserStatusMes)
			OutPutOnlineUser()
		case model.SmsMesType:
			OutputGroupMes(&mes)
		default:
			fmt.Println("客户端收到了未知的信息类型")
		}
		fmt.Printf("mes: %v\n", mes)
	}
}
