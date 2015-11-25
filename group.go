package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func putGroup(c *gin.Context) {

	var json GroupPutData
	if c.BindJSON(&json) == nil {
		if json.State == "" || json.Name == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid message format"})
		}
	}

	// start transaction
	err, tid := f5.StartTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	group := cfg.Groups[json.Name]
	log.Printf("debug putGroup processing group %s, setting to %s\n", group, json.State)

	for pkey, pool := range group.Pools {
		log.Printf("processing pool %s\n", pkey)

		for _, member := range pool.Blue {

			if json.State == "blue" {
				log.Printf("enabling %s blue member: %s\n", pkey, member)
				f5.OnlinePoolMember(pkey, member)
			}
			if json.State == "green" {
				log.Printf("disabling %s blue member: %s\n", pkey, member)
				f5.OfflinePoolMember(pkey, member)
			}
		}

		for _, member := range pool.Green {

			if json.State == "green" {
				log.Printf("enabling %s green member: %s\n", pkey, member)
				f5.OnlinePoolMember(pkey, member)
			}
			if json.State == "blue" {
				log.Printf("disabling %s green member: %s\n", pkey, member)
				f5.OfflinePoolMember(pkey, member)
			}
		}
	}
	err = f5.CommitTransaction(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	showGroup(c)

}

func showGroup(c *gin.Context) {

	log.Printf("processing groups\n")

	for gkey, group := range cfg.Groups {

		log.Printf("processing group %s\n", gkey)
		var gblue int = 0
		var ggreen int = 0

		// if val, ok := dict["foo"]; ok { //do something here }

		for pkey, pool := range group.Pools {

			log.Printf("processing pool %s\n", pkey)
			var pblue int = 0
			var pgreen int = 0

			err, resp, members := f5.ShowPoolMembers(pkey)
			if err != nil {
				group.Pools[pkey] = group.Pools[pkey].SetError(resp.Message)
				//				c.JSON(http.StatusInternalServerError, gin.H{"error": err, "message": resp.Message})
				continue
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
				group.Pools[pkey] = group.Pools[pkey].SetState("blue")
			} else if ggreen > 0 {
				group.Pools[pkey] = group.Pools[pkey].SetState("green")
			}
			if (pblue > 0) && (pgreen > 0) {
				group.Pools[pkey] = group.Pools[pkey].SetState("orange")
				//c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "pool state": state})
			}
			//			c.JSON(http.StatusOK, gin.H{"status": 200, "pool state": state})

		} // end range group.Pools
		if gblue > 0 {
			cfg.Groups[gkey] = cfg.Groups[gkey].SetState("blue")
		} else if ggreen > 0 {
			cfg.Groups[gkey] = cfg.Groups[gkey].SetState("green")
		}
		if (gblue > 0) && (ggreen > 0) {
			cfg.Groups[gkey] = cfg.Groups[gkey].SetState("orange")
		}

	} // end range cfg.Groups
	c.JSON(http.StatusOK, gin.H{"data": cfg.Groups, "user": currentUser})

}

/*
type Poolstate struct {
	MemberCount int
	Status      string
}
*/
