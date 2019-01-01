package module

import (
	"game_server/client_mgr"
	"github.com/Cyinx/einx/slog"
	"github.com/yuin/gopher-lua"
)

var buffer_table *lua.LTable = nil

func RpcCall(L *lua.LState) int {
	if L.GetTop() < 3 {
		slog.LogInfo("lua", "%v", L.GetTop())
		return 1
	}
	sid := L.CheckNumber(1)
	f := L.CheckAny(2)
	t := L.CheckAny(3)

	s := clientmgr.GetClient(uint64(sid))
	if s != nil {
		if buffer_table == nil {
			buffer_table = L.NewTable()
		}

		buffer_table.RawSetInt(1, f)
		buffer_table.RawSetInt(2, t)
		s.RpcCall(buffer_table)
	} else {
		slog.LogInfo("lua", "not found client %v", sid)
	}

	return 0
}
