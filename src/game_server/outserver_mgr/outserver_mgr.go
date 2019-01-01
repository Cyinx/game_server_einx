package outserver_mgr

import (
	"github.com/Cyinx/einx"
	//"github.com/Cyinx/einx/slog"
	"msg_def"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type NetLinker = einx.NetLinker
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID
type ITcpClientMgr = einx.ITcpClientMgr
type ModuleRouter = einx.ModuleRouter
type ProtoTypeID = uint32
type Context = einx.Context

var logic = einx.GetModule("logic")
var logic_router = logic.(ModuleRouter)

const (
	Type_AuthServer = 1
)

type OutServerMgr struct {
	link_map         map[AgentID]Agent
	auth_server      *AuthServer
	client_component ITcpClientMgr
}

var Instance = &OutServerMgr{
	link_map: make(map[AgentID]Agent),
}

func (this *OutServerMgr) OnLinkerConneted(id AgentID, agent Agent) {
	linker := agent.(NetLinker)
	switch linker.GetUserType() {
	case Type_AuthServer:
		this.OnAuthServerConnected(id, agent)
	default:
		this.OnClusterServerConnected(id, agent)
	}
}

func (this *OutServerMgr) OnLinkerClosed(id AgentID, agent Agent) {
	delete(this.link_map, id)
}

func (this *OutServerMgr) OnComponentError(ctx Context, err error) {

}

func (this *OutServerMgr) OnComponentCreate(ctx Context, id ComponentID) {
	component := ctx.GetComponent()
	component.Start()
	this.client_component = component.(ITcpClientMgr)
	this.client_component.Connect("127.0.0.1:2206", Type_AuthServer)
}

func (this *OutServerMgr) OnAuthServerConnected(id AgentID, agent Agent) {
	this.auth_server = &AuthServer{
		linker: agent.(NetLinker),
	}

	msg := &msg_def.VersionCheck{}
	msg.Type = 1
	this.auth_server.SendMsg(msg_def.VersionCheckMsgID, msg)
}

func (this *OutServerMgr) OnClusterServerConnected(id AgentID, agent Agent) {
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
