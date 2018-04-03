PlayerMgr = PlayerMgr or { players_ = {},}

local players = PlayerMgr.players_

function PlayerMgr:GetPlayer(id)
	return players[id]
end

function PlayerMgr:AddPlayer(id)
	local player_data = {}
	players[id] = player_data
	return player_data
end

function PlayerMgr:RemovePlayer(id)
	players[id] = nil
end
