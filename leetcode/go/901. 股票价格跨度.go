package main

import "fmt"

func main() {
	stockSpanner := Constructor()
	// res := stockSpanner.Next(100)
	// fmt.Println(res)
	res := stockSpanner.Next(31)
	fmt.Println(res)
	res = stockSpanner.Next(41)
	fmt.Println(res)
	res = stockSpanner.Next(48)
	fmt.Println(res)
	res = stockSpanner.Next(59)
	fmt.Println(res)
	res = stockSpanner.Next(79)
	fmt.Println(res)
	// res = stockSpanner.Next(85)
	// fmt.Println(res)
}

type StockSpanner struct {
	stack []Stock
}

func Constructor() StockSpanner {
	stack := make([]Stock, 0)
	spanner := StockSpanner{stack: stack}
	return spanner
}

func (this *StockSpanner) Next(price int) int {
	if len(this.stack) == 0 {
		stock := Stock{price: price, ran: 1}
		this.stack = append(this.stack, stock)
		return 1
	}
	tempStock := Stock{price: price, ran: 1}
	for len(this.stack) > 0 {
		lastStock := this.stack[len(this.stack)-1]
		if tempStock.price >= lastStock.price {
			this.stack = this.stack[:len(this.stack)-1]
			tempStock.ran += lastStock.ran
			if len(this.stack) == 0 {
				this.stack = append(this.stack, tempStock)
				return tempStock.ran
			}
		} else {
			this.stack = append(this.stack, tempStock)
			break
		}
	}
	return tempStock.ran
}

type Stock struct {
	price int
	ran   int
}
