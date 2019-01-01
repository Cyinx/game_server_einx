package outserver_mgr

import (
	"msg_def"
)

type AuthServer struct {
	linker NetLinker
}

func (this *AuthServer) SendLuaMsg(msg interface{}) {
	b, _, _ := msg_def.MarshalMsg(msg)
	this.linker.WriteMsg(msg_def.LuaMsgID, b)
}

func (this *AuthServer) SendMsg(msg_id ProtoTypeID, msg interface{}) {
	b, _, _ := msg_def.MarshalMsg(msg)
	this.linker.WriteMsg(msg_id, b)
}
