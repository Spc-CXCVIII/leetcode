package leetcode

import "sort"

// m = จำนวนตัวเลขที่ไม่ใช่ 0 ใน array nums1
// n = จำนวนตัวเลขที่ไม่ใช่ 0 ใน array nums2
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	return nums1
}
