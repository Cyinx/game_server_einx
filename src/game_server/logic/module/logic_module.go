package module

import (
	"github.com/Cyinx/einx"
)

var Instance = einx.GetModule("logic")
var mongodb = einx.GetModule("mysql")
var Lua = einx.NewLuaStae()

func InitLuaApi() {
	Lua.RegisterFunction("RpcCall", RpcCall)
	registerDBService(Lua.GetVm())
}
