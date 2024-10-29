package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var a, b float64
	var d, out string
	for {

		fmt.Println("请输入第一个数")
		fmt.Scanln(&a)
		fmt.Println("请输入第二个数")
		fmt.Scanln(&b)
		fmt.Println("请输入运算符")
		fmt.Scanln(&d)
		result := JustDoIt(d)
		//测试函数
		if result(1, 1) == 0 {
			panic("无效预算符")
		}
		fmt.Println(result(a, b))
		fmt.Println("还想再来一次吗？键入yes/no")
		fmt.Scanln(&out)
		if out == "no" {
			break
		}
	}
}

func JustDoIt(c string) func(a, b float64) float64 {
	switch c {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return func(a, b float64) float64 {
			return 0
		}
	}
}
