package mysql

import (
	"github.com/Cyinx/einx"
)

type WorkerPool = einx.WorkerPool
type Module = einx.Module

var worker_pool WorkerPool = einx.CreateModuleWorkers("mysql", 4)
