package main

import (
	"reflect"
	"testing"
)

func TestMajorityN3(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{}, []int{}},
		{[]int{2, 2, 2, 2}, []int{2}},
	}
	for _, c := range cases {
		if got := MajorityN3(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("MajorityN3(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
