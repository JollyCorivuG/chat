package dao

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Manager interface {
	GetUserById(c redis.Conn, id int) (user *model.User, err error)
	Login(userId int, userPwd string) (user *model.User, err error)
	Register(user *model.User) (err error)
}

var Mgr Manager

type manager struct {
	pool *redis.Pool
}

func InitPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	Mgr = &manager{
		pool: pool,
	}
}

func (this *manager) GetUserById(c redis.Conn, id int) (user *model.User, err error) {
	res, err := redis.String(c.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = model.ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &model.User{}
	// 把res反序列化成User对象
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}

func (this *manager) Login(userId int, userPwd string) (user *model.User, err error) {
	c := this.pool.Get()
	defer c.Close()
	user, err = this.GetUserById(c, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = model.ERROR_USER_WRONGPWD
		return
	}
	return
}

func (this *manager) Register(user *model.User) (err error) {
	c := this.pool.Get()
	defer c.Close()
	_, err = this.GetUserById(c, user.UserId)
	if err == nil {
		err = model.ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	_, err = c.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("用户注册时发生错误", err)
		return
	}
	return
}
