package clientmgr

import (
	"msg_def"
)

type Client struct {
	agent Agent
}

func (this *Client) RpcCall(msg interface{}) {
	this.agent.WriteMsg(msg_def.LuaMsgID, msg)
}
