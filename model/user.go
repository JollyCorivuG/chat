package model

import "net"

// 定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

// 用户的结构体
type User struct {
	UserId     int    `json:"userid"`   // 用户id
	UserName   string `json:"username"` // 用户名
	UserPwd    string `json:"userpwd"`  // 用户密码
	UserStatus int    `json:"userstatus"`
}



type CurUser struct {
	C net.Conn
	User
}