package main

import "fmt"

type Commodity struct {
	Name  string
	Price float64
	Stock int
}

type Electronic struct {
	Commodity
	Brand string
	Model string
}

func main() {
	Bread := Commodity{
		"bread",
		10,
		5,
	}
	Phone := Electronic{
		Commodity: Commodity{
			Name:  "phone",
			Price: 2000,
			Stock: 10,
		},
		Brand: "Apple",
		Model: "16ProMax冷锋蓝",
	}
	Bread.NewPrice(20)
	Bread.NewStock(15)
	Phone.NewPrice(9999)
	Phone.ShowStock()
	Phone.Print()
}

type Do interface {
	NewName(N string)
	NewPrice(P float64)
	NewStock(S int)
	ShowStock()
	Sell()
	Buy()
}

type DoForElectronic interface {
	Do
}

func (c *Commodity) NewName(N string) {
	c.Name = N
	fmt.Println("名字已改为", c.Name)
}

func (c *Commodity) NewPrice(P float64) {
	c.Price = P
	fmt.Println("价格已改为", c.Price)
}

func (c *Commodity) NewStock(S int) {
	c.Stock = S
	fmt.Println("库存已改为", c.Stock)
}

func (c *Commodity) ShowStock() {
	fmt.Println("The Commodity Stock is", c.Stock)
}

func (c *Commodity) Sell() {
	c.Stock--
	fmt.Println("Now,The Commodity Stock is", c.Stock)
}

func (c *Commodity) Buy() {
	c.Stock++
	fmt.Println("Now,The Commodity Stock is", c.Stock)
}

func (e *Electronic) Print() {
	fmt.Println(e.Brand, e.Model)
}
