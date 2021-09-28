package znet

import (
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {
	m1 := &Message{
		Id:      1,
		Data:    []byte("hello"),
		DataLen: 5,
	}
	m2 := &Message{
		Id:      2,
		Data:    []byte("nanase"),
		DataLen: 6,
	}

	dp := DataPack{}

	b1, _ := dp.Pack(m1)
	b2, _ := dp.Pack(m2)

	fmt.Println(b1)
	fmt.Println(b2)

	all := append(b1, b2...)
	fmt.Println(all)

	mm1, _ := dp.Unpack(all)
	fmt.Println(mm1)
}