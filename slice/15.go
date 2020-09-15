// +build ignore

package main

import "fmt"
import "sort"

func main(){
	s:=[]int{0,1,1}
	fmt.Printf("%+v",threeSum(s))
}


func threeSum(nums []int) [][]int {
	//TODO M九大排序
	sort.Slice(nums, func(i int, j int) bool { //升序 （小在前）
		return nums[i]< nums[j]
	})
	var ans [][]int
	ll:=len(nums)
	if ll ==0{
		return ans
	}
	i:=0
	for i<ll{
		tempAnss:=towSum(nums[i+1:],-nums[i])
		for _,tAns:=range tempAnss{
			ans = append(ans,[]int{nums[i],tAns[0],tAns[1]})
		}
		k:=nums[i]
		for i<ll&&k==nums[i]{i++}
	}
	return ans
}
func towSum(nums []int,sum int) [][]int{
	var ans [][]int
	ll := len(nums)
	left :=0
	right:=ll-1
	for left<right{
		nowSum:=nums[left]+nums[right]
		if nowSum> sum{
			k:=nums[right]
			for left<right&&k==nums[right]{right--}
		}else if nowSum<sum{
			k:=nums[left]
			for left<right&&k==nums[left]{left++}
		}else{
			ans = append(ans,[]int{nums[left],nums[right]})
			k:=nums[left]
			for left<right&&k==nums[left]{left++}
			k=nums[right]
			for left<right&&k==nums[right]{right--}
		}
	}
    return ans

}