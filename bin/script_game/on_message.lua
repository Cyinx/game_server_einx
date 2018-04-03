RpcService = RpcService or { rpc_handlers = {}, }

local rpc_handlers = RpcService.rpc_handlers

setmetatable(RpcService, {
  __newindex = function(t,k,v)
  	rpc_handlers[k] = v
  end
})

function on_message_handler(f,id,arg)
	local handler = rpc_handlers[f]
	if handler ~= nil then
		handler(id,arg)
	end
end