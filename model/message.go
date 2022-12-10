package model

// 消息的类型
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 向服务器发送的消息
type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息本身
}

// 用户登录的消息
type LoginMes struct {
	UserId   int    `json:"userid"`   // 用户id
	UserName string `json:"username"` // 用户名
	UserPwd  string `json:"userpwd"`  // 用户密码
}

// 服务器返回的用户登录消息
type LoginResMes struct {
	Code    int    `json:"code"`  // 状态码，500表示用户未注册，200表示登录成功
	Error   string `json:"error"` // 返回错误信息
	UsersId []int  `json:"usersid"`
}

// 用户注册的消息
type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 状态码，200表示注册成功
	Error string `json:"error"` // 返回错误信息
}

// 为了配合服务器推送用户状态变化的信息
type NotifyUserStatusMes struct {
	UserId int `json:"userid"`
	Status int `json:"status"`
}

// 发送的信息
type SmsMes struct {
	User    `json:"user"` // 匿名结构体，指发送消息的用户
	Content string        `json:"content"`
}
