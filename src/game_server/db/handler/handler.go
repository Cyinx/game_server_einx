package msghandler

import (
	"game_server/db/module"
	"github.com/Cyinx/einx"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type ProtoTypeID = einx.ProtoTypeID
type MsgHandler = einx.MsgHandler
type RpcHandler = einx.RpcHandler

func RegisterHandler(type_id ProtoTypeID, handler MsgHandler) {
	module.Instance.RegisterHandler(type_id, handler)
}

func RegisterRpcHandler(rpc_name string, handler RpcHandler) {
	module.Instance.RegisterRpcHandler(rpc_name, handler)
}
