// +build ignore

package main

import "fmt"

func main(){
	s:=[][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}}
	fmt.Printf("%+v",maxAreaOfIsland(s))
}



func maxAreaOfIsland(grid [][]int) int {
	max:=0
  for i,list :=range grid{
  	  for j,_ := range list{
  	  	if grid[i][j]!=1{continue}
		  num:=count(grid,i,j)
  	  	  if max<num{max=num}
	  }
  }

	return max
}

func count (grid [][]int,startI,startJ int)int {
	landCount:=0
	var BFSList []int
	BFSList = append(BFSList, startI,startJ)
	for i:=0;i<len(BFSList);i=i+2{
		I:=BFSList[i]
		J:=BFSList[i+1]
		if I<0||I>=len(grid){continue}
		if J<0||J>=len(grid[0]){continue}
		if grid[I][J] != 1{continue}
		grid[I][J] = 0
		landCount++
		BFSList = append(BFSList, I+1,J)
		BFSList = append(BFSList, I-1,J)
		BFSList = append(BFSList, I,J+1)
		BFSList = append(BFSList, I,J-1)
	}
    return landCount
}