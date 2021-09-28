package znet

type Message struct {
	Id      uint32
	Data    []byte
	DataLen uint32
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		Data:    data,
		DataLen: uint32(len(data)),
	}
}

func (m *Message) GetId() uint32 {
	return m.Id
}

func (m *Message) SetId(id uint32) {
	m.Id = id
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
