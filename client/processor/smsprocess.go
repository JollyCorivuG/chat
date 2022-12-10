package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
	"go_web/chat/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes model.Message
	mes.Type = model.SmsMesType

	var smsMes model.SmsMes
	smsMes.Content = content
	smsMes.UserName = CurUser.UserName
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
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
		C : CurUser.C,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err =", err)
	}
	return
}