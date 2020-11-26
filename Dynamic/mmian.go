// +build ignore

package main

import "fmt"
import "sort"

func main(){
	input:=[][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	ans:= maximalSquare(input)
	fmt.Printf("\nans:%d\n",ans)
}

//TODO check 信封 正反排序后 最大子正序数列 学习一下
func maxEnvelopes(envelopes [][]int) int {
	//正反排序 先正序，再反序
	//正序 保证最大子正序列 结果是信封结果 ，后反序是保证同高（宽）之间不互相套
	sort.Slice(envelopes, func (i,j int)bool{
		if envelopes[i][0] == envelopes[j][0]{
			return   envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	ll:=len(envelopes)
	dp :=make([]int,ll)
	for i:=0;i<ll;i++ {
		dp[i]=1 //至少自己算1个子串
	}
	max:=0
	for i:=0;i<ll;i++{
		for j:=0;j<i;j++{
			dp[i] = max(dp[i],dp[j])
		}
		if max<dp[i]{max = dp[i]}
	}
    return max
}
