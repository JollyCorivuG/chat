package utils

import (
	"encoding/binary"
	"encoding/json"
	"go_web/chat/model"
	"log"
	"net"
)

type Transfer struct {
	C net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes model.Message, err error) {
	
	_, err = this.C.Read(this.Buf[:4])
	if err != nil {
		log.Fatal(err)
	}

	// 根据this.Buf[:4]转换成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.C.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		log.Fatal(err)
	}

	// 将this.Buf[:pkgLen]反序列化 -> mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	// 先发送长度
	n, err := this.C.Write(this.Buf[:4])
	if n != 4 || err != nil {
		log.Fatal(err)
	}

	// 发送消息本身
	n, err = this.C.Write(data)
	if n != int(pkgLen) || err != nil {
		log.Fatal(err)
	}
	return
}
