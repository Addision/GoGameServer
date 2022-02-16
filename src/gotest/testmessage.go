package main

import (
	"fmt"
	. "network"
)

func main() {
	msg := NewMessage(123456, 1, []byte("egwegegreasge"))
	dp := NewDataPack()
	data, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack data error, ", err.Error())
		return
	}

	msg2, err := dp.UnPack(data)
	if err != nil {
		fmt.Println("unpack data error, ", err.Error())
		return
	}
	fmt.Println("datalen=", msg2.GetDataLen(), " msgId=", msg2.GetMsgID(), " msgType=", msg2.GetMsgType(), " data=", string(data[dp.GetHeadLen():]))
}
