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

	for gindex, group := range cfg.Groups {

		log.Printf("processing group %s\n", group.Name)
		var gblue int = 0
		var ggreen int = 0

		// if val, ok := dict["foo"]; ok { //do something here }

		for pindex, pool := range group.Pools {

			log.Printf("processing pool %s\n", pool.Name)
			var pblue int = 0
			var pgreen int = 0

			err, members := f5.ShowPoolMembers(pool.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			}

			for _, member := range members.Items {
				// check each member - whether blue or green - maybe use a map here?

				if member.Session == "monitor-enabled" {
					for _, bmember := range pool.Blue {
						if bmember == member.FullPath {
							gblue++
							pblue++
							break
						}
					}
					for _, bmember := range pool.Green {
						if bmember == member.FullPath {
							ggreen++
							pgreen++
							break
						}
					}

				}

			}

			if pblue > 0 {
				group.Pools[pindex].State = "blue"
			} else if ggreen > 0 {
				group.Pools[pindex].State = "green"
			}
			if (pblue > 0) && (pgreen > 0) {
				group.Pools[pindex].State = "orange"
				//c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "pool state": state})
			}
			//			c.JSON(http.StatusOK, gin.H{"status": 200, "pool state": state})

		} // end range group.Pools
		if gblue > 0 {
			cfg.Groups[gindex].State = "blue"
		} else if ggreen > 0 {
			cfg.Groups[gindex].State = "green"
		}
		if (gblue > 0) && (ggreen > 0) {
			cfg.Groups[gindex].State = "orange"
		}

	} // end range cfg.Groups
	c.JSON(http.StatusOK, gin.H{"data": cfg.Groups})

}

/*
type Poolstate struct {
	MemberCount int
	Status      string
}
*/
