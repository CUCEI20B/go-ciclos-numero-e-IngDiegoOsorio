package main

import "fmt"

func main()  {
	var e float64
	var exponencial int
	var limit int
	e = 1
	fmt.Scanln(&limit)
	for i := 1; i <= limit; i++ {
		exponencial = 1
		for j := 1; j <= i; j++ {
			exponencial = exponencial * j	
		}
		e = e + 1/float64(exponencial)
	}
	fmt.Println(e)
}