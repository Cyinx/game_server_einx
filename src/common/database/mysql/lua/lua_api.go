package luaApi

import (
	"github.com/yuin/gopher-lua"
)

func PCall(f string, args ...interface{}) {
	lua_runtime.PCall(f, args...)
}

func RpcCall(L *lua.LState) int {

	return 0
}
