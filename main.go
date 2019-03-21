package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("Hello World!")
    fmt.Println("GOARCH:", runtime.GOARCH)
}
