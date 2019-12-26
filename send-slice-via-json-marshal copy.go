package main
import (
  "fmt"
  
)

type User struct {
	Name string `json:name"`
  Age int `json:"age"`
}
      
// slice 를 json 으로 전달할때는 익명 구조체를 선언하지 말고, slice 자체를 바로 보내버리면 된다. 
func main() {
  a := User {"Kim", 24}
  b := User {"Han", 36}
  c := []User{a,b}
  
  bytes,_ = json.Marshal(c)
  fmt.Println(string(bytes))
 
} 
