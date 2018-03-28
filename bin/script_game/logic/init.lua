

local a = {
  ["1"] = 1,
   [1] = 2,
   [3] = {1,2,3,4},
}

DBService.RpcCall("test",a)