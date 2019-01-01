package luaApi

import (
	"github.com/Cyinx/einx"
)

var lua_runtime = einx.NewLuaStae()
var logic_module = einx.GetModule("logic")

func InitLuaApi() {
	lua_runtime.DoFile("script_db/init.lua")
	registerLogicService(lua_runtime.GetVm())
	registerDBManager(lua_runtime.GetVm())
}

func LoadLuaConfig() {

}
