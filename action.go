package imparse

import (
	"context"
	"github.com/txchat/imparse/proto/auth"
)

type Channel int

const (
	Undefined Channel = 0
	UniCast   Channel = 1
	GroupCast Channel = 2
)

var ChannelTypeName = map[Channel]string{
	Undefined: "Undefined",
	UniCast:   "UniCast",
	GroupCast: "GroupCast",
}

func (ch Channel) String() string {
	return ChannelTypeName[ch]
}

type Answer interface {
	Check(context.Context, Checker, Frame) error
	Filter(context.Context, Frame) (uint64, error)
	Transport(context.Context, Frame) error
	Ack(context.Context, Frame) (int64, error)
}

type Pusher interface {
}

type Storage interface {
	SaveMsg(context.Context, auth.Device, Frame) error
}
