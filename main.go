package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pr8kerl/f5-switcher/F5"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	bindaddress          string = "127.0.0.1:5000"
	f5                   *F5.Device
	appRoot              string = "/public"
	currentUser          string = "luser"
	currentUserFirstName string = "luser"
	debug                bool   = false
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
	r.Use(MyLogger(gin.DefaultWriter))
	r.Use(gin.Recovery())
	r.Use(SetJellyBeans())
	r.Use(GetUser)

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
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Powered-By", "Black Jelly Beans")
		c.Next()
	}
}

func GetUser(c *gin.Context) {
	luser := c.Request.Header.Get("X-Remote-User")
	if len(luser) > 0 {
		currentUser = luser
		fmt.Printf("current user: %s\n", currentUser)
	}
	c.Next()
}

func MyLogger(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		fmt.Fprintf(out, "[GIN] %v | %3d | %13v | %s | %s | %-7s %s\n%s",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			currentUser,
			method,
			path,
			comment,
		)
	}
}
