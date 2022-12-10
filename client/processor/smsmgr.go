package processor

import (
	"encoding/json"
	"fmt"
	"go_web/chat/model"
)

func OutputGroupMes(mes *model.Message) {
	var smsMes model.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	info := fmt.Sprintf("用户%s:\t 对大家说：\t%s", smsMes.UserName, smsMes.Content)
	fmt.Printf("info: %v\n", info)
}
