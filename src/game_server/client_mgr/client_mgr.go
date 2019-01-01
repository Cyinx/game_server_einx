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
type ModuleRouter = einx.ModuleRouter
type ComponentID = einx.ComponentID
type Context = einx.Context
type ProtoTypeID = uint32

var logic = einx.GetModule("logic")
var logic_router = logic.(ModuleRouter)

type ClientMgr struct {
	client_map map[AgentID]*Client
	tcp_link   Component
}

var Instance = &ClientMgr{
	client_map: make(map[AgentID]*Client),
}

func GetClient(agent_id uint64) *Client {
	client, _ := Instance.client_map[AgentID(agent_id)]
	return client
}

func (this *ClientMgr) OnLinkerConneted(id AgentID, agent Agent) {
	this.client_map[id] = &Client{linker: agent.(NetLinker)}
}

func (this *ClientMgr) OnLinkerClosed(id AgentID, agent Agent) {
	delete(this.client_map, id)
}

func (this *ClientMgr) OnComponentError(ctx Context, err error) {

}

func (this *ClientMgr) OnComponentCreate(ctx Context, id ComponentID) {
	component := ctx.GetComponent()
	this.tcp_link = component
	component.Start()
	slog.LogInfo("gate_client", "Tcp sever start success")
}

func (this *ClientMgr) ServeHandler(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalMsg(id, b)
	if msg != nil {
		logic_router.RouterMsg(agent, id, msg)
	}

}

func (this *ClientMgr) ServeRpc(agent Agent, id ProtoTypeID, b []byte) {
	msg := msg_def.UnmarshalRpc(id, b)
	if msg != nil {
		logic_router.RouterMsg(agent, id, msg)
	}
}
