// +build ignore

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	l1 := &ListNode{1, &ListNode{2, &ListNode{4,nil}}}
	l2 := &ListNode{1, &ListNode{4, &ListNode{5,nil}}}
	l3 := &ListNode{1, &ListNode{8, &ListNode{10,nil}}}
	l4 := &ListNode{-9, &ListNode{2, &ListNode{44,nil}}}
	var input []*ListNode
	input = append(input,l1,l2,l3,l4)
	//l2 := &ListNode{1, &ListNode{1, &ListNode{1,nil}}}
	//l1 := &ListNode{2,nil}
	//l2 := &ListNode{1, nil}
	ans:= mergeKLists(input)

	fmt.Printf("\nans:\n")
	for ans!=nil{
		fmt.Printf("%d,",ans.Val)
		ans = ans.Next
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	ll:=len(lists)
	if ll==0{
		return nil
	}else if ll==1{
		return lists[0]
	}else if ll==2{
		return mergetwo(lists[0],lists[1])
	}
	mid:=ll/2

	left :=mergeKLists(lists[0:mid])
	right :=mergeKLists(lists[mid:])
	return mergetwo(left,right)
}


func mergetwo(l1,l2 *ListNode)*ListNode{
	if l1 == nil{
		return l2
	}else if l2==nil {
		return l1
	}
	tHead:=&ListNode{0,nil}
	cur :=tHead
	for l1!=nil&& l2!=nil{
		if l1.Val<l2.Val{
			cur.Next = l1
			l1= l1.Next
		}else{
			cur.Next = l2
			l2= l2.Next
		}
		cur = cur.Next
	}
	if l1!=nil{
		cur.Next = l1
	}else if l2!=nil {
		cur.Next = l2
	}

	return tHead.Next
}