// +build ignore

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
var pp int=5
var qq int=1
func Totree(nums []int)*TreeNode{
	ll:=len(nums)
	if ll ==0{
		return nil
	}
	var LBFList []*TreeNode
	for i:=0;i<ll;i++{
		if nums[i] ==-1{
			LBFList = append(LBFList,nil)
			continue
		}
		subTree:=&TreeNode{nums[i],nil,nil}
		LBFList = append(LBFList,subTree)
	}
	for i:=1;i<=ll;i++{
		//加东西

		if LBFList[i-1] ==nil{
			continue
		}
		if 2*i <=ll{
			LBFList[i-1].Left = LBFList[2*i-1]
		}
		if 2*i +1 <=ll{
			LBFList[i-1].Right = LBFList[2*i]
		}
	}
	return LBFList[0]
}

func main(){

	input := []int{3,9,20,-1,-1,15,7}
	tree := Totree(input)
	ans:= zigzagLevelOrder(tree)
	fmt.Printf("\nans:%+v\n",ans)
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	var LBFList [5000]*TreeNode
	var LeveList [5000]int
	LBFleng :=1
	LBFList[0]=root
	LeveList[0]=0
	var ans [][]int
	if root == nil{
		return ans
	}
	ans =append(ans,[]int{})
	ceng :=1
	for i:=0;i<LBFleng;i++{
		cur:=LBFList[i]
		if cur == nil{
			continue
		}
		if ceng <= LeveList[i]{
			ans =append(ans,[]int{})
			ceng++
		}
		if LeveList[i]%2 ==0 {
			ans[LeveList[i]] = append (ans[LeveList[i]],cur.Val)
		}else{
			ans[LeveList[i]] = append ([]int{cur.Val},ans[LeveList[i]]...)
		}
		LBFList[LBFleng] = cur.Left;LeveList[LBFleng]=LeveList[i]+1; LBFleng++
		LBFList[LBFleng] = cur.Right;LeveList[LBFleng]=LeveList[i]+1; LBFleng++
	}
	return ans
}