// +build ignore

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	l1 := &ListNode{1, &ListNode{1, &ListNode{1,nil}}}
	//l2 := &ListNode{1, &ListNode{1, &ListNode{1,nil}}}
	//l1 := &ListNode{2,nil}
	//l2 := &ListNode{1, nil}
	ans:= detectCycle(l1)

	fmt.Printf("\nans:\n")
	for ans!=nil{
		fmt.Printf("%d,",ans.Val)
		ans = ans.Next
	}

}

func detectCycle(head *ListNode) *ListNode {
	var fast* ListNode
	var slow* ListNode
	slow = head
	fast = head
	flag:=false
	for fast!=nil && fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow==fast {
			flag = true
			break
		}
	}
	if flag {
		fast=head
		for fast!=slow {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}


	return nil
}