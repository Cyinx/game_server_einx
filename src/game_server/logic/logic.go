package logic

import (
	"game_server/logic/client_mgr"
	"game_server/logic/handler"
	"game_server/logic/module"
	"game_server/logic/outserver_mgr"
)

func InitLogicModule() {
	module.InitLuaApi()
	module.Lua.DoFile("script_game/init.lua")
	msghandler.InitLoginHandler()
	msghandler.InitDBHandler()
}

func StartTcpServer(port string) {
	module.AddTcpServer(port, clientmgr.Instance)
}

func StartClusterClient() {
	module.StartTcpClient("outserver", outserver_mgr.Instance)
}
