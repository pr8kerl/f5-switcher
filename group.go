package main

import (
        "fmt"
        "strings"
	      "github.com/gin-gonic/gin"
        //      "github.com/kr/pretty"
        "encoding/json"
        "io/ioutil"
        "log"
)


func postGroup(c *gin.Context) {
	var json SMS
	if c.BindJSON(&json) == nil {
		if json.Mobile == "" || json.Message == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": "invalid message format"})
		} else {
			msgs <- json
			c.JSON(http.StatusOK, gin.H{"status": 200, "message": "message received"})
		}
	}
}

func showGroup(c *gin.Context) {


}

