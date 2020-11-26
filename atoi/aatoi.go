package main

import (
	"fmt"
	"strconv"
)

func  main() {
	remain :="10M"
	file :="20M"
	remainInt,_ :=strconv.ParseInt(remain[:2], 10, 64)
	fileInt,_ :=strconv.ParseInt(file[:2], 10, 64)
	fmt.Printf("%v %v %v",remainInt,remainInt > fileInt,fileInt)
}


