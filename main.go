package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	combined := append(nums1, nums2...)
	sort.Ints(combined)
	n := len(combined)
	if n%2 == 0 {
		return (float64(combined[n/2] + combined[(n/2)-1])) / 2
	} else {
		return float64(combined[(n-1)/2])
	}
}
