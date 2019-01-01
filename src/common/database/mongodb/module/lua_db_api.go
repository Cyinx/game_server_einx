package module

import (
	"game_server/db/dbmanager"
	"github.com/Cyinx/einx/lua"
	"github.com/yuin/gopher-lua"
)

const db_lua_table_name = "DBManager"

func registerDBManager(L *lua.LState) {
	mt := L.NewTypeMetatable(db_lua_table_name)
	L.SetGlobal(db_lua_table_name, mt)
	// methods
	L.SetField(mt, "Insert", L.NewFunction(Insert))
	L.SetField(mt, "QueryOne", L.NewFunction(QueryOne))
}

func Insert(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	collection := L.CheckString(1)
	args := L.CheckTable(2)

	q := lua_state.ConvertLuaTable(args)

	if dbmanager.GetInstance().Insert(collection, q) == nil {
		L.Push(lua.LBool(true))
	} else {
		L.Push(lua.LBool(false))
	}

	return 1
}

func QueryOne(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	collection := L.CheckString(1)
	args := L.CheckTable(2)

	q := lua_state.ConvertLuaTable(args)

	result := make(map[string]interface{})

	if dbmanager.GetInstance().DBQueryOneResult(collection, q, result) == nil {
		L.Push(lua.LBool(true))
		L.Push(lua_state.ConvertMap(L, result))
	} else {
		L.Push(lua.LBool(false))
	}

	return 2
}
