package module

import (
	"github.com/Cyinx/einx"
)

var Instance = einx.GetModule("logic")
var db_module = einx.GetModule("mongodb")
var Lua = einx.NewLuaStae()

func AddTcpServer(addr string, mgr interface{}) {
	einx.AddTcpServerMgr(Instance, addr, mgr)
}

func StartTcpClient(addr string, mgr interface{}) {
	einx.StartTcpClientMgr(Instance, addr, mgr)
}

func InitLuaApi() {
	Lua.RegisterFunction("RpcCall", RpcCall)
	registerPersonType(Lua.GetVm())
}
