package main

import "fmt"

func main(){
	A :="asdasd"
	B:= (A[2] =='d')
	c :=A[2:3]




	fmt.Printf("%v\n",B)
	fmt.Printf("%s\n", c )

	//TODO check  检查char的两种形式
	CC:= make( map[uint8]int)
	d :=A[2]
	CC[d]=23
	fmt.Printf("%+v\n", CC )

	DD := make(map[rune]int)
  	for _,char :=range A{
		DD[char] +=1
  	}
  	fmt.Printf("%+v\n",DD['d'])

  	EE :=make([]byte,2)
	EE[0] =A[2]
	FF:=make([]uint8,2)
	FF[0] =A[2]

}