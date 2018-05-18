package faasBackend

import (
	joycampProc "github.com/abhishekkr/joycamp/proc"
)

type CrimeThinker interface {
	NewFunction(jproc joycampProc.Proc) error
	FunctionStatus(procId uint64) error
	KillFunction(procId uint64) error
}

/*
FaasEngines acts as map for all available FaaS Backends.
*/
var FaasEngines = make(map[string]CrimeThinker)

/**/
func faasBackend(backend string) CrimeThinker {
	if backend == "" {
		backend = "local"
	}
	return FaasEngines[backend]
}

/*
RegisterFaasEngine gets used by adapters to register themselves.
*/
func RegisterFaasEngine(name string, faasEngine CrimeThinker) {
	FaasEngines[name] = faasEngine
}

/*
GetFaasEngine gets used by client to fetch a required FaaS-engine.
*/
func GetFaasEngine(name string) CrimeThinker {
	return FaasEngines[name]
}
