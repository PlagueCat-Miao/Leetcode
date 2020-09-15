// +build ignore

package main

import "fmt"

func main(){
	s:=[]int{4,5,6,7,0,1,2}
	fmt.Printf("%+v",search(s,9))
}

var ll int
var tgt int
func search(nums []int, target int) int {
	ll=len(nums)
	tgt =target
	line:=0
	for line=1;line<ll;line++{
		if nums[line-1]>nums[line]{
			break;
		}
	}

  return Dsearch(0,ll,nums,line)
}

func Dsearch(left,right int,nums []int,offset int)int{
	if right<=left{
		return -1
	}
	fmt.Printf("%v,%v\n",left,right)

	mid :=(right+left)/2
	realMid:=(mid+offset)%ll
	midVal:=nums[realMid]

	if midVal>tgt{
		return Dsearch(left,mid,nums,offset)
	}else if  midVal<tgt{
		return Dsearch(mid+1,right,nums,offset)
	}else {
		return realMid
	}

}
