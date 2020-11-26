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
	ans:= getIntersectionNode(l1)

	fmt.Printf("\nans:\n")
	for ans!=nil{
		fmt.Printf("%d,",ans.Val)
		ans = ans.Next
	}

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	Aman:= headA
	Bman:= headB
	AHp := 1
	BHp := 1
	if headA == nil{
		return headA
	}
	for Aman !=nil && Bman !=nil{
		if Aman == Bman{
			return Aman
		}
		Aman= Aman.Next
		Bman= Bman.Next
		if Aman ==nil&&AHp>0{
			Aman = headB
			AHp--
		}
		if Bman ==nil&&BHp>0{
			Bman = headA
			BHp--
		}
	}
	if Aman == Bman{
		return Aman
	}
	return nil

}