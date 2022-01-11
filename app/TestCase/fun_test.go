package testcase

import (
	"fmt"
	"testing"
)


func TestDefer(t *testing.T) {
/*	var v int32 =10
	fmt.Print("----START-----")
	r := defer1(&v)
	fmt.Print(r,v)
	fmt.Print("----END-----")*/

	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int){ println(i) } (i)
	}
}

func defer1(v *int32) int32 {
	defer func() {
		*v++
		fmt.Print("----RUN DEFER-----")

	}()
	fmt.Print("----defer-MAIN-----")
	return *v + 10
}
