package main

import (
	"fmt"
)

func intSeq() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}

func outfunc(outer string, s int) func(i int) int {
	return func(in int) int{
		s ++
		return s+in
	}
}

func main(){
	nextInt := intSeq() 
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("hello world")
	nextInt = intSeq() 
	fmt.Println(nextInt())
	fmt.Println("hello world")	
	outf := outfunc("china", 0)
	fmt.Println(outf(0))
	fmt.Println(outf(0))
	fmt.Println(outf(0))

}
