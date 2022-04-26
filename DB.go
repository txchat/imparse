package imparse

import (
	"context"
	"github.com/txchat/imparse/proto"
)

type MsgIndex struct {
	Mid        string
	Seq        string
	SenderId   string
	CreateTime uint64
}

type Cache interface {
	GetMsg(ctx context.Context, from, seq string) (*MsgIndex, error)
	AddMsg(ctx context.Context, uid string, m *MsgIndex) error
	GetMid(ctx context.Context) (id int64, err error)
}

type DB interface {
	SaveMsg(deviceType proto.Device, frame Frame) error
}
