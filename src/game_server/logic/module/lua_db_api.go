package module

import (
	"github.com/Cyinx/einx/slog"
	"github.com/yuin/gopher-lua"
)

var cb_handlers map[string]*lua.LFunction = make(map[string]*lua.LFunction)

func GetDBServiceCallBack(f string) *lua.LFunction {
	if cb_handler, ok := cb_handlers[f]; ok == true {
		return cb_handler
	}
	return nil
}

const lua_table_name = "DBService"

func registerDBService(L *lua.LState) {
	mt := L.NewTypeMetatable(lua_table_name)
	L.SetGlobal(lua_table_name, mt)
	// methods
	L.SetField(mt, "__newindex", L.NewFunction(newIndexMethod))
	L.SetField(mt, "Insert", L.NewFunction(Insert))
	L.SetField(mt, "QueryOne", L.NewFunction(QueryOne))
	L.SetField(mt, "RpcCall", L.NewFunction(DbRpcCall))
}

func DbRpcCall(L *lua.LState) int {

	if L.GetTop() < 2 {
		return 0
	}
	rpc_name := L.CheckAny(1)
	args := L.NewTable()
	for i := 2; i <= L.GetTop(); i++ {
		args.RawSetInt(i-1, L.CheckAny(i))
	}

	mongodb.RpcCall("db_lua_rpc", rpc_name, args)
	return 1
}

func newIndexMethod(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}
	f := L.CheckString(1)
	h := L.CheckFunction(2)

	cb_handlers[f] = h
	return 0
}

func Insert(L *lua.LState) int {
	slog.LogDebug("db", "insert %v", L.GetTop())
	if L.GetTop() < 4 {
		return 0
	}

	collection := L.CheckString(1)
	content := L.CheckTable(2)
	cb := L.CheckString(3)
	args := L.CheckTable(4)

	mongodb.RpcCall("Insert", collection, content, cb, args)
	return 0
}

func QueryOne(L *lua.LState) int {
	if L.GetTop() < 4 {
		return 0
	}

	collection := L.CheckString(1)
	content := L.CheckTable(2)
	cb := L.CheckString(3)
	args := L.CheckTable(4)

	mongodb.RpcCall("QueryOne", collection, content, cb, args)
	return 0
}
