package config

import (
	"common/cfg_reader"
	"common/database/mysql"
	//"github.com/Cyinx/einx/slog"
)

var cfg *cfg_reader.Config = cfg_reader.NewConfig()

func InitServerConfig() {
	cfg.ReadConfig("config")
	mysql.SetDBCfg(
		cfg.GetNamedString("mysql", "ipaddr"),
		cfg.GetNamedInt("mysql", "port"),
		cfg.GetNamedString("mysql", "dbname"),
		cfg.GetNamedString("mysql", "user"),
		cfg.GetNamedString("mysql", "pwd"))
}
