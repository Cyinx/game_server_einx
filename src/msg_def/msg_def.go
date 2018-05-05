package msg_def

import (
	"protobuf_gen"
)

type VersionCheck = pbgen.VersionCheck
type Login = pbgen.Login
type TestData = pbgen.TestData
type LuaRpcMsg = pbgen.LuaRpcMsg

var LuaMsgID ProtoTypeID = 1

var VersionCheckMsgID = RegisterMsgProto(uint16(pbgen.MainMsgID_GENERAL_MSG),
	uint16(pbgen.HandlerMsgID_VERSION_CHECK),
	(*VersionCheck)(nil))
