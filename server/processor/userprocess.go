package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/dao"
	"go_web/chat/model"
	"go_web/chat/utils"
	"net"
)

type UserProcess struct {
	C      net.Conn
	UserId int
}

// 编写通知所有在线用户的方法
func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	for _, up := range Mgr.(*manager).OnlineUsers {
		if up.UserId == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes model.Message
	mes.Type = model.NotifyUserStatusMesType

	var notifyUserStatusMes model.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = model.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
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
		C: this.C,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err =", err)
	}
}

// 处理登录
func (this *UserProcess) SeverProcessLogin(mes *model.Message) (err error) {
	var loginMes model.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	// 先声明一个resMes
	var resMes model.Message
	resMes.Type = model.LoginResMesType

	var loginResMes model.LoginResMes

	user, err := dao.Mgr.Login(loginMes.UserId, loginMes.UserPwd)
	// fmt.Println("debug")

	if err == nil {
		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		Mgr.AddOnlineUser(this)
		this.NotifyOtherOnlineUser(loginMes.UserId)
		for _, up := range Mgr.GetAllOnlineUser() {
			loginResMes.UsersId = append(loginResMes.UsersId, up.UserId)
		}

		fmt.Println(user.UserName, "登录成功")

	} else {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_WRONGPWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误...."
		}
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	// 我们创建一个Transfer
	tf := &utils.Transfer{
		C: this.C,
	}
	err = tf.WritePkg(data)
	return
}

// 处理注册
func (this *UserProcess) SeverProcessRegister(mes *model.Message) (err error) {
	var registerMes model.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	var resMes model.Message
	resMes.Type = model.RegisterResMesType
	var registerResMes model.RegisterResMes
	err = dao.Mgr.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册时发生未知错误...."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	// 我们先创建一个Transfer
	tf := &utils.Transfer{
		C: this.C,
	}

	err = tf.WritePkg(data)
	return
}
