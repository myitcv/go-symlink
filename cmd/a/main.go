package main

import (
	"fmt"

	"github.com/myitcv/go-symlink/cmd/a/internal/cmdinternallib1"
	"github.com/myitcv/go-symlink/mylib1"
)

func main() {
	fmt.Println(mylib1.DoSomething(), cmdinternallib1.DoSomething())
}
