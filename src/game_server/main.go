package main

import (
	"game_server/db"
	"game_server/logic"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"runtime"
)

func main() {
	slog.SetLogPath("log/game_server/")
	slog.LogInfo("game_server", "开始服务器...")
	slog.LogInfo("game_server", "服务器CPU核心数: [%d]", runtime.NumCPU())
	slog.LogInfo("game_server", "正在初始化数据库模块")
	dbmodule.InitDBModule()
	slog.LogInfo("game_server", "正在初始化逻辑模块")
	logic.InitLogicModule()
	slog.LogInfo("game_server", "注册消息处理器")
	slog.LogInfo("game_server", "Listen tcp port:2205")
	logic.StartTcpServer(":2205")
	einx.Run()
	einx.Close()
}
