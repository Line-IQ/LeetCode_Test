package main

import (
	"LeetCode/funcs"
	"fmt"
)

func test2() {
	nums := []int{2, 7, 11, 15, 23}
	an1 := funcs.TwoSum(nums, 9)
	fmt.Println(an1)

	l4 := &funcs.ListNode{}
	l3 := l4
	for i := 0; i < 7; i++ {
		l3.Next = &funcs.ListNode{Val: 9}
		l3 = l3.Next
	}
	l1 := l4.Next
	l2 := funcs.ListNode{
		Val: 9,
		Next: &funcs.ListNode{
			Val: 9,
			Next: &funcs.ListNode{
				Val:  9,
				Next: &funcs.ListNode{Val: 9},
			},
		},
	}
	l7 := funcs.AddTwoNumbers(l1, &l2)
	for cur := l7; cur != nil; cur = cur.Next {
		fmt.Printf("%d", cur.Val)
		if cur.Next != nil {
			fmt.Printf(",")
		}
	}
}

func test3() {
	s := "abba"
	sLen := funcs.LengthOfLongestSubstring(s)
	fmt.Println(sLen)
}

func test4() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	num := funcs.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(num)
}

func main() {
	test3()

	test4()
	n := 5
	fmt.Println(n / 2.0)
}
