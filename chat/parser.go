package chat

import (
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	comet "github.com/txchat/im/api/comet/grpc"
	"github.com/txchat/imparse"
	biz "github.com/txchat/imparse/proto"
)

type _handleEvent func(*StandardFrame, []byte) (imparse.Frame, error)

var events = map[biz.Proto_EventType]_handleEvent{
	biz.Proto_common: func(s *StandardFrame, body []byte) (imparse.Frame, error) {
		var pro biz.Common
		err := proto.Unmarshal(body, &pro)
		if err != nil {
			return nil, err
		}
		switch pro.ChannelType {
		case biz.Channel_ToUser:
			return NewPrivateFrame(s, &pro), nil
		case biz.Channel_ToGroup:
			return NewGroupFrame(s, &pro), nil
		default:
			return nil, imparse.ErrExecSupport
		}
	},
	biz.Proto_Signal: func(s *StandardFrame, body []byte) (imparse.Frame, error) {
		var pro biz.Signal
		err := proto.Unmarshal(body, &pro)
		if err != nil {
			return nil, err
		}
		return NewNoticeFrame(s, &pro), nil
	},
}

//标准解析器 定义了解析标准协议的方法
type StandardParse struct {
}

func (s *StandardParse) NewFrame(key, from string, in io.Reader, opts ...Option) (imparse.Frame, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	//
	var p comet.Proto
	err = proto.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	switch p.GetVer() {
	case 0, 1:
	default:
	}

	switch p.GetOp() {
	case int32(comet.Op_SendMsg):
	case int32(comet.Op_Auth):

	}
	//业务服务协议解析
	var pro biz.Proto
	err = proto.Unmarshal(p.Body, &pro)
	if err != nil {
		return nil, err
	}

	//解析event事件
	event, ok := events[pro.EventType]
	if !ok || event == nil {
		return nil, imparse.ErrorEnvType
	}
	frame, err := event(NewStandardFrame(&p, key, from, opts...), pro.Body)
	if err != nil {
		return nil, err
	}
	return frame, err
}
