package main

import "fmt"

func main() {
	listing()
}

func listing() {
	s := []int{1, 2, 3}
	f(&s)
	fmt.Println(s)
}

func f(s *[]int) {
	*s = append(*s, 10)
}z
