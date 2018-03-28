package msghandler

import (
	//"game_server/db/dbmanager"
	"game_server/db/module"
	//"github.com/Cyinx/einx/db/mongodb"
	//"github.com/Cyinx/einx/slog"
	//"msg_def"
)

func InitLoginHandler() {
	RegisterRpcHandler("test", TestRpc)
}

func TestRpc(a interface{}, args []interface{}) {
	//db_manager.GetInstance().Insert("test", mongodb.M{"username": "test", "pass": "111"})
	//slog.LogInfo("testrpc", "%v", args[0].([]byte))
	t, _ := module.Lua.UnMarshal(args[0].([]byte))
	module.Lua.PCall2("test", t)
}

func CheckVersion(agent Agent, args interface{}) {

}
