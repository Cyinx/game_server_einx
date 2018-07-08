package clientmgr

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"msg_def"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type NetLinker = einx.NetLinker
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID
type ProtoTypeID = uint32

var logic = einx.GetModule("logic")

type ClientMgr struct {
	client_map map[AgentID]*Client
	tcp_link   Component
}

var Instance = &ClientMgr{
	client_map: make(map[AgentID]*Client),
}

func (this *ClientMgr) GetClient(agent_id uint64) (*Client, bool) {
	client, ok := this.client_map[AgentID(agent_id)]
	return client, ok
}

func (this *ClientMgr) OnAgentEnter(id AgentID, agent Agent) {
	net_linker := agent.(NetLinker)
	this.client_map[id] = &Client{linker: net_linker}
	if id%1000 == 0 {
		slog.LogWarning("client", "client id [%v]", id)
	}
	var msg msg_def.VersionCheck
	b, _, _ := msg_def.MarshalMsg(msg)
	net_linker.WriteMsg(msg_def.VersionCheckMsgID, b)
}

func (this *ClientMgr) OnAgentExit(id AgentID, agent Agent) {
	delete(this.client_map, id)
}

func (this *ClientMgr) OnComponentError(c Component, err error) {

}

func (this *ClientMgr) OnComponentCreate(id ComponentID, component Component) {
	this.tcp_link = component
	component.Start()
	slog.LogInfo("gate_client", "Tcp sever start success")
}

func (this *ClientMgr) ServeHandler(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalMsg(id, b)
	if msg != nil {
		logic.RouterMsg(agent, id, msg)
	}

}

func (this *ClientMgr) ServeRpc(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalRpc(id, b)
	if msg != nil {
		logic.RouterMsg(agent, id, msg)
	}
}
