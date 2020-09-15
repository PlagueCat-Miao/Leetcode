// +build ignore

package main

import "fmt"

func main(){


	fmt.Println(multiply("123","456"))
}


func multiply(num1 string, num2 string) string {
	if num1=="0" ||num2 == "0"{
		return "0"
	}
	ans:= [300]int{}
	llA :=len(num1)
	llB :=len(num2)
	for i,ANumStr := range num1{
		for j,BNumStr := range num2{

			ten,one:=cheng(ANumStr,BNumStr)
			now := llA-i-1+llB-j-1

			ans[now]+=one
			ans[now+1]+=ans[now]/10
			ans[now]%=10

			ans[now+1]+=ten
			ans[now+2]+=ans[now+1]/10
			ans[now+1]%=10
			for q:=2 ;ans[now+q]>9;q++{
				ans[now+q+1]+=ans[now+q]/10
				ans[now+q]%=10
			}

		}
	}
	ll:=len(ans)
	var ansStr []rune
	flag:=false
	for i:=ll-1;i>=0;i--{
		if ans[i]==0&&flag==false{
			continue
		}
		flag=true;
		ansStr = append(ansStr,rune(ans[i]+'0'))

	}

	return string(ansStr)
}

func cheng(A rune, B rune)(int ,int){
	a := A-'0'
	b := B-'0'
	c:=int(a*b)
	return c/10 , c%10
}