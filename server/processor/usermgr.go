package processor

import (
	"fmt"
)

type Manager interface {
	AddOnlineUser(up *UserProcess)
	DelOnlineUser(up *UserProcess)
	GetAllOnlineUser() map[int]*UserProcess
	GetOnlineUserById(userId int) (up *UserProcess, err error)
}

var Mgr Manager

type manager struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	Mgr = &manager{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (this *manager) AddOnlineUser(up *UserProcess) {
	this.OnlineUsers[up.UserId] = up
}

func (this *manager) DelOnlineUser(up *UserProcess) {
	delete(this.OnlineUsers, up.UserId)
}

func (this *manager) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUsers
}

func (this *manager) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.OnlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
