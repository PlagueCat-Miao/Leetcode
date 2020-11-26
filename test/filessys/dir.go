
package main

import (
"fmt"
"os"
)
//
//// 判断文件夹是否存在
//func PathExists(path string) (bool, error) {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true, nil
//	}
//	if os.IsNotExist(err) {
//		return false, nil
//	}
//	return false, err
//}
//
//func main() {
//	_dir := "./gzFiles2"
//	exist, err := PathExists(_dir)
//	if err != nil {
//		fmt.Printf("get dir error![%v]\n", err)
//		return
//	}
//
//	if exist {
//		fmt.Printf("has dir![%v]\n", _dir)
//	} else {
//		fmt.Printf("no dir![%v]\n", _dir)
//		// 创建文件夹
//		err := os.Mkdir(_dir, os.ModePerm)
//		if err != nil {
//			fmt.Printf("mkdir failed![%v]\n", err)
//		} else {
//			fmt.Printf("mkdir success!\n")
//		}
//	}
//}

//func main(){
//	mypath , _ := user.Current() ///home/hellcat
//	fmt.Printf("%s", mypath.HomeDir)
//}
//
func main(){
	_, err := os.OpenFile("./ash/as.go", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("[func RtmpPushExec OpenFile]  err= %+v ", err)
	}
	fmt.Printf("[OK]  err= %+v ", err)
}