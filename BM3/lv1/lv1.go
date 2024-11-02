package main

import "fmt"

type KQ struct {
	Fahrenheit float64
	Celsius    float64
}

func (KQ *KQ) ToFahrenheit(a float64) {
	KQ.Fahrenheit = 32 + a*1.8
}

func (KQ *KQ) ToCelsius(a float64) {
	KQ.Celsius = (a - 32) / 1.8
}

func main() {
	Kq := new(KQ)
	Kq.ToFahrenheit(30)
	Kq.ToCelsius(30)
	fmt.Println(Kq.Fahrenheit, Kq.Celsius)
}
