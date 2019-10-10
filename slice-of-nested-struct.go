// https://play.golang.org/p/zrDMlt_IuBY
// check if there is no problem SLICE of the NESTED STRUCT 
// result - no problem
// 1. declare struct "UserServerStruct"
// 2. nest it to another struct "UserServer"
// 3. make slice []UserServer

package main

import (
	"fmt"
)

type UserServerStruct struct {
	UserAdded    string
	UserModified string
	UserDeleted  string
}

type UserServer struct {
	GUID          string
	UserServerURL UserServerStruct
}

var UserServerSubscribeList []UserServer


func main() {
	uss := UserServerStruct{
		UserAdded:    "a.com/added",
		UserModified: "a.com/modifyed",
		UserDeleted:  "a.com/deleted",
	}

	us1 := UserServer{
		"GUID-01",
		uss,
	}

	uss.UserAdded = "b.com/added"
	uss.UserModified = "b.com/modifyed"
	uss.UserDeleted = "b.com/deleted"

	us2 := UserServer{
		"GUID-02",
		uss,
	}

	uss.UserAdded = "c.com/added"
	uss.UserModified = "c.com/modifyed"
	uss.UserDeleted = "c.com/deleted"

	us3 := UserServer{
		"GUID-03",
		uss,
	}

	UserServerSubscribeList = append(UserServerSubscribeList, us1, us2, us3)

	for _, server := range UserServerSubscribeList {
		fmt.Printf("%+v\n", server)
	}
}
