package F5

import (
        "fmt"
        "strings"
        "encoding/json"
        "io/ioutil"
        "log"
)

// a pool member
type LBPoolMember struct {
        Name            string `json:"name"`
        Partition       string `json:"partition"`
        FullPath        string `json:"fullPath"`
        Address         string `json:"address"`
        ConnectionLimit int    `json:"connectionLimit"`
        DynamicRatio    int    `json:"dynamicRatio"`
        Ephemeral       string `json:"ephemeral"`
        InheritProfile  string `json:"inheritProfile"`
        Logging         string `json:"logging"`
        Monitor         string `json:"monitor"`
        PriorityGroup   int    `json:"priorityGroup"`
        RateLimit       string `json:"rateLimit"`
        Ratio           int    `json:"ratio"`
        Session         string `json:"session"`
        State           string `json:"state"`
}

// a pool member reference - just a link and an array of pool members
type LBPoolMemberRef struct {
        Link  string         `json:"link"`
        Items []LBPoolMember `json":items"`
}

type LBPoolMembers struct {
        Link  string         `json:"selfLink"`
        Items []LBPoolMember `json":items"`
}

// used by online/offline
type MemberState struct {
        State   string `json:"state"`
        Session string `json:"session"`
}

type LBPool struct {
        Name                   string          `json:"name"`
        FullPath               string          `json:"fullPath"`
        Generation             int             `json:"generation"`
        AllowNat               string          `json:"allowNat"`
        AllowSnat              string          `json:"allowSnat"`
        IgnorePersistedWeight  string          `json:"ignorePersistedWeight"`
        IpTosToClient          string          `json:"ipTosToClient"`
        IpTosToServer          string          `json:"ipTosToServer"`
        LinkQosToClient        string          `json:"linkQosToClient"`
        LinkQosToServer        string          `json:"linkQosToServer"`
        LoadBalancingMode      string          `json:"loadBalancingMode"`
        MinActiveMembers       int             `json:"minActiveMembers"`
        MinUpMembers           int             `json:"minUpMembers"`
        MinUpMembersAction     string          `json:"minUpMembersAction"`
        MinUpMembersChecking   string          `json:"minUpMembersChecking"`
        Monitor                string          `json:"monitor"`
        QueueDepthLimit        int             `json:"queueDepthLimit"`
        QueueOnConnectionLimit string          `json:"queueOnConnectionLimit"`
        QueueTimeLimit         int             `json:"queueTimeLimit"`
        ReselectTries          int             `json:"reselectTries"`
        ServiceDownAction      string          `json:"serviceDownAction"`
        SlowRampTime           int             `json:"slowRampTime"`
        MemberRef              LBPoolMemberRef `json:"membersReference"`
}

type LBPools struct {
        Items []LBPool `json:"items"`
}

func (f *F5) ShowPool(pname string) (error, *LBPool) {

        //u := "https://" + f5Host + "/mgmt/tm/ltm/pool/~" + partition + "~" + pname + "?expandSubcollections=true"
        pool := strings.Replace(pname, "/", "~", -1)
        u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "?expandSubcollections=true"
        res := LBPool{}

        err, resp := f.SendRequest(u, GET, &sessn, nil, &res)
        if err != nil {
                log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
                return err, nil
        } else {
          return nil, &res
        }
}

func (f *F5) ShowPoolMembers(pname string) (error,*LBPoolMembers) {

        pool := strings.Replace(pname, "/", "~", -1)
        //      member := strings.Replace(pmember, "/", "~", -1)
        //u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members?expandSubcollections=true"
        u := "https://" + f5Host + "/mgmt/tm/ltm/pool/" + pool + "/members"
        res := LBPoolMembers{}

        err, resp := f.SendRequest(u, GET, &sessn, nil, &res)
        if err != nil {
                log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
                return err, nil
        } else {
          f.PrintResponse(&res.Items)
          return nil, &res
        }

}

