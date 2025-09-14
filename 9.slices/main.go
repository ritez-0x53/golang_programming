package main

import (
	"fmt"
	"slices"
)

//* slices -> dynamic array
//* most used construct in go
//* +prebuilt useful methods

func main() {

	//* uninitialised slice is nill by default
	var arr []int
	fmt.Println(arr == nil)
	fmt.Println(len(arr))

	//* initialised slice
	var nums = make([]int, 2, 5)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	
	nums = append(nums, 300)
	nums = append(nums, 100)
	nums[0] = 3
	fmt.Println(nums)

	//* copy function
	var nums1 = make([]int , 0, 4)
	nums1 = append(nums1, 100)
	nums1 = append(nums1, 200)
	fmt.Println(nums1)

	var nums2 = make([]int, len(nums1))
	copy(nums2, nums1)
	fmt.Println(nums1,nums2)


	// slice operator
	var intArrs = []int{1,2,3}
	fmt.Println(intArrs[0:2])

	fmt.Println(slices.Equal(nums1 , nums2))
}
