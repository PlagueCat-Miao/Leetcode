// +build ignore

package main

import "fmt"

func main(){
	s:="abba"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	ll:=len(s)
	if ll ==0{
		return ll
	}
	max:=1
	maxStr := s[0:1]
	left:=0
	right:=1
	hashmap:=make(map[uint8]int)
	hashmap[s[0]] = 0
	for i:=1;i<ll;i++{
		nowStr:=s[i]
		val,ok:=hashmap[nowStr]
		if ok{
			if val+1>left{ //注意 map中记录了滞后的字符串
				left = val+1
			}
			hashmap[nowStr] = i
		}else{
			hashmap[nowStr] = i
		}
		// for j:=left;j<i;j++{
		//     if nowStr == s[j]{ //有重复
		//         left = j+1
		//     }
		// }
		right = i+1
		if max<right-left{
			max = right-left
			maxStr = s[left:right]
		}
	}
	fmt.Printf("%s",maxStr)
	return max
}
