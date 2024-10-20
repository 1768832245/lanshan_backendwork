package main

import "fmt"

func main() {
	PrimeNumber(97)
}

func PrimeNumber(a int) {
	if a <= 1 {
		fmt.Println("Not PrimeNumber")
	} else if a == 2 {
		fmt.Println("is PrimeNumber")
	}
	for i := 2; i < a; i++ {
		if a%i == 0 {
			fmt.Println("Not PrimeNumber")
			break
		} else if a%i != 0 && i == a-1 {
			fmt.Println("is PrimeNumber")

		}
	}
}

//似乎有一种更简单的数学方法能看是不是素数，但我忘了QAQ
