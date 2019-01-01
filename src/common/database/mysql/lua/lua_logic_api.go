package luaApi

import (
	"github.com/yuin/gopher-lua"
)

const logic_lua_table_name = "LogicService"

func registerLogicService(L *lua.LState) {
	mt := L.NewTypeMetatable(logic_lua_table_name)
	L.SetGlobal(logic_lua_table_name, mt)
	// methods
	L.SetField(mt, "RpcCall", L.NewFunction(LogicRpcCall))
}

func LogicRpcCall(L *lua.LState) int {

	if L.GetTop() < 2 {
		return 0
	}

	rpc_name := L.CheckAny(1)
	args := L.NewTable()
	for i := 2; i <= L.GetTop(); i++ {
		args.RawSetInt(i-1, L.CheckAny(i))
	}

	logic_module.RpcCall("db_lua_rpc", rpc_name, args)
	return 0
}
