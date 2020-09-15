// +build ignore

package main

import "fmt"

func main(){

	s1:=" "
	fmt.Println(reverseWords(s1)+"e")
}

func reverseWords(s string) string {
	sChars := []rune(s)
	ll :=len(sChars)
	if ll==0{
		return ""
	}
	sChars =reverse(sChars,0,ll)
	sChars =no_(sChars)
	ll=len(sChars)
	left:=0
	for i:=0;i<ll;i++{
		if sChars[i] ==' '{
			sChars = reverse(sChars,left,i)
			//fmt.Printf("%s\n",string(sChars))
			left = i+1
		}
	}
	sChars = reverse(sChars,left,ll)
	return string(sChars)
}

func reverse(s []rune,left,right int)[]rune {
	ansChars :=make([]rune,len(s))
	copy(ansChars,s)
    //TODO check 你又把反向算错了
	for i:=0;i<right-left;i++{
		ansChars[right-1-i] = s[left+i]
	}
	return ansChars
}


func no_(s []rune)[]rune {
	ll := len(s)
	var ansChar []rune
	if  s[0] !=' '{
		ansChar=append(ansChar,s[0])
	}
	for i:=1;i<ll;i++{
		if s[i]==' '&&s[i-1]==' '{
			continue
		}
		ansChar=append(ansChar,s[i])
	}
	ll = len(ansChar)
	if ll>0&&ansChar[ll-1] == ' '{
		return ansChar[:ll-1]
	}
	return ansChar
}