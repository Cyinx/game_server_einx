package module

import (
	"game_server/db/dbmanager"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/db/mongodb"
	"github.com/Cyinx/einx/lua"
	"github.com/yuin/gopher-lua"
	"time"
)

var Instance = einx.GetModule("mongodb")
var Lua = einx.NewLuaStae()
var db_cfg = mongodb.NewMongoDBInfo("127.0.0.1", 27916, "test", "", "")

func InitDBComponent() {
	var DBInstance = mongodb.NewMongoDBMgr(Instance, db_cfg, 5*time.Second)
	einx.AddModuleComponent(Instance, DBInstance, db_manager.Instance)
}

func InitLuaApi() {
	//Lua.RegisterFunction("InitDBCfg", InitDBCfg)
}

func LoadLuaConfig() {
	t := Lua.GetVm().GetGlobal("db_config").(*lua.LTable)
	cfg := lua_state.ConvertLuaTable(t)
	db_cfg = mongodb.NewMongoDBInfo(cfg["ip"].(string), int(cfg["port"].(float64)), cfg["dbname"].(string), "", "")
}
