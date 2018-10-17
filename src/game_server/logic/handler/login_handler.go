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
	RegisterHandler(msg_def.LuaMsgID, MsgLuaRpc)
	RegisterRpcHandler("test", TestRpc)
}

func TestRpc(ctx Context, args []interface{}) {

	slog.LogInfo("testrpc", "%s", args[0])
}

func CheckVersion(ctx Context, args interface{}) {

}

func MsgLuaRpc(ctx Context, args interface{}) {
	agent := ctx.GetSender()
	msg := args.(*lua_state.LTable)
	lua_runtime := module.Lua
	lua_runtime.PCall("on_message_handler",
		msg.RawGetInt(1), uint64(agent.GetID()), msg.RawGetInt(2))
}
