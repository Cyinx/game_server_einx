package msghandler

import (
	"game_server/db/dbmanager"
	"github.com/Cyinx/einx/lua"
	"github.com/yuin/gopher-lua"
)

func InitDBHandler() {
	RegisterRpcHandler("Insert", Insert)
	RegisterRpcHandler("QueryOne", QueryOne)
}

func QueryOne(sender interface{}, args []interface{}) {
	collection := args[0].(string)
	content := args[1].(*lua.LTable)
	cb := args[2].(string)
	cb_args := args[3]

	q := lua_state.ConvertLuaTable(content)

	is_success := false
	result := make(map[string]interface{})

	if db_manager.GetInstance().DBQueryOneResult(collection, q, result) == nil {
		is_success = true
	}

	logic_module.RpcCall("mongodb_query_back", cb, cb_args, is_success, result)
}

func Insert(sender interface{}, args []interface{}) {
	collection := args[0].(string)
	cond := args[1].(*lua.LTable)
	cb := args[2].(string)
	cb_args := args[3]

	q := lua_state.ConvertLuaTable(cond)

	is_success := false

	if db_manager.GetInstance().Insert(collection, q) == nil {
		is_success = true
	}

	logic_module.RpcCall("mongodb_insert_back", cb, cb_args, is_success)
}

func CheckVersion(agent Agent, args interface{}) {

}
