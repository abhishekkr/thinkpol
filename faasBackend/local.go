package faasBackend

import (
	"fmt"

	"github.com/abhishekkr/gol/golbin"
)

type LocalCrime struct {
}

func init() {
	RegisterFaasEngine("local", new(LocalCrime))
}

func NewFunction(jproc joycampProc.Proc) error {
	golbin.ExecWithEnv("env", map[string]string{"GOL": fmt.Sprintf("%v", jproc)})
}

func FunctionStatus(procId uint64) error {
	golbin.ExecWithEnv("env", map[string]string{"GOL": "status"})
}

func KillFunction(procId uint64) error {
	golbin.ExecWithEnv("env", map[string]string{"GOL": "kill"})
}
