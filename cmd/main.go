package main

import (
	"fmt"
	"hexarch/internal/adapters/core/arithmetic"
	"hexarch/internal/ports"
)

func main() {
	// --- PORTS ---
	var (
		core ports.ArithmeticPort
	)

	// --- ADAPTERS ---
	core = arithmetic.NewAdapter()

	test, _ := core.Multiply(20, 21)
	fmt.Println(test)
}
