package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
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
*/

func showGroup(c *gin.Context) {

	log.Printf("processing groups\n")
	f5.PrintResponse(cfg)

	for _, group := range cfg.Groups {
		log.Printf("processing group %s\n", group.Name)
		for _, pool := range group.Pools {
			log.Printf("processing group %s\n", pool.Name)
			err, members := f5.ShowPoolMembers(pool.Name)
			if err != nil {
				log.Printf("pool members %s\n", members)
				c.JSON(http.StatusOK, gin.H{"status": 200, "members": members})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			}
		}
	}

}
