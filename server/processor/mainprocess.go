package processor

import (
	"fmt"
	"go_web/chat/model"
	"go_web/chat/utils"
	"net"
)

type MainProcess struct {
	C net.Conn
}

// 编写一个SeverProcessMes函数
func (this *MainProcess) SeverProcessMes(mes *model.Message) (err error) {
	fmt.Printf("mes: %v\n", mes)
	switch mes.Type {
	case model.LoginMesType:
		// 处理登录
		up := &UserProcess{
			C: this.C,
		}
		err = up.SeverProcessLogin(mes)
	case model.RegisterMesType:
		// 处理注册
		up := &UserProcess{
			C: this.C,
		}
		up.SeverProcessRegister(mes)
	case model.SmsMesType:
		sp := &SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理....")
	}
	return
}

// 处理和客户端的通讯
func (this *MainProcess) Start() (err error) {

	// 读取客户端发送的数据
	for {
		// 读包
		tf := &utils.Transfer{
			C: this.C,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg err =", err)
			return err
		}

		err = this.SeverProcessMes(&mes)
		if err != nil {
			fmt.Println("severProcessMes err =", err)
			return err
		}
	}
}
