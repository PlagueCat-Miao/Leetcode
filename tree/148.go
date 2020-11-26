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
	ans:= sortList(l1)

	fmt.Printf("\nans:\n")
	for ans!=nil{
		fmt.Printf("%d,",ans.Val)
		ans = ans.Next
	}

}

//TODO check 归并好好学
func sortList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}else if head.Next == nil{
		return head
	}

	fast :=head.Next
	slow :=head
	for fast!=nil && fast.Next!=nil{
		fast= fast.Next.Next
		slow=  slow.Next
	}
	leftStart :=slow.Next
	slow.Next =nil
	left:=sortList(head)
	right:=sortList(leftStart)
	return mergeTwoLists(left,right)
}


func mergeTwoLists(l1 *ListNode ,l2 *ListNode)*ListNode{
	if l2 == nil{
		return l1
	}else if l1 ==nil {
		return l2
	}

	tempHead:=&ListNode{0,nil}
	cur := tempHead

	for l1!=nil &&l2 !=nil{
		if l1.Val < l2.Val{
			cur.Next = l1
			l1 = l1.Next
		}else{
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1!=nil {
		cur.Next = l1
	}else if l2!=nil {
		cur.Next = l2
	}

	return tempHead.Next

}
