package mysql

import (
	"common/database/mysql/dbmanager"
	"common/database/mysql/handler"
	"common/database/mysql/lua"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/db/mysql"
	"time"
)

var db_cfg *mysql.MysqlConnInfo = nil //mysql.NewMysqlConnInfo("127.0.0.1", 3306, "gamedb", "root", "123456")

func InitDBComponent() {
	worker_pool.ForEachModule(func(m Module) {
		var component = mysql.NewMysqlMgr(m, db_cfg, 5*time.Second)
		mgr := &dbmanager.DBManager{}
		einx.AddModuleComponent(m, component, mgr)
	})
}

func SetDBCfg(ip string, port int, dbname string, user string, pwd string) {
	db_cfg = mysql.NewMysqlConnInfo(
		ip,
		port,
		dbname,
		user,
		pwd)
}

func InitDBModule() {
	luaApi.InitLuaApi()
	InitDBComponent()
	msghandler.InitDBHandler()
}
