// +build ignore

package main

import "fmt"

func main(){
	s:=[]string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if 0==len(strs){
		return ""
	}
	minll:=9999
	minIndex:=-1
	for i,str:=range strs{
		ll:=len(str)
		if ll<minll{
			minll = ll
			minIndex = i
		}
	}
	//TODO check 二分模板
	left:=0
	right:=minll
	for left <right{
		mid := (right - left +1)/2 +left
		if judge(left,mid,minIndex,strs){
			left = mid
		}else{
			right = mid-1
		}

	}
	return strs[minIndex][:left]
}

func judge(left,right int,minIndex int,strs []string) bool{
	target:=strs[minIndex][left:right]
	for _,str:=range strs{
		if target != str[left:right]{
			return false
		}
	}
	return true
}
