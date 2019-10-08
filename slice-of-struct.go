// slice-of-struct check if it works
// I tested I just not 100% sure about nesting struct to slice
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

// {GUID:GUID-01 UserServerURL:{UserAdded:a.com/added UserModified:a.com/modifyed UserDeleted:a.com/deleted}}
// {GUID:GUID-02 UserServerURL:{UserAdded:b.com/added UserModified:b.com/modifyed UserDeleted:b.com/deleted}}
// {GUID:GUID-03 UserServerURL:{UserAdded:c.com/added UserModified:c.com/modifyed UserDeleted:c.com/deleted}}
