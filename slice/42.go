// +build ignore

package main

import "fmt"

func main(){
	s:=[]int{0}
	fmt.Printf("%+v",trap(s))
}


func trap(height []int) int {

	ll :=len(height)

	left:=0
	right:=ll
	for i:=0;i<ll;i++{
		left =i
		if height[i]!=0{
			break
		}
	}
	for i:=ll-1;i>=0;i--{
		right =i
		if height[i]!=0{
			break
		}
	}
	wa:= 0
	var tWa int
	for left<right{
		//fmt.Printf("%d %d %d\n",left,right,wa)
		if height[left] < height[right]{
			left,tWa=moveLeft(height,left,right)
			wa+=tWa
		}else {
			right,tWa=moveRight(height,left,right)
			wa+=tWa
		}
	}
	return wa
}

func moveLeft(height []int,left,right int) (int,int) {
	waAns :=0
	var i int
	for i=left+1;i<right;i++{
		if height[left] <=height[i]{
			break
		}
		waAns += height[left] - height[i]
	}
	return i,waAns
}

func moveRight(height []int,left,right int) (int,int){
	waAns :=0
	var i int
	for i=right-1;left<i;i--{
		if height[right] <=height[i]{
			break
		}
		waAns += height[right] - height[i]
	}
	return i,waAns
}