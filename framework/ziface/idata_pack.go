package ziface

type IDataPack interface {
	// header长度
	GetHeaderLen() uint32
	// 封包
	Pack(message IMessage) ([]byte, error)
	// 拆包
	Unpack([]byte) (IMessage, error)
}