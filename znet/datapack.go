package znet

import (
	"bytes"
	"encoding/binary"
	"github.com/0RAJA/zinx/ziface"
	"unsafe"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	return 2 * uint32(unsafe.Sizeof(uint32(1))) // 2*sizeof(uint32)
}

// Pack 封包
func (d *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// UnPack 此处的拆包是通过两步 进行的,先拆出长度,再根据长度进行进行读.直到读取完毕
func (d *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	dataBuff := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	return msg, nil
}
