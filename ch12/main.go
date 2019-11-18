package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Println()

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	// '%T' uses 'reflect.TypeOf' internally
	fmt.Printf("%T\n", w)
	fmt.Println()

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	t = v.Type()
	fmt.Println(t.String())
	fmt.Println()

	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x)
}
