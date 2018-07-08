package msg_def

import (
	"github.com/Cyinx/einx/lua"
	"github.com/Cyinx/einx/network"
	"github.com/Cyinx/einx/slog"
	"github.com/Cyinx/protobuf/proto"
	"github.com/yuin/gopher-lua"
	"reflect"
	//	"sync"
)

type ProtoTypeID = uint32
type Message = proto.Message

var lua_vm *lua_state.LuaRuntime = nil

var MsgNewMap = make(map[ProtoTypeID]func() interface{})

func SetLuaVm(vm *lua_state.LuaRuntime) {
	lua_vm = vm
}

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

func UnmarshalMsg(type_id ProtoTypeID, data []byte) interface{} {
	if type_id == LuaMsgID {
		lua_msg, _ := lua_state.UnMarshal(data, lua_vm.GetVm())
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

func MarshalMsg(msg interface{}) ([]byte, error, bool) {
	switch v := msg.(type) {
	case Message:
		b, err := proto.Marshal(v)
		return b, err, false
	default:
		b := make([]byte, 0, 128)
		b = lua_state.Marshal(b, msg.(lua.LValue))
		return b, nil, true
	}
}

func UnmarshalRpc(type_id ProtoTypeID, data []byte) interface{} {
	p, _ := network.RpcUnMarshal(data)
	return p
}

func MarshalRpc(msg interface{}) ([]byte, error, bool) {
	b := make([]byte, 0, 128)
	b = network.RpcMarshal(b, msg)
	return b, nil, true
}
