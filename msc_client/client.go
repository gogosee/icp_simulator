package main

import (
	"flag"
	"fmt"
)

func main() {
	phone := flag.String("p", "13800001111,13300001111,18600001111", "phone")
	message := flag.String("m", "hello world, 你好世界！", "message")
	iterates := flag.Int("i", 1, "iterates")
	flag.Parse()

	fmt.Printf("phone: %s\n", *phone)
	fmt.Printf("message: %s\n", *message)
	fmt.Printf("iterates: %d\n", *iterates)

	fmt.Println(*phone)
	fmt.Println(*message)
	fmt.Println(*iterates)
}
