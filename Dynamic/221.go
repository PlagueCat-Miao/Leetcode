// +build ignore

package main

import "fmt"

func main(){
    input:=[][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	ans:= maximalSquare(input)
	fmt.Printf("\nans:%d\n",ans)
}

//TODO check 正方性动归 学习一下
func maximalSquare(matrix [][]byte) int {

	l1:=len(matrix)
	if l1 == 0{
		return 0
	}
	l2:=len(matrix[0])
	if l2 == 0{
		return 0
	}
	var ansMatrix [1000][1000]int
	max:=0
	for i:=0;i<l1;i++{
		for j:=0;j<l2;j++{
			if matrix[i][j] == '0'{
				ansMatrix[i+1][j+1] =0
				continue
			}
			//fmt.Printf("%d %d %d", ansMatrix[i][j], ansMatrix[i+1][j],ansMatrix[i][j+1])
			ansMatrix[i+1][j+1] = Min( ansMatrix[i][j], ansMatrix[i+1][j],ansMatrix[i][j+1])+1
			if max < ansMatrix[i+1][j+1]{
				max = ansMatrix[i+1][j+1]
			}
		}
	}
	return  max*max
}
func Min(x, y,z int) int {
	if z <= y && z <=x {
		return z
	}
	if x <= y && x <=z {
		return x
	}
	return y
}