package main

import (
	"fmt"
)

func go_step(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	return go_step(n-1)+go_step(n-2)+go_step(n-3)
}

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"zy", 32})
	fmt.Println(person{name: "zzg"})
	fmt.Println(person{name: "lj", age: 18})
	fmt.Println(&person{name: "wm", age: 31})
	fmt.Println(go_step(6))
}
