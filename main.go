package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pr8kerl/f5-switcher/F5"
	"net/http"
	//	"os"
	//	"time"
)

var (
	bindaddress = "127.0.0.1:5000"
)

func init() {
	// setup config
	InitialiseConfig(cfgfile)
	if config.BindAddress != "" {
		if config.BindPort > 0 {
			bindaddress = fmt.Sprintf("%s:%d", config.BindAddress, config.BindPort)
		} else {
			bindaddress = fmt.Sprintf("%s:5000", config.BindAddress)
		}
	}
  f5 := F5.New(&config)
}

func main() {

	// Creates a router without any middleware by default
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(SetJellyBeans())

	api := r.Group("/api")
	{
		api.GET("/group", showGroup)
		//	api.POST("/group", postGroup)
	}

  // init napping session
  InitSession()

	r.Run(bindaddress)

}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "hello"})
}

func SetJellyBeans() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Powered-By", "Black Jelly Beans")
		c.Next()
	}
}

