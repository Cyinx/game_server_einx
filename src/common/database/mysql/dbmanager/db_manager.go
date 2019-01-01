package dbmanager

import (
	"database/sql"
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/db/mysql"
	"github.com/Cyinx/einx/slog"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID
type Context = einx.Context

const (
	DB_MYSQLMGR_CONTEXT_IDX = 10001
)

var prepare_sql []string

type DBManager struct {
	db    *mysql.MysqlMgr
	stmts []*sql.Stmt
}

func (this *DBManager) GetSession() *sql.DB {
	return this.db.GetSession()
}

func (this *DBManager) OnComponentError(ctx Context, err error) {
	slog.LogInfo("mysql", "reconnect to mysql,error:%v.", err)
	c := ctx.GetComponent()
	c.Start()
	if err := this.db.Ping(); err != nil {
		slog.LogInfo("mysql", "mysql error:%v", err)
	}
}

func (this *DBManager) OnComponentCreate(ctx Context, id ComponentID) {
	component := ctx.GetComponent()
	component.Start()
	this.db = component.(*mysql.MysqlMgr)
	if err := this.db.Ping(); err != nil {
		slog.LogInfo("mysql", "mysql error:%v", err)
		return
	}
	slog.LogInfo("mysql", "mysql start")
	this.DoPrepare()
	ctx.Store(DB_MYSQLMGR_CONTEXT_IDX, this)
}

func (this *DBManager) DoPrepare() {
	this.stmts = make([]*sql.Stmt, len(prepare_sql))

	for idx, query := range prepare_sql {
		if stmt, err := this.GetSession().Prepare(query); err == nil {
			this.stmts[idx] = stmt
		}
	}

}

func (this *DBManager) GetPrepareStmt(idx int) *sql.Stmt {
	if idx >= len(this.stmts) {
		return nil
	}
	return this.stmts[idx]
}

func PrepareStatement(idx int, query string) {
	if prepare_sql == nil {
		prepare_sql = make([]string, 128)
	}

	length := len(prepare_sql)
	if idx >= length {
		slice := make([]string, length+64)
		copy(slice, prepare_sql)
		prepare_sql = slice
	}
	prepare_sql[idx] = query
}
