package faasBackend

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	golbin "github.com/abhishekkr/gol/golbin"
	golenv "github.com/abhishekkr/gol/golenv"
	gollog "github.com/abhishekkr/gol/gollog"
)

var (
	CmdDir = golenv.OverrideIfEnv("JOYCAMP_CMD_DIR", "/tmp/joycamp")
)

type LocalCrime struct {
	procId string
}

func init() {
	os.MkdirAll(CmdDir, 0755)

	RegisterFaasEngine("local", new(LocalCrime))
}

func (l *LocalCrime) NewFunction(jprocDef []byte) (string, error) {
	procId := "changeit"
	joycampCfg := path.Join(CmdDir, procId)
	err := ioutil.WriteFile(joycampCfg, jprocDef, 0644)
	if err != nil {
		return "", err
	}

	cmd := fmt.Sprintf("joycamp -cfg %s", joycampCfg)
	stdout, err := golbin.Exec(cmd)
	gollog.Info(stdout)
	return procId, err
}

func (l *LocalCrime) FunctionStatus(procId string) error {
	_, err := golbin.Exec("env")
	return err
}

func (l *LocalCrime) KillFunction(procId string) error {
	_, err := golbin.Exec("env")
	return err
}
