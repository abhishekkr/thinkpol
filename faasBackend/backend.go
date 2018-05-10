package faasBackend

import "github.com/gin-gonic/gin"

type NewSpeak struct {
}

func InitNewSpeak(cacheName string) *NewSpeak {
	return &NewSpeak{}
}

func FaasHelp(ctx *gin.Context) {
	help := map[string]string{
		"name": "thinkpol",
	}

	ctx.JSON(200, help)
}

func FaasPing(ctx *gin.Context) {
	ping := map[string]string{
		"total-proc-count": "-1",
	}

	ctx.JSON(200, ping)
}

func (newspeak *NewSpeak) FunctionStatus(ctx *gin.Context) {
	response := map[string]string{
		"http-method": "get",
	}

	ctx.JSON(200, response)
}

func (newspeak *NewSpeak) NewFunction(ctx *gin.Context) {
	response := map[string]string{
		"http-method": "post",
	}

	ctx.JSON(200, response)
}

func (newspeak *NewSpeak) DeleteFunction(ctx *gin.Context) {
	response := map[string]string{
		"http-method": "del",
	}

	ctx.JSON(200, response)
}
