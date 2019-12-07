// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Test_impl(t *testing.T) {
	x := IntSet{}
	x.Add(0)
	x.Add(1)
	x.Add(2)
	x.Add(3)
	fmt.Println(x)
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		name   string
		caller *IntSet
		want   int
	}{
		{
			name:   "one element",
			caller: setFromSlice([]int{1}),
			want:   1,
		},
		{
			name: "three elements",
			caller: setFromSlice([]int{1, 2, 3}),
			want: 3,
		},
		{
			name: "seven elements",
			caller: setFromSlice([]int{1, 2, 3, 5, 8, 13, 21}),
			want: 7,
		},
		{
			name: "fifteen elements",
			caller: setFromSlice([]int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 601, 987}),
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.caller.Len(); got != tt.want {
				t.Errorf("IntSet.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setFromSlice(slice []int) *IntSet {
	x := &IntSet{}
	for _, e := range slice { x.Add(e) }
	return x
}
