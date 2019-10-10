// Just checked if Golang can unmarshal with superset - check omitempty is working or not
// It works!
// 1. make Some struct instance
// 2. marshal it
// 3. unmarshal with Total instance
// 4. print Total instance
// 5. assign to GUID and print Total Instance

package main

import (
	"encoding/json"
	"fmt"
)

type Total struct {
	GUID         string `json:"guid"`
	UserAdded    string `json:"userAdded"`
	UserModified string `json:"userModified"`
	UserDeleted  string `json:"userDeleted"`
}

type TotalEmpty struct {
	GUID         string `json:"guid,omitempty"`
	UserAdded    string `json:"userAdded"`
	UserModified string `json:"userModified"`
	UserDeleted  string `json:"userDeleted"`
}

type Some struct {
	UserAdded    string `json:"userAdded"`
	UserModified string `json:"userModified"`
	UserDeleted  string `json:"userDeleted"`
}

// https://play.golang.org/p/63uEV5g2iWx
func main() {
	b, _ := json.Marshal(Some{
		"http://add.com",
		"http://modify.com",
		"http://delete.com",
	})

	var total Total

	json.Unmarshal(b, &total)
	fmt.Printf("+%v\n", total) // {GUID: UserAdded:http://add.com UserModified:http://modify.com UserDeleted:http://delete.com}

	total.GUID = "myGUID-1234"
	fmt.Printf("%+v\n", total) // {GUID:myGUID-1234 UserAdded:http://add.com UserModified:http://modify.com UserDeleted:http://delete.com}

	// if use omitempty. GUID field will not visible
	// Case 1. no omitempty
	totalNormal := Total{
		UserAdded:    "added",
		UserModified: "modified",
		UserDeleted:  "deleted",
	}

	totalOmitEmpty := TotalEmpty{
		UserAdded:    "added",
		UserModified: "modified",
		UserDeleted:  "deleted",
	}

	b, _ = json.MarshalIndent(totalNormal, "", "  ")
	fmt.Println("Normal: ", string(b))

	// Normal:  {
	// 	"guid": "",
	// 	"userAdded": "added",
	// 	"userModified": "modified",
	// 	"userDeleted": "deleted"
	// }

	b, _ = json.MarshalIndent(totalOmitEmpty, "", "  ")
	fmt.Println("OmitEmpty: ", string(b))

	// OmitEmpty:  {
	// 	"userAdded": "added",
	// 	"userModified": "modified",
	// 	"userDeleted": "deleted"
	// }
}
