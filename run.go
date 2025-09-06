package main

import (
	"LeetCode/funcs"
	"fmt"
	"strconv"
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

func test6() {
	s := "PAYPALISHIRING"
	res := funcs.Convert(s, 1)
	fmt.Println(res)
}

func test7() {
	x := -123
	res := funcs.Reverse(x)
	fmt.Println(res)
}

func test264() {
	n := 4
	num := funcs.ThUglyNumber(n)
	fmt.Println(num)
}

func test8() {
	r := -22313456348
	a := funcs.Reverse(r)
	fmt.Println(a)
}

func test9() {
	s := "42"
	res := funcs.MyAtoi(s)
	fmt.Println(res)
}

func main() {
	test3()

	test4()
	n := 5
	fmt.Println(n / 2.0)
	a := make([][]int, 5, 8)
	for i := 0; i < len(a); i++ {

	}
	fmt.Println()

	str := [][]string{{"a", "b", "c"}, {"d", "e", "f"}}
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			fmt.Print(str[i][j], ",")
		}
	}

	mode := 1
	switch mode {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
	s := "abcde"
	fmt.Println(string(s[0]))

	//test6()

	b := -32
	fmt.Println(strconv.Itoa(b)[0:2])

	test7()

	bb := make([]int, 0, 5)
	bb = append(bb, 2, 3)
	fmt.Println(bb)
	fmt.Println(len(bb))

	test264()

	test8()

	//c := "+"
	//fmt.Println(strconv.Atoi(c))
	test9()
}
