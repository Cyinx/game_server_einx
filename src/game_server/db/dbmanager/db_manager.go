package db_manager

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/db/mongodb"
	"github.com/Cyinx/einx/slog"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID

type DBManager struct {
	db *mongodb.MongoDBMgr
}

var Instance = &DBManager{nil}

func GetInstance() *mongodb.MongoDBMgr {
	return Instance.db
}

func (this *DBManager) OnComponentError(c Component, err error) {
	slog.LogInfo("mongodb", "reconnect to mongodb,error:%v.", err)
	c.Start()
	if err := this.db.Ping(); err != nil {
		slog.LogInfo("mongodb", "mongodb error:%v", err)
	}
}

func (this *DBManager) OnComponentCreate(id ComponentID, component Component) {
	component.Start()
	this.db = component.(*mongodb.MongoDBMgr)
	if err := this.db.Ping(); err != nil {
		slog.LogInfo("mongodb", "mongodb error:%v", err)
	}
	slog.LogInfo("mongodb", "mongodb start")
}
