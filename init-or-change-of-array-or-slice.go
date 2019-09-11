package main

import (
	"fmt"
)

// initialize or change Array or Slice
// https://play.golang.org/p/20scD3uef_c
// Array 나 Slice 를 초기화 (또는 특정 값으로 일괄 변경) 하려면?
func main() {
	// 문자열 슬라이스를 만든 다음, make 로 슬라이스 길이만큼 len, cap 을 지정해서 make 하니 깨끗하게 "" 로 초기화 됨
	a := []string{"A", "B", "C", "D", "E"}
	a = make([]string, len(a), len(a)) 
	fmt.Println("a: ",a, len(a), cap(a))

	// 정수 슬라이스 역시 같은 방법을 사용하니 default 값인 0 으로 잘 초기화 됨
	b := []int{1,2,3,4,5}
	b = make([]int, len(b), len(b)) 
	fmt.Println("b: ",b, len(b), cap(b))
	
	// copy 명령을 써도 된다. 
	x := []int{1,2,3,4,5}
	y := make([]int, len(x), len(x))
	copy(x,y)
	fmt.Println("x: ",x, len(x), cap(x))
	
		
	// 배열(Array) 는 어떨까? 그냥 대입하면 된다. 
	// 또한 하나를 변경하였다 하여 다른게 변경되지도 않는다. (값의 대입인 것이다)
	c := [5]int{1,2,3,4,5}
	d := [5]int{}
	c = d
	d[0]= 3
	c[0]= 4
	fmt.Println("c: ",c, len(c), cap(c))
	fmt.Println("d: ",d, len(d), cap(d))
}
