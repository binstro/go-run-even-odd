package main

import (
	"fmt"

	goevenodd "github.com/binstro/go_even_odd"
	goevenoddV2 "github.com/binstro/go_even_odd/v2"
)

func main() {
	result := goevenodd.CheckEvenOdd(2)
	resultV2 := goevenoddV2.CheckEvenOdd(1, 2, 3, 4, 5)
	fmt.Println("V1 : ", result)
	fmt.Println("V2 : ", resultV2)
}
