package main

import (
	"fmt"

	"github.com/learning-go-book/package_example/do-format"
	"github.com/learning-go-book/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
