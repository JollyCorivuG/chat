package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
	"go_web/chat/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *model.Message) (err error) {
	// 遍历服务器的map
	var smsMes model.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	for id, up := range Mgr.(*manager).OnlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.C)
	}
	return
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, c net.Conn) {
	tf := &utils.Transfer{
		C : c,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err =", err)
	}
}