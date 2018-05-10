package main

import (
	"flag"

	golenv "github.com/abhishekkr/gol/golenv"
	gollog "github.com/abhishekkr/gol/gollog"
	gin "github.com/gin-gonic/gin"

	faasBackend "github.com/abhishekkr/thinkpol/faasBackend"
)

var (
	/*
		HTTPAt specifies server's listen-at config, can be overridden by env var DORY_HTTP. Defaults to '':8080'.
	*/
	HTTPAt   = golenv.OverrideIfEnv("DORY_HTTP", ":8080")
	FaasType = golenv.OverrideIfEnv("FAAS_TYPE", "local")
)

func main() {
	flag.Parse()

	gollog.Debug("starting thinkpol as server")
	ginUp(HTTPAt)
	gollog.Debug("bye .")
}

/*
ginCors to set required HTTP configs.
*/
func ginCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}

/*
ginHandleErrors to manage issues at server side.
*/
func ginHandleErrors(ctx *gin.Context) {
	ctx.Next()
	errorToPrint := ctx.Errors.ByType(gin.ErrorTypePublic).Last()
	if errorToPrint != nil {
		ctx.JSON(500, gin.H{
			"status":  500,
			"message": errorToPrint.Error(),
		})
	}
}

func ginUpFaas(router *gin.Engine, newspeak *faasBackend.NewSpeak) {
	router.GET("/faas", newspeak.FunctionStatus)
	router.POST("/faas", newspeak.NewFunction)
	router.DELETE("/faas", newspeak.DeleteFunction)
}

/*
ginUp maps all routing logic and starts server.
*/
func ginUp(listenAt string) {
	newspeak := faasBackend.InitNewSpeak(FaasType)

	router := gin.Default()
	router.Use(ginCors())
	router.Use(ginHandleErrors)
	router.Use(gollog.GinLogrus(), gin.Recovery())

	router.GET("/help", faasBackend.FaasHelp)
	router.GET("/ping", faasBackend.FaasPing)

	ginUpFaas(router, newspeak)

	router.Run(listenAt)
}
