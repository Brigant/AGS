package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string)

	resultch <- fetchResource(1)
	result := <-resultch
	fmt.Println(result)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Resource %d", n)
}
