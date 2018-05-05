package msg_def

import (
	"github.com/Cyinx/einx/lua"
	"github.com/Cyinx/einx/slog"
	"github.com/Cyinx/protobuf/proto"
	"github.com/yuin/gopher-lua"
	"reflect"
)

type ProtoTypeID = uint32
type Message = proto.Message

type msgSerializer struct {
	lua *lua_state.LuaRuntime
}

var Serializer *msgSerializer = &msgSerializer{}

var MsgNewMap = make(map[ProtoTypeID]func() interface{})

func RegisterMsgProto(msg_type uint16, msg_id uint16, x Message) ProtoTypeID {
	proto_id := (uint32(msg_type) << 16) | uint32(msg_id)
	if proto_id == LuaMsgID {
		slog.LogInfo("msg_proto", "id %v is provide for lua.msg_type %v msg_id %v", LuaMsgID, msg_type, msg_id)
		return 0
	}
	t := reflect.TypeOf(x)
	f := proto.MessageNewFunc(t)
	if f == nil {
		slog.LogInfo("msg_proto", "unregister message new func [%s]", t.Name())
	} else {
		MsgNewMap[proto_id] = f
	}
	return proto_id
}

func (this *msgSerializer) SetLuaRuntime(l *lua_state.LuaRuntime) {
	this.lua = l
}

func (this *msgSerializer) UnmarshalMsg(type_id ProtoTypeID, data []byte) interface{} {
	if type_id == LuaMsgID {
		lua_msg, _ := lua_state.UnMarshal(data, this.lua.GetVm())
		return lua_msg
	}

	msg_new, ok := MsgNewMap[type_id]
	if !ok {
		return nil
	}

	msg := msg_new()
	proto.UnmarshalMerge(data, msg.(Message))
	return msg
}

func (this *msgSerializer) MarshalMsg(msg interface{}) ([]byte, error) {
	switch v := msg.(type) {
	case Message:
		return proto.Marshal(v)
	default:
		b := lua_state.Marshal(make([]byte, 0, 16), msg.(lua.LValue))
		return b, nil
	}
}

func (this *msgSerializer) UnmarshalRpc(type_id ProtoTypeID, data []byte) interface{} {
	return nil
}

func (this *msgSerializer) MarshalRpc(msg interface{}) ([]byte, error) {
	return nil, nil
}
