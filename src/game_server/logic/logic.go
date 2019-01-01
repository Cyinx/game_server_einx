package logic

import (
	"game_server/logic/handler"
	"game_server/logic/module"
	"github.com/Cyinx/einx/lua"
	"msg_def"
)

func GetLuaRuntime() *lua_state.LuaRuntime {
	return module.Lua
}

func InitLogicModule() {
	module.InitLuaApi()
	module.Lua.DoFile("script_game/init.lua")
	msghandler.InitLoginHandler()
	msghandler.InitDBHandler()
	msg_def.SetLuaVm(GetLuaRuntime())
}
