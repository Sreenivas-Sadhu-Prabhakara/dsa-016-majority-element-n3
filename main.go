package main

import (
	"fmt"
	"sort"
)

// MajorityN3 returns every value occurring more than len(nums)/3 times —
// at most two can — using extended Boyer–Moore voting with two candidate
// slots, then a verification pass. Result is sorted.
func MajorityN3(nums []int) []int {
	cand1, cand2 := 0, 1 // any two distinct seeds; counts gate their use
	cnt1, cnt2 := 0, 0
	for _, v := range nums {
		switch {
		case cnt1 > 0 && v == cand1:
			cnt1++
		case cnt2 > 0 && v == cand2:
			cnt2++
		case cnt1 == 0:
			cand1, cnt1 = v, 1
		case cnt2 == 0:
			cand2, cnt2 = v, 1
		default:
			cnt1--
			cnt2--
		}
	}
	out := []int{}
	for _, cand := range []int{cand1, cand2} {
		n := 0
		for _, v := range nums {
			if v == cand {
				n++
			}
		}
		if n > len(nums)/3 {
			out = append(out, cand)
		}
	}
	if len(out) == 2 && out[0] == out[1] {
		out = out[:1]
	}
	sort.Ints(out)
	return out
}

func main() {
	fmt.Println(MajorityN3([]int{3, 2, 3}))
	fmt.Println(MajorityN3([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
