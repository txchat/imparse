package imparse

import "errors"

var (
	ErrorEnvType = errors.New("unsupported event type")
	ErrorChType  = errors.New("unsupported channel type")
	ErrorMsgType = errors.New("unsupported message type")

	ErrExecSupport          = errors.New("biz proto exec not support")
	ErrMsgTypeSupport       = errors.New("biz msg type not support")
	ErrAppTypeSupport       = errors.New("app type not support")
	ErrPermissionDeny       = errors.New("permission deny")
	ErrGroupMemberNotExists = errors.New("group member not exists")
)
