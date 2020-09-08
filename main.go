package main

import (
	"fmt"
	"github.com/aminnairi/golang-template/utils/math"
	"github.com/aminnairi/golang-template/utils/strings"
)

func main() {
	fmt.Println(strings.Reverse("Hello World!")) // !dlroW olleH
	fmt.Println(math.Divide(12, 2))              // 6
}
