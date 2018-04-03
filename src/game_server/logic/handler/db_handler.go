package msghandler

import (
	"game_server/logic/module"
	"github.com/Cyinx/einx/lua"
	"github.com/Cyinx/einx/slog"
	"github.com/yuin/gopher-lua"
)

func InitDBHandler() {
	RegisterRpcHandler("mongodb_query_back", OnMongodbQueryBack)
	RegisterRpcHandler("mongodb_insert_back", OnMongodbInsertBack)
}

func OnMongodbQueryBack(sender interface{}, args []interface{}) {
	cb := args[0].(string)
	cb_args := args[1]
	is_success := args[2].(bool)
	result := args[3].(map[string]interface{})

	f := module.GetDBServiceCallBack(cb)
	if f != nil {
		vm := module.Lua.GetVm()
		vm.Push(f)
		vm.Push(lua.LBool(is_success))
		vm.Push(cb_args.(lua.LValue))
		vm.Push(lua_state.ConvertMap(vm, result))
		if err := vm.PCall(3, -1, nil); err != nil {
			slog.LogError("lua", "db callback lua err:%v", err)
		}
	}
}

func OnMongodbInsertBack(sender interface{}, args []interface{}) {
	cb := args[0].(string)
	cb_args := args[1]
	is_success := args[2].(bool)

	f := module.GetDBServiceCallBack(cb)
	if f != nil {
		vm := module.Lua.GetVm()
		vm.Push(f)
		vm.Push(lua.LBool(is_success))
		vm.Push(cb_args.(lua.LValue))
		if err := vm.PCall(3, -1, nil); err != nil {
			slog.LogError("lua", "db callback lua err:%v", err)
		}
	}
}
