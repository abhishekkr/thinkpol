package faasBackend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/abhishekkr/gol/golenv"
	gollog "github.com/abhishekkr/gol/gollog"
	joycampProc "github.com/abhishekkr/joycamp/proc"
	gin "github.com/gin-gonic/gin"
)

var (
	LocalJoycampPath = golenv.OverrideIfEnv("LOCAL_JOYCAMP_PATH", "/tmp/joycamp")
)

type ThoughtCrime struct {
	CrimeThinker CrimeThinker
}

func InitThoughtCrime(cacheName string) *ThoughtCrime {
	return &ThoughtCrime{}
}

func FaasHelp(ctx *gin.Context) {
	help := map[string]string{
		"name": "thinkpol",
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(200, help)
}

func FaasPing(ctx *gin.Context) {
	ping := map[string]string{
		"total-proc-count": "-1",
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(200, ping)
}

func (thoughtcrime *ThoughtCrime) FunctionStatus(ctx *gin.Context) {
	faasBackend(ctx.Param("backend")).FunctionStatus(ctx.Param("procId"))
	response := map[string]string{
		"http-method": "get",
	}

	backend := ctx.Param("backend")

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(200, response)
}

func (thoughtcrime *ThoughtCrime) NewFunction(ctx *gin.Context) {
	joycampCfg, err = ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		gollog.Err(err)
	}

	var jproc joycampProc.Proc
	err = json.Unmarshal(joycampCfg, &jproc)
	if err != nil {
		gollog.Err(err)
	}
	faasBackend(ctx.Param("backend")).NewFunction(jproc)

	response := map[string]string{
		"http-method": "post",
		"cfg":         joycampCfg,
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(200, response)
}

func (thoughtcrime *ThoughtCrime) KillFunction(ctx *gin.Context) {
	faasBackend(ctx.Param("backend")).KillFunction(ctx.Param("procId"))
	response := map[string]string{
		"http-method": "del",
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(200, response)
}
