package znet

type Message struct {
	ID   uint32
	Len  uint32
	Data []byte
}

func NewMessage(ID uint32, data []byte) *Message {
	return &Message{ID: ID, Len: uint32(len(data)), Data: data}
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(length uint32) {
	m.Len = length
}

func (m *Message) SetDataID(id uint32) {
	m.ID = id
}

func (m *Message) GetDataLen() uint32 {
	return m.Len
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetData() []byte {
	return m.Data
}
