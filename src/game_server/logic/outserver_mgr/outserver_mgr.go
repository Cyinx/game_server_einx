package outserver_mgr

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"msg_def"
	"protobuf_gen"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID
type ITcpClientMgr = einx.ITcpClientMgr

const (
	ServerType_DBServer = 1
)

type OutServerMgr struct {
	link_map         map[AgentID]Agent
	client_component ITcpClientMgr
}

var Instance = &OutServerMgr{
	link_map: make(map[AgentID]Agent),
}

func (this *OutServerMgr) OnAgentEnter(id AgentID, agent Agent) {
	this.link_map[id] = agent
	slog.LogInfo("outserver", "outserver connect : %v", id)
	switch agent.GetUserType() {
	case ServerType_DBServer:
		this.OnDBServerConnected(id, agent)
	}
}

func (this *OutServerMgr) OnAgentExit(id AgentID, agent Agent) {
	delete(this.link_map, id)
}

func (this *OutServerMgr) OnComponentError(c Component, err error) {

}

func (this *OutServerMgr) OnComponentCreate(id ComponentID, component Component) {
	component.Start()
	this.client_component = component.(ITcpClientMgr)
	this.client_component.Connect("127.0.0.1:2206", ServerType_DBServer)
}

func (this *OutServerMgr) OnDBServerConnected(id AgentID, agent Agent) {
	var b msg_def.VersionCheck
	b.Type = int32(pbgen.VersionType_VersionGateServer)
	agent.WriteMsg(msg_def.VersionCheckMsgID, &b)
}
