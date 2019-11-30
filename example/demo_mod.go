package main

import "fmt"

func main() {
	message := hello()
	fmt.Println(message)
}

func hello() string {
	message := "你好, 北京"
	return message
}
