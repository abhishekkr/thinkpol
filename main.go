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

/*
ginUpFaas to manage FaaS routes
*/
func ginUpFaas(router *gin.Engine, thoughtcrime *faasBackend.ThoughtCrime) {
	faasAPI := router.Group("/faas")
	{
		faasAPI.GET("/:backend/:procId", thoughtcrime.FunctionStatus)
		faasAPI.POST("/:backend", thoughtcrime.NewFunction)
		faasAPI.DELETE("/:backend/:procId", thoughtcrime.KillFunction)
	}
}

/*
ginUp maps all routing logic and starts server.
*/
func ginUp(listenAt string) {
	thoughtcrime := faasBackend.InitThoughtCrime(FaasType)

	router := gin.Default()
	router.Use(ginCors())
	router.Use(ginHandleErrors)
	router.Use(gollog.GinLogrus(), gin.Recovery())

	router.GET("/help", faasBackend.FaasHelp)
	router.GET("/ping", faasBackend.FaasPing)

	ginUpFaas(router, thoughtcrime)

	router.Run(listenAt)
}
