package msghandler

import (
	"game_server/client_mgr"
	"game_server/logic/module"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/lua"
	"msg_def"
)

var mysql einx.Module = einx.GetModule("mysql")

func InitLoginHandler() {
	RegisterHandler(msg_def.VersionCheckMsgID, CheckVersion2)
	RegisterHandler(msg_def.LuaMsgID, MsgLuaRpc)
	RegisterRpcHandler("test", TestRpc)
}

func TestRpc(ctx Context, args []interface{}) {

}

func CheckVersion(ctx Context, args interface{}) {
	agent := ctx.GetSender()
	client := clientmgr.GetClient(uint64(agent.GetID()))
	msg := args.(*msg_def.VersionCheck)
	m := make(map[int]float64)
	m[-1] = 123
	for i := 0; i < 1024; i++ {
		//m[i] = make(map[int]interface{})
		m[i] = float64(i*123) + m[i-1]
	}
	mysql.RpcCall("test", m)
	client.SendMsg(msg_def.VersionCheckMsgID, msg)
}

func CheckVersion2(ctx Context, args interface{}) {
	agent := ctx.GetSender()
	client := clientmgr.GetClient(uint64(agent.GetID()))
	msg := args.(*msg_def.VersionCheck)
	m := make(map[string]interface{})
	m["user_name"] = "name"
	m["id"] = 132
	m["pwd"] = "123sdfsd"
	mysql.RpcCall("test", m)
	client.SendMsg(msg_def.VersionCheckMsgID, msg)
	//	agent.WriteMsg(msg_def.VersionCheckMsgID, msg)
}

var lua_on_message_handler lua_state.LValue = nil

func MsgLuaRpc(ctx Context, args interface{}) {
	agent := ctx.GetSender()
	msg := args.(*lua_state.LTable)
	lua_runtime := module.Lua
	if lua_on_message_handler == nil {
		lua_on_message_handler = lua_runtime.GetGlobal("on_message_handler")
	}
	lua_runtime.PCall3(lua_on_message_handler,
		msg.RawGetInt(1), lua_state.LNumber(agent.GetID()), msg.RawGetInt(2))
}
