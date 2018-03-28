package msg_def

import (
	"github.com/Cyinx/einx/network"
	"protobuf_gen"
)

type VersionCheck = pbgen.VersionCheck
type Login = pbgen.Login
type TestData = pbgen.TestData
type LuaRpcMsg = pbgen.LuaRpcMsg

var VersionCheckMsgID = network.RegisterMsgProto(uint16(pbgen.MainMsgID_GENERAL_MSG),
	uint16(pbgen.HandlerMsgID_VERSION_CHECK),
	(*VersionCheck)(nil))

//var LoginMsgID = network.RegisterMsgProto(1, 1, (*Login)(nil))
//var TestDataMsgID = network.RegisterMsgProto(1, 2, (*TestData)(nil))

var MsgLuaRpcMsgID = network.RegisterMsgProto(uint16(pbgen.MainMsgID_GENERAL_MSG),
	uint16(pbgen.HandlerMsgID_LUARPC_REQUEST),
	(*LuaRpcMsg)(nil))
