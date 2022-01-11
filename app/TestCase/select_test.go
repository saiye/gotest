package testcase

import (
	"fmt"
	"runtime"
	"testing"
)





func TestNoRaceSelect1(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool)
	c := make(chan bool)
	c1 := make(chan bool)
	fmt.Println("----run1---")
	go func() {
		x = 1
		// At least two channels are needed because
		// otherwise the compiler optimizes select out.
		// See comment in runtime/select.go:^func selectgo.
		select {
		case c <- true:
			fmt.Println("----run2---")
		case c1 <- true:
			fmt.Println("----run3---")
		}
		fmt.Println("----run4---")
		compl <- true
		fmt.Println("----run5---")
	}()

	fmt.Println("----run--tt---")
	select {
	case <-c:
		fmt.Println("----run6---")
	case c1 <- true:
		fmt.Println("----run7---")
	}
	x = 2
	fmt.Println("----run8---")
	<-compl
	fmt.Println("----run9---")
}

func TestNoRaceSelect2(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool)
	c := make(chan bool)
	c1 := make(chan bool)
	fmt.Println("----run-start---")
	go func() {
		select {
		case <-c:
			fmt.Println("----run1---")
		case <-c1:
			fmt.Println("----run2---")
		}
		x = 1
		fmt.Println("----run3---")
		compl <- true
		fmt.Println("----run4---")
	}()
	x = 2
	fmt.Println("----run5---")
	close(c)
	fmt.Println("----run6---")
	runtime.Gosched()
	fmt.Println("----run7---")
	<-compl
	fmt.Println("----run8---")
}

func TestNoRaceSelect3(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool)
	c := make(chan bool, 10)
	c1 := make(chan bool)
	go func() {
		x = 1
		select {
		case c <- true:
			fmt.Println("----run1---")
		case <-c1:
			fmt.Println("----run2---")
		}
		fmt.Println("----run-com start---")
		compl <- true
		fmt.Println("----run-com end---")
	}()
	fmt.Println("----run3---")
	<-c
	x = 2
	fmt.Println("----run4---")
	<-compl

	fmt.Println("----run5---")
}

func TestNoRaceSelect4(t *testing.T) {
	type Task struct {
		f    func()
		done chan bool
	}

	queue := make(chan Task)
	dummy := make(chan bool)
	fmt.Println("----run1---")
	go func() {
		for {
			select {
			case t := <-queue:
				t.f()
				t.done <- true
				fmt.Println("----run2---")
			}
		}
	}()

	doit := func(f func()) {
		done := make(chan bool, 1)
		select {
		case queue <- Task{f, done}:
			fmt.Println("----run3---")
		case <-dummy:
			fmt.Println("----run4---")
		}
		select {
		case <-done:
			fmt.Println("----run5---")
		case <-dummy:
			fmt.Println("----run6---")
		}
	}

	var x int
	fmt.Println("----run7---")
	doit(func() {
		x = 1
		fmt.Println("doit--",x)
	})
	fmt.Println("----run8---")
	_ = x
}

func TestNoRaceSelect5(t *testing.T) {
	test := func(sel, needSched bool) {
		var x int
		_ = x
		ch := make(chan bool)
		c1 := make(chan bool)

		done := make(chan bool, 2)
		go func() {
			if needSched {
				runtime.Gosched()
			}
			// println(1)
			x = 1
			if sel {
				select {
				case ch <- true:
				case <-c1:
				}
			} else {
				ch <- true
			}
			done <- true
		}()

		go func() {
			// println(2)
			if sel {
				select {
				case <-ch:
				case <-c1:
				}
			} else {
				<-ch
			}
			x = 1
			done <- true
		}()
		<-done
		<-done
	}

	test(true, true)
	test(true, false)
	test(false, true)
	test(false, false)
}

func TestRaceSelect1(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool, 2)
	c := make(chan bool)
	c1 := make(chan bool)
	fmt.Println("-----start----")
	go func() {
		<-c
		<-c
		fmt.Println("-----go1----")
	}()
	f := func() {
		select {
		case c <- true:
			fmt.Println("-----f1----")
		case c1 <- true:
			fmt.Println("-----f2----")
		}
		x = 1
		fmt.Println("-----comp1----")
		compl <- true
		fmt.Println("-----comp2----")
	}
	fmt.Println("-----end1----")
	go f()
	go f()
	fmt.Println("-----end2----")
	<-compl
	fmt.Println("-----end3----")
	<-compl
	fmt.Println("-----end4----")
}

func TestRaceSelect2(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool)
	c := make(chan bool)
	c1 := make(chan bool)
	go func() {
		x = 1
		select {
		case <-c:
		case <-c1:
		}
		compl <- true
	}()
	close(c)
	x = 2
	<-compl
}

func TestRaceSelect3(t *testing.T) {
	var x int
	_ = x
	compl := make(chan bool)
	c := make(chan bool)
	c1 := make(chan bool)
	go func() {
		x = 1
		select {
		case c <- true:
		case c1 <- true:
		}
		compl <- true
	}()
	x = 2
	select {
	case <-c:
	}
	<-compl
}

func TestRaceSelect4(t *testing.T) {
	done := make(chan bool, 1)
	var x int
	go func() {
		select {
		default:
			x = 2
		}
		done <- true
	}()
	_ = x
	<-done
}

// The idea behind this test:
// there are two variables, access to one
// of them is synchronized, access to the other
// is not.
// Select must (unconditionally) choose the non-synchronized variable
// thus causing exactly one race.
// Currently this test doesn't look like it accomplishes
// this goal.
func TestRaceSelect5(t *testing.T) {
	done := make(chan bool, 1)
	c1 := make(chan bool, 1)
	c2 := make(chan bool)
	var x, y int
	fmt.Print("start\n")
	go func() {
		select {
		case c1 <- true:
			x = 1
			fmt.Print("RUN-C1\n")
		case c2 <- true:
			y = 1
			fmt.Print("RUN-C2\n")
		default:
			fmt.Print("----RUN-default\n")
		}
		done <- true
		fmt.Print("RUN-done\n")
	}()
	fmt.Printf("----X%v--Y%v---\n",x,y)
	_ = x
	_ = y
	//<-done
	fmt.Print("end1")
	if <-done{
		//fmt.Print("end2",	done)
		fmt.Printf("end2----X%v--Y%v---\n",x,y)
	}

}

// select statements may introduce
// flakiness: whether this test contains
// a race depends on the scheduling
// (some may argue that the code contains
// this race by definition)
/*
func TestFlakyDefault(t *testing.T) {
	var x int
	c := make(chan bool, 1)
	done := make(chan bool, 1)
	go func() {
		select {
		case <-c:
			x = 2
		default:
			x = 3
		}
		done <- true
	}()
	x = 1
	c <- true
	_ = x
	<-done
}*/