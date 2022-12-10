package processor

import (
	"fmt"
	"go_web/chat/model"
)

var OnlineUsers map[int]*model.User = make(map[int]*model.User, 1024)
var CurUser model.CurUser // 在用户登录成功后完成对CurUser的初始化

func OutPutOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for _, user := range OnlineUsers {
		fmt.Println("用户id:/t", user.UserId)
	}
}

func UpdateUserStatus(notifyUserStatusMes *model.NotifyUserStatusMes) {
	user, ok := OnlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &model.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	// user.UserStatus = notifyUserStatusMes.Status
	OnlineUsers[notifyUserStatusMes.UserId] = user
}
