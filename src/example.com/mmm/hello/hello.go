package main

import (
	"example.com/mmm/mymath"
	"fmt"
)

func main() {
	fmt.Println("Hello from golang!")
	x := 5
	y := mymath.Square(x)
	fmt.Printf("%d square = %d\n", x, y)
}
