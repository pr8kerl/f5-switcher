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
    var blue int = 0
    var green int = 0

// if val, ok := dict["foo"]; ok { //do something here }


		for _, pool := range group.Pools {
			log.Printf("processing pool %s\n", pool.Name)
      
			err, members := f5.ShowPoolMembers(pool.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			}

      for _, member := range members {
        // check each member - whether blue or green - maybe use a map here?
        
        if member.session == "monitor-enabled" {
          for _, bmember range pool.Blue {
            if bmember == member.fullPath {
              blue++
              break
            }
          }
          for _, bmember range pool.Green {
            if bmember == member.fullPath {
              blue++
              break
            }
          }

        }

      }

			if err != nil {
				log.Printf("error %s\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 200, "members": members})
			}
		}
	}

}

struct Poolstate {
  ExpectedCount int
  ActiveCount int
  State string
}

