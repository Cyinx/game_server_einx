package dbmodule

import (
	"game_server/db/handler"
	"game_server/db/module"
)

func InitDBModule() {
	module.InitLuaApi()
	module.Lua.DoFile("script_db/init.lua")
	module.InitDBComponent()
	msghandler.InitDBHandler()
}
