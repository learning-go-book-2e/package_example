package main

import (
	"fmt"

	format "github.com/learning-go-book-2e/package_example/do-format"
	"github.com/learning-go-book-2e/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
