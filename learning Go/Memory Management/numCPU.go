package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime.numCPU() = %v\n", runtime.NumCPU())
}
