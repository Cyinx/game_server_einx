package main

import (
	"game_server/client_mgr"
	"game_server/logic"
	"game_server/outserver_mgr"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
)

var module = einx.GetModule("logic")

func StartTcpServer() {
	slog.LogInfo("game_server", "Listen tcp port:2205")
	einx.AddTcpServerMgr(module, ":2203", clientmgr.Instance)
}

func StartClusterClient() {
	einx.StartTcpClientMgr(module, "outserver", outserver_mgr.Instance)
}

func InitServer() {
	logic.InitLogicModule()
}
