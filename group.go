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

	for _, group := range cfg.Groups {
		log.Printf("processing group %s\n", group.Name)
		var gblue int = 0
		var ggreen int = 0

		// if val, ok := dict["foo"]; ok { //do something here }

		for _, pool := range group.Pools {
			log.Printf("processing pool %s\n", pool.Name)

			err, members := f5.ShowPoolMembers(pool.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			}

			for _, member := range members.Items {
				// check each member - whether blue or green - maybe use a map here?

				if member.Session == "monitor-enabled" {
					for _, bmember := range pool.Blue.Members {
						if bmember == member.FullPath {
							gblue++
							break
						}
					}
					for _, bmember := range pool.Green.Members {
						if bmember == member.FullPath {
							ggreen++
							break
						}
					}

				}

			}

			state := Poolstate{}
			if gblue > 0 {
				state.MemberCount = gblue
				state.Status = "blue"
			} else if ggreen > 0 {
				state.MemberCount = ggreen
				state.Status = "green"
			}
			if (gblue > 0) && (ggreen > 0) {
				state.MemberCount = -1
				state.Status = "orange"
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "pool state": state})
			}
			c.JSON(http.StatusOK, gin.H{"status": 200, "pool state": state})

		}
	}

}

type Poolstate struct {
	MemberCount int
	Status      string
}
