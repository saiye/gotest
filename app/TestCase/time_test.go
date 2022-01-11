package testcase

import (
	"fmt"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	//ps 踩坑字符串必须要与layout 格式一致，否则结果是00010101
	timeStr:="2019-04-07"
	layout := "2006-01-02" //转换的时间字符串带秒则 为 2006-01-02 15:04:05
	//now, _:= time.ParseInLocation(layout, timeStr, time.Local)
	now, _:= time.Parse(layout, timeStr)

	fmt.Println("RES1:",now.Format("20060102")) //年月日
}

