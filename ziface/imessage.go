package ziface

type IMessage interface {
	GetId() uint32
	SetId(id uint32)

	GetDataLen() uint32
	SetDataLen(len uint32)

	GetData() []byte
	SetData(data []byte)
}