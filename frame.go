package imparse

import "context"

type FrameType string

type BizProto interface {
	AckBody() ([]byte, error)
	PushBody() ([]byte, error)
}

type Frame interface {
	Type() FrameType
	Filter(ctx context.Context, db Cache, filters ...Filter) (uint64, error)
	Transport(ctx context.Context, exec Exec) error
	Ack(ctx context.Context, exec Exec) (int64, error)
}
