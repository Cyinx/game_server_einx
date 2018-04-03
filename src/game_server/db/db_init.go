package dbmodule

import (
	"game_server/db/handler"
	"game_server/db/module"
)

func InitDBModule() {
	module.InitLuaApi()
	module.InitDBComponent()
	msghandler.InitDBHandler()
}
