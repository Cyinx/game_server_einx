package dbmodule

import (
	"game_server/db/handler"
	"game_server/db/module"
)

func InitDBModule() {
	module.InitLuaApi()
	module.Lua.DoFile("script_game/db/init.lua")
	module.LoadLuaConfig()
	module.InitDBComponent()
	msghandler.InitLoginHandler()
}
