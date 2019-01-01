package msghandler

import (
	//"auth_server/db/dbmanager"
	"common/database/mysql/lua"
	"github.com/yuin/gopher-lua"
)

func InitDBHandler() {
	RegisterRpcHandler("db_lua_rpc", OnDBLuaRpc)
}

func OnDBLuaRpc(ctx Context, args []interface{}) {
	rpc_name := args[0].(lua.LValue)
	lua_args := args[1].(*lua.LTable)
	luaApi.PCall("on_rpc_handler", rpc_name, 0, lua_args)
}
