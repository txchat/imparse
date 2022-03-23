package imparse

import "context"

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
	SaveMsg(frame Frame) error
}
