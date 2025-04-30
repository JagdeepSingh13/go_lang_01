package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [4]bool
	fmt.Println(vals)

	nums1 := [3]int{1, 2, 3}
	fmt.Println(nums1)

	// 2-D arrays
	nums2 := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums2)
}
