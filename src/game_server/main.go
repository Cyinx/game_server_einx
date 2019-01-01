package main

import (
	//"common/database/mysql"
	"game_server/config"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"runtime"
)

func init() {
	slog.SetLogPath("log/game_server/")
}

func main() {
	config.InitServerConfig()
	slog.LogInfo("game_server", "开始服务器...")
	slog.LogInfo("game_server", "服务器CPU核心数: [%d]", runtime.NumCPU())
	slog.LogInfo("game_server", "正在初始化数据库模块")
	//mysql.InitDBModule()
	slog.LogInfo("game_server", "正在初始化逻辑模块")
	InitServer()
	slog.LogInfo("game_server", "注册消息处理器")
	StartTcpServer()
	StartClusterClient()
	einx.Run()
	einx.Close()
}
