package msghandler

import (
	"game_server/logic/module"
	"github.com/Cyinx/einx"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type ProtoTypeID = einx.ProtoTypeID
type MsgHandler = einx.MsgHandler
type RpcHandler = einx.RpcHandler
type TimerHandler = einx.TimerHandler
type Context = einx.Context
type ModuleRouter = einx.ModuleRouter

var module_router = module.Instance.(ModuleRouter)

func RegisterHandler(type_id ProtoTypeID, handler MsgHandler) {
	module_router.RegisterHandler(type_id, handler)
}

func RegisterRpcHandler(rpc_name string, handler RpcHandler) {
	module_router.RegisterRpcHandler(rpc_name, handler)
}

func AddTimer(delay uint64, op TimerHandler, args ...interface{}) uint64 {
	return module.Instance.AddTimer(delay, op, args...)
}

func RemoveTimer(timer_id uint64) {
	module.Instance.RemoveTimer(timer_id)
}
