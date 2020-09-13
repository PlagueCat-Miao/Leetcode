package main

import (
	"fmt"
	"strconv"
)

func main(){

	K:="124"
	int, _ := strconv.Atoi(K)
	string := strconv.Itoa(int)
	fmt.Printf("%v %v",int,string)
}
