package msghandler

import (
	"game_server/db/module"
	logic "game_server/logic/module"
	"github.com/Cyinx/einx"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type ProtoTypeID = einx.ProtoTypeID
type MsgHandler = einx.MsgHandler
type RpcHandler = einx.RpcHandler

var logic_module = logic.Instance

func RegisterHandler(type_id ProtoTypeID, handler MsgHandler) {
	module.Instance.RegisterHandler(type_id, handler)
}

func RegisterRpcHandler(rpc_name string, handler RpcHandler) {
	module.Instance.RegisterRpcHandler(rpc_name, handler)
}
