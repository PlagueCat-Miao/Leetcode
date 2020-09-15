// +build ignore

package main

import "fmt"

func main(){
	s:=[]int{1,3,5,4,7}
	fmt.Printf("%+v",findLengthOfLCIS(s))
}

func findLengthOfLCIS(nums []int) int {
	max:=1
	now:=1
	for i,_:=range nums{
		if i==0{
			continue
		}
		//fmt.Printf("%v  %v>=%v%v %v\n",i,nums[i-1],nums[i-1] >= nums[i], nums[i],now )
		if  nums[i-1] >= nums[i]{
			if now>max{max =now}
			now =1
			continue
		}
		now++
	}
	if now>max{max =now}
	return max
}