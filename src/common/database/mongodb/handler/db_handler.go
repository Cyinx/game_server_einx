package msghandler

import (
	"game_server/db/dbmanager"
	"game_server/db/module"
	"github.com/Cyinx/einx/lua"
	"github.com/yuin/gopher-lua"
)

func InitDBHandler() {
	RegisterRpcHandler("Insert", Insert)
	RegisterRpcHandler("QueryOne", QueryOne)
	RegisterRpcHandler("db_lua_rpc", OnDBLuaRpc)
}

func OnDBLuaRpc(ctx Context, args []interface{}) {
	rpc_name := args[0].(lua.LValue)
	lua_args := args[1].(*lua.LTable)
	lua_runtime := module.Lua
	lua_runtime.PCall("on_rpc_handler", rpc_name, 0, lua_args)
}

func QueryOne(ctx Context, args []interface{}) {
	collection := args[0].(string)
	content := args[1].(*lua.LTable)
	cb := args[2].(string)
	cb_args := args[3]

	q := lua_state.ConvertLuaTable(content)

	is_success := false
	result := make(map[string]interface{})

	if dbmanager.GetInstance().DBQueryOneResult(collection, q, result) == nil {
		is_success = true
	}

	logic_module.RpcCall("mongodb_query_back", cb, cb_args, is_success, result)
}

func Insert(ctx Context, args []interface{}) {
	collection := args[0].(string)
	cond := args[1].(*lua.LTable)
	cb := args[2].(string)
	cb_args := args[3]

	q := lua_state.ConvertLuaTable(cond)

	is_success := false

	if dbmanager.GetInstance().Insert(collection, q) == nil {
		is_success = true
	}

	logic_module.RpcCall("mongodb_insert_back", cb, cb_args, is_success)
}

func CheckVersion(ctx Context, args interface{}) {

}
