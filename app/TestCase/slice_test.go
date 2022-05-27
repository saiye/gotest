package testcase

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	//理解数学的开闭区间 [a,b)， a<=x<b
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	a := s[1:]
	fmt.Println(a)
	b := s[0:2]
	fmt.Println(b)
	c := s[1:2]
	fmt.Println(c)
}

func TestDownMp4(t *testing.T) {
	urlArr := make([]string, 0)
	urlArr = append(urlArr, "https://1252524126.vod2.myqcloud.com/522ff1e0vodcq1252524126/9e3aa711387702299670171610/playlist.f3.m3u8?time=1650773487567")
	//urlArr = append(urlArr, "")
}
