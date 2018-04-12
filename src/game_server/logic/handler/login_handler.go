package msghandler

import (
	//"game_server/logic/client_mgr"
	"game_server/logic/module"
	"github.com/Cyinx/einx/lua"
	"github.com/Cyinx/einx/slog"
	"msg_def"
)

func InitLoginHandler() {
	RegisterHandler(msg_def.VersionCheckMsgID, CheckVersion)
	RegisterHandler(msg_def.MsgLuaRpcMsgID, MsgLuaRpc)
	RegisterRpcHandler("test", TestRpc)
}

func TestRpc(a interface{}, args []interface{}) {

	slog.LogInfo("testrpc", "%s", args[0])
}

func CheckVersion(agent Agent, args interface{}) {

}

func MsgLuaRpc(agent Agent, args interface{}) {
	msg := args.(*msg_def.LuaRpcMsg)
	lua_runtime := module.Lua
	lua_arg, _ := lua_state.UnMarshal(msg.Payload, lua_runtime.GetVm())
	lua_runtime.PCall("on_message_handler", msg.Rpcname, uint64(agent.GetID()), lua_arg)
}
