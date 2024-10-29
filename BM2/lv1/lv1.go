package main

import "fmt"

type b map[int]int

func main() {
	a := []int{1, 1, 3, 4, 5}
	fmt.Println(Abc(a))
}

func Abc(a []int) b {
	result := map[int]int{} //map类型必须分配内存再操作，不然会panic  ((result b)下面必须再对result = make(map[int]int))
	for _, v := range a {
		result[v]++
	}
	return result
}
