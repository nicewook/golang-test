package main

import (
	"fmt"
)

type User struct {
	name string
	age int
}

// change element value of the struct slice
// https://play.golang.org/p/r1u-cNtMC0J
// range 를 쓰면서, 구조체 슬라이스 내의 값을 바꾸고 싶다. 
func main() {
	a := []User{
		{"A", 10},
		{"B", 12},
		{"C", 15},
	}
	
	// 일단 출력
	fmt.Println(a)
	
	// range 를 받는 변수가 하나 뿐이면 slice 의 index 이다. 
	for i := range a {
		fmt.Println(i)
	}
	
	// 이렇게 index 와 value 를 받을 수도 있다.
	for i, v := range a {
		fmt.Println(i, v)
	}
	
	// slice 의 구조체의 값을 바꾸려면 index 만 받아서 아래와 같이 처리하면 된다. 
	for i := range a {
		if a[i].name == "A" {
			fmt.Println("age is: ", a[i].age)
			a[i].age = 20
		}
	}
	
	fmt.Println(a)
}
