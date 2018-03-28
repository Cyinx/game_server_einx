package module

import (
	"github.com/yuin/gopher-lua"
)

type DBService struct {
}

const lua_table_name = "DBService"

var func_map = map[string]lua.LGFunction{
	"RpcCall": DBRpcCall,
}

func registerPersonType(L *lua.LState) {
	mt := L.NewTypeMetatable(lua_table_name)
	L.SetGlobal(lua_table_name, mt)
	// methods
	L.SetField(mt, "RpcCall", L.NewFunction(DBRpcCall))
	//L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), func_map))
}

func DBRpcCall(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}
	rpc_name := L.CheckString(1)
	t := L.CheckTable(2)
	b := make([]byte, 0, 128)
	b = Lua.Marshal(b, t)
	db_module.RpcCall(rpc_name, b)
	return 1
}
