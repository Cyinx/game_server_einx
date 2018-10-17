package outserver_mgr

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
	"msg_def"
	"protobuf_gen"
)

type Agent = einx.Agent
type NetLinker = einx.NetLinker
type AgentID = einx.AgentID
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID
type ITcpClientMgr = einx.ITcpClientMgr
type ProtoTypeID = uint32

var logic = einx.GetModule("logic")
var logic_router = logic.(einx.ModuleRouter)

const (
	ServerType_DBServer = 1
)

type OutServerMgr struct {
	link_map         map[AgentID]NetLinker
	client_component ITcpClientMgr
}

var Instance = &OutServerMgr{
	link_map: make(map[AgentID]NetLinker),
}

func (this *OutServerMgr) OnLinkerConneted(id AgentID, agent Agent) {
	this.link_map[id] = agent.(NetLinker)
	slog.LogInfo("outserver", "outserver connect : %v", id)
	linker := agent.(NetLinker)
	switch linker.GetUserType() {
	case ServerType_DBServer:
		this.OnDBServerConnected(id, agent)
	}
}

func (this *OutServerMgr) OnLinkerClosed(id AgentID, agent Agent) {
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
	var msg msg_def.VersionCheck
	msg.Type = int32(pbgen.VersionType_VersionGateServer)
	linker := agent.(NetLinker)
	b, _, _ := msg_def.MarshalMsg(msg)
	linker.WriteMsg(msg_def.VersionCheckMsgID, b)
}

func (this *OutServerMgr) ServeHandler(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalMsg(id, b)
	if msg != nil {
		logic_router.RouterMsg(agent, id, msg)
	}
}

func (this *OutServerMgr) ServeRpc(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalRpc(id, b)
	if msg != nil {
		logic_router.RouterMsg(agent, id, msg)
	}
}
