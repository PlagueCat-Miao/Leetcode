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
		if nums[i]==pp{P=subTree}
		if nums[i]==qq{Q=subTree}
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

	input := []int{3,5,1,6,2,0,8,-1,-1,7,4}
	tree := Totree(input)
	ans:= lowestCommonAncestor(tree,P,Q)
	fmt.Printf("\nans:%+v\n",ans)
	//var ans2 [][]int
	//ans2 =append(ans2,[]int{})
	//fmt.Printf("\nans:%+v\n",ans2)
}

var P *TreeNode
var Q *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	P=p
	Q=q
	//fmt.Printf("%+v %+v %+v\n",P,Q,root)
	_,_,ans :=dfs(root)
	return ans
}
func dfs(now *TreeNode)(bool,bool,*TreeNode){

	pOK:=false
	qOK:=false
	if now ==nil{
		return false,false,nil
	}else if now == P {
		pOK = true
	}else if now == Q {
		qOK = true
	}



	LpOk,LqOk,Lans:=dfs(now.Left)
	if Lans !=nil{
		return  true,true,Lans
	}
	RpOk,RqOk,Rans:=dfs(now.Right)
	if Rans !=nil{
		return  true,true,Rans
	}
	pOK = pOK || LpOk || RpOk
	qOK = qOK || LqOk ||RqOk
	if pOK == qOK&&pOK == true{
		return true,true,now
	}
	return pOK,qOK,nil

}
