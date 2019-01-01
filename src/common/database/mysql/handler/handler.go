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

var worker_pool = einx.GetWorkerPool("mysql")

func RegisterHandler(type_id ProtoTypeID, handler MsgHandler) {
	worker_pool.RegisterHandler(type_id, handler)
}

func RegisterRpcHandler(rpc_name string, handler RpcHandler) {
	worker_pool.RegisterRpcHandler(rpc_name, handler)
}
