package module

import (
	"game_server/db/dbmanager"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/db/mongodb"
	//"github.com/Cyinx/einx/lua"
	//"github.com/yuin/gopher-lua"
	"time"
)

var Instance = einx.GetModule("mongodb")
var Logic = einx.GetModule("logic")
var Lua = einx.NewLuaStae()
var db_cfg = mongodb.NewMongoDBInfo("192.168.1.88", 27916, "test", "", "")

func InitDBComponent() {
	var DBInstance = mongodb.NewMongoDBMgr(Instance, db_cfg, 5*time.Second)
	einx.AddModuleComponent(Instance, DBInstance, dbmanager.Instance)
}

func InitLuaApi() {
	//Lua.RegisterFunction("InitDBCfg", InitDBCfg)
	registerLogicService(Lua.GetVm())
	registerDBManager(Lua.GetVm())
}

func LoadLuaConfig() {

}
