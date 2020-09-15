// +build ignore

package main

import "fmt"

func main(){
	fmt.Printf("%+v",getPermutation(3,2))
}
//https://leetcode-cn.com/problems/permutation-sequence/
var NN [14]int
var NNN int
func getPermutation(n int, k int) string {
	NN[1] =1
	NN[0] =1
	NNN= n
	var reNum []int
	reNum =append(reNum,1)
	for i:=2;i<=n;i++{
		NN[i] = NN[i-1]*i
		reNum =append(reNum,i)
	}

	var ans []int
	return Dget(0,k-1,reNum,ans)

}

func Dget(index int,k int, reNUM[]int,ans[]int)string {

	if index== NNN{
		return print(ans)
	}
	No:=k / NN[NNN-1-index]

	//fmt.Printf("%+v %+v %+v\n",k,NN[NNN-1-index],reNUM)
	ans = append(ans,reNUM[No])
	reNUM = append(reNUM[:No],reNUM[No+1:]...)
	return Dget(index+1,k%NN[NNN-1-index],reNUM,ans)
}

func print(ans []int) string{
	ll := len(ans)
	var  ansStr string
 	for i:=0;i<ll;i++{
		ansStr+=fmt.Sprintf("%d",ans[i])
	}
	return ansStr
}
