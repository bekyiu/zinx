package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx/util"
	"zinx/ziface"
)

type DataPack struct {
}

// dataLen uint32 + id uint32
func (dp *DataPack) GetHeaderLen() uint32 {
	return 8
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// |dataLen|id|data|
func (dp *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	// 将dataLen写入buffer
	if err := binary.Write(buffer, binary.LittleEndian, message.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, message.GetId()); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (dp *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	buffer := bytes.NewBuffer(data)

	// 读前8个字节
	msg := &Message{}
	// msg.DataLen的大小限制了读取的字节数
	if err := binary.Read(buffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(buffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 一个包长度太大
	if msg.DataLen > util.GlobalConfig.MaxPkgSize {
		return nil, errors.New(fmt.Sprintf("package is too large: %d byte", msg.DataLen))
	}

	return msg, nil
}
