package luaApi

import (
	//	"common/database/mysql/dbmanager"
	//"database/sql"
	//"github.com/Cyinx/einx/lua"
	//"github.com/Cyinx/einx/slog"
	"github.com/yuin/gopher-lua"
)

const db_lua_table_name = "DBManager"

func registerDBManager(L *lua.LState) {
	mt := L.NewTypeMetatable(db_lua_table_name)
	L.SetGlobal(db_lua_table_name, mt)
	// methods
	L.SetField(mt, "Query", L.NewFunction(Query))
}

func Query(L *lua.LState) int {
	/*
		if L.GetTop() < 1 {
			L.Push(lua.LBool(false))
			L.Push(lua.LNil)
			return 2
		}

		query := L.CheckString(1)
		var rows *sql.Rows
		var err error
		if L.GetTop() >= 2 {
			args := make([]interface{}, 0, L.GetTop())
			for i := 2; i <= L.GetTop(); i++ {
				args = append(args, lua_state.ConvertLuaValue(L.CheckAny(i)))
			}
			rows, err = dbmanager.GetInstance().GetSession().Query(query, args...)
		} else {
			rows, err = dbmanager.GetInstance().GetSession().Query(query)
		}
		if err != nil {
			slog.LogWarning("mysql", "query error %v", err)
			L.Push(lua.LBool(false))
			L.Push(lua.LNil)
			rows.Close()
			return 2
		}

		column_types, column_err := rows.ColumnTypes()
		if column_err != nil {
			slog.LogError("mysql", "columns error:%v", column_err)
			L.Push(lua.LBool(false))
			L.Push(lua.LNil)
			rows.Close()
			return 2
		}

		values := make([]interface{}, len(column_types))
		results := L.NewTable()

		for c := true; c || rows.NextResultSet(); c = false {
			values = values[:0]
			for k, c := range column_types {
				switch c.DatabaseTypeName() {
				case "INT", "BIGINT":
					values[k] = new(int64)
				case "DOUBLE", "FLOAT":
					values[k] = new(float64)
				case "VARCHAR":
					values[k] = new(string)
				case "BLOB":
					values[k] = new([]byte)
				default:
					values[k] = new([]byte)
				}
			}

			for i := 1; rows.Next(); i++ {
				if err = rows.Scan(values...); err != nil {
					slog.LogError("mysql", "Scan error:%v", err)
					L.Push(lua.LBool(false))
					L.Push(lua.LNil)
					rows.Close()
					return 2
				}

				result := L.NewTable()

				for k, v := range values {
					key := column_types[k]
					switch s := v.(type) {
					case *int64:
						result.RawSetString(key.Name(), lua.LNumber(*s))
					case *float64:
						result.RawSetString(key.Name(), lua.LNumber(*s))
					case *string:
						result.RawSetString(key.Name(), lua.LString(*s))
					case *[]byte:
						ud := L.NewUserData()
						ud.Value = *s
						result.RawSetString(key.Name(), ud)
					}
				}
				results.RawSetInt(i, result)
			}
		}

		L.Push(lua.LBool(true))
		L.Push(results)
		rows.Close()
	*/
	return 2
}
