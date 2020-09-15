// +build ignore

package main

import "fmt"
import "sort"

func main(){
	//s:=[][]int{{1,3},{2,6},{8,10},{15,18}}
	s:=[][]int{{1,4},{0,4}}
	fmt.Printf("%+v",merge(s))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func (i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	ll := len(intervals)
	var i int
	var left,rightMax int
	for i=0;i<ll;{
		left = intervals[i][0]
		rightMax = intervals[i][1]
		j:=i+1
		for;j<ll;j++{
			if rightMax >= intervals[j][0]{
				if rightMax< intervals[j][1]{rightMax =  intervals[j][1]}
				continue
			}

			ans = append(ans, []int{left, rightMax})
			break
		}
		if j==ll{
			ans = append(ans, []int{left, rightMax})
		}
		i=j

	}
	return ans
}