package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(BinarySearchElement())
}

func BinarySearchElement() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(100) + 1
	fmt.Println(n)
	//随机数按时间戳生成
	begin := 1
	end := 100
	var middle int
	//[1,100] 左闭右闭,begin可以等于end
	for begin <= end {
		middle = (begin + end) / 2
		if n > middle {
			begin = middle - 1 //begin可以等于end
		} else if n < middle {
			end = middle + 1 //begin可以等于end
		} else {
			return middle
		}
	}
	panic("not recover")
}

/*1 100 n=99时
  51 76 89 96 100（闭区间允许相等，除完加、减1）
*/
