package ziface

/*
将请求的消息封装到message中,定义抽象层接口
*/

type IMessage interface {
	GetDataLen() uint32
	GetMsgID() uint32
	GetData() []byte

	SetData([]byte)
	SetDataLen(length uint32)
	SetDataID(id uint32)
}
