// +build ignore

package main

import "fmt"


func main(){

	s1:="adc"
	s2:="dcda"
	fmt.Println(checkInclusion(s1,s2))
}
func checkInclusion(s1 string, s2 string) bool {
	s1hash:=map[rune]int{}
	s1ll:=len(s1)
	s2ll:=len(s2)
	if s2ll<s1ll{
		return false
	}

	for _,char:=range s1{
		s1hash[char]+=1
	}
	s2hash:=map[rune]int{}
	for i:=0;i+s1ll<=s2ll;i++{
		if i==0{
			for _,char:=range s2[i:i+s1ll]{
				s2hash[char]+=1
			}
		}else{
			s2hash[rune(s2[i-1])] -=1
			s2hash[rune(s2[i+s1ll-1])] +=1
		}

		//fmt.Printf("%v:%+v %+v\n",i,s1hash,s2hash)
		ok:=eq(s1hash,s2hash)
		if ok{
			fmt.Println(s2[i:i+s1ll])
			return true
		}
	}
	return false

}

func eq(amap,bmap map[rune]int)bool{
	//TODO check  你忘了 'a' = 97  '0' 是48 ‘A’
	for i:='a';i<'a'+26;i++{
		aval,aok:=amap[i]
		bval,bok:=bmap[i]
		if !aok{
			aval =0
		}
		if !bok{
			bval =0
		}
		if aval != bval{
			return false
		}
	}
	return true
}