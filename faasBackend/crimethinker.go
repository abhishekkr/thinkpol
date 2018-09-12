package faasBackend

type CrimeThinker interface {
	NewFunction(jprocDef []byte) (string, error)
	FunctionStatus(procId string) error
	KillFunction(procId string) error
}

/*
FaasEngines acts as map for all available FaaS Backends.
*/
var FaasEngines = make(map[string]CrimeThinker)

/*
faasBackend to return CrimeThinker of correct backend type.
*/
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
