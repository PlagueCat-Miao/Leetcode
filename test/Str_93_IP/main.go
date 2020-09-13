package main

import "strconv"

func restoreIpAddresses(s string) []string {
	ans :=make([]string,0)
	input :=make([]string,0)
	l:= len(s)
	if l > 12{
		return ans
	}
	ans = dps(0,s,input)

	return ans
}

func dps(index int, remStr string,nowIpList []string)[]string {
	ans :=make([]string,0)
	l:= len(remStr)
	if l > 12 - index*3{
		return ans
	}

	if index == 4&&l==0{
		ans =append(ans,nowIpList[0]+"."+nowIpList[1]+"."+nowIpList[2]+"."+nowIpList[3])
		return ans
	}

	for i:=0;i<3&&i<l ;i++{
		nextRemStr := remStr[1+i:]
		IpStr := remStr[:1+i]
		nextIpList := nowIpList
		ispass:=judge(IpStr)
		if ispass{
			nextIpList = append(nextIpList,IpStr)
			tempAns:= dps(index+1,nextRemStr,nextIpList)
			ans=append(ans,tempAns...)
		}
	}
	return ans
}
func judge(s string) bool{
	if s[0]=='0' && len(s) >1{
		return false
	}
	intt,err:=strconv.Atoi(s)
	if err != nil{
		return false
	} else if intt>255{
		return false
	}
	return true
}


