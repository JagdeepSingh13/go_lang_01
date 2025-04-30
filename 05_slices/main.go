package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
func main() {
	// uninitialized slices -> nil
	// var nums []int
	var nums = make([]int, 2, 5)
	nums2 := []int{}

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)

	// after appending the ele. in nums
	var nums3 = make([]int, len(nums))

	nums2 = append(nums2, 1)

	fmt.Println(nums)
	fmt.Println(cap(nums))

	fmt.Println(nums2)
	fmt.Println(cap(nums2))

	// copy-fn.
	copy(nums3, nums)

	fmt.Println(nums3)
	fmt.Println(slices.Equal(nums3, nums))

	// slice-operator
	var num = []int{1, 2, 3}

	fmt.Println(num[0:2])
	fmt.Println(num[:2])
	fmt.Println(num[1:])
}
