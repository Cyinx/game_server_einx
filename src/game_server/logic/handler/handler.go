package msghandler

import (
	"github.com/Cyinx/einx"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type ProtoTypeID = einx.ProtoTypeID
type MsgHandler = einx.MsgHandler
type RpcHandler = einx.RpcHandler
type Context = einx.Context

var logic_module = einx.GetModule("logic")
var module_router = logic_module.(einx.ModuleRouter)

func RegisterHandler(type_id ProtoTypeID, handler MsgHandler) {
	module_router.RegisterHandler(type_id, handler)
}

func RegisterRpcHandler(rpc_name string, handler RpcHandler) {
	module_router.RegisterRpcHandler(rpc_name, handler)
}
