package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(offer.GetRandPage(7*24*3600))
	//fmt.Println(2818785 /3600/24)
	//offer.TestOffer05()
	//daily_practice.TestCalculate2()
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano() / 1000 % 10)
	fmt.Println(now.Unix() % 10)
}
