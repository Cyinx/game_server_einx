local RpcService = _G.RpcService
local DBService = _G.DBService
local PlayerMgr = _G.PlayerMgr

RpcService.Register = function(agent_id,args)
	local user_name = args[1]
	local user_data = {}
	user_data.user_name = user_name
	user_data.id = 123
	user_data.pwd = 123.01

	DBService.Insert("user",user_data,"OnUserRegister",{abc = "123"})
end