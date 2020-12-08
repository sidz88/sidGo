package main

import "fmt"

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	fmt.Println(rectangle{9.5, 18.5, "blue"})
}
