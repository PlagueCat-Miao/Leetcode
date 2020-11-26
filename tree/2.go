// +build ignore

package main

import "fmt"

func main(){
	//l1 := &ListNode{1, &ListNode{1, &ListNode{1,nil}}}
	//l2 := &ListNode{1, &ListNode{1, &ListNode{1,nil}}}
	l1 := &ListNode{2,nil}
	l2 := &ListNode{1, nil}
	ans:= addTwoNumbers(l1,l2)

	fmt.Printf("\nans:\n")
	for ans!=nil{
		fmt.Printf("%d,",ans.Val)
		ans = ans.Next
	}

}
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode
	var ll *ListNode
	i:=0
	c:=0
	for l1 != nil && l2 != nil{
		a := l1.Val
		b := l2.Val
		c += a+b
		point :=ListNode{c%10,nil}
		c =c/10
		if i==0{
			head = &point
			cur = head

		}else{
			cur.Next = &point
			cur = cur.Next
		}
		l1 =l1.Next
		l2 =l2.Next
		i++
	}
	ll =nil
	if l1 != nil {
		ll = l1
	}else if l2 != nil{
		ll = l2
	}
	for ll != nil{
		a := ll.Val
		c += a
		point :=ListNode{c%10,nil}
		c =c/10
		if i==0{
			head = &point
			cur = head
		}
		cur.Next = &point
		cur = cur.Next
		ll =ll.Next
		i++
	}
	if c >0{
		point :=ListNode{c%10,nil}
		cur.Next = &point
	}
	return head

}