package main

import (
	"fmt"
	"testing"
)

func Test_bytesDisplay(t *testing.T) {
	fmt.Println([]byte("hello"))
	fmt.Printf("%c\n", 104)
}
