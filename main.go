package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pr8kerl/f5-switcher/F5"
	"net/http"
	"os"
	//	"time"
)

var (
	bindaddress = "127.0.0.1:5000"
	f5          *F5.Device
	appRoot     = "/public"
)

func init() {
	// setup config
	err := InitialiseConfig(cfgfile)
	if err != nil {
		fmt.Printf("error reading config: %s\n", err)
		os.Exit(1)
	}
	if cfg.Webconfig.BindAddress != "" {
		if cfg.Webconfig.BindPort > 0 {
			bindaddress = fmt.Sprintf("%s:%d", cfg.Webconfig.BindAddress, cfg.Webconfig.BindPort)
		} else {
			bindaddress = fmt.Sprintf("%s:5000", cfg.Webconfig.BindAddress)
		}
	}
	f5 = F5.New(cfg.F5config.Hostname, cfg.F5config.Username, cfg.F5config.Password)
}

func main() {

	// Creates a router without any middleware by default
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(SetJellyBeans())

	//r.GET("/", index)
	r.StaticFS(appRoot, http.Dir("public/"))

	api := r.Group("/api")
	{
		api.GET("/group", showGroup)
		api.PUT("/group", putGroup)
	}

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
