
package main

import (
	"fmt"
	"unsafe"
)

var lll = 1

var printl = fmt.Println

func learning() {
	a := "a"
	fmt.Println(a, "hello, world")
	a = "asdfghqwertqwertqweqrqweqtqqwtyyeuritoyusdghf"
	fmt.Println(&a)
	fmt.Println(unsafe.Sizeof(a))

	p := &a;
	printl("p",p)
	fmt.Printf("%T",p) // 格式化打印

	const (
		s1 = iota
		s2
		s3
	)
	fmt.Println(s1, s2, s3)

	const (
		x1 = 2 * iota
		x2
		x3
	)
	fmt.Println(x1, x2, x3)

	s := 'a'
	fmt.Printf("%T", s)
	printl()
	printl(s)

	p1 := &s
	printl(p1)

	// goto
	var i = 0
	LOOP_LABEL: if (i < 20) {
		printl(i)
		i++
		goto LOOP_LABEL
	}

	// for
	i = 0
	for true {
		if (i < 10) {
			i++
			printl(i)
		} else {
			break
		}
	}

	outer := "foo"
	foo := func() {
		printl(outer)
	}
	foo()

	sum(1, 2)
	two(1, 2)

	var c Circle // 新建对象
	c.radius = 2
	c.area() // 调用类方法

	adder := createAdder()
	fmt.Printf("%T", adder)
	printl()
	adder()
	adder()
	adder()

/*	t1 := 1
	t2 := "2"
	t3 := t1 + t2 */
	// t1 t2 不可加

	deferTest()
}

func createAdder() func() int { // 匿名函数返回类型声明
	i := 100
	return func() int {
		i++
		printl(i)
		return i
	}
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func two(num1, num2 int) (int, int) {
	return num1,num2
}

// 声明类
type Circle struct {
	radius float64
}

// 类方法
func (c Circle) area () {
	printl(c.radius * c.radius * 3.14)
}

func deferTest() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func f1() (r int) {
	r = 9
	t := 0
	defer func() {
		t = t + 5
	}()
	// 就算return后面跟着空，但由于返回值已经定义为r，就会返回r=9
	return
}

func f2() (r int) {
	r = 9
	t := 0
	defer func() {
		t = t + 5
	}()
	// return t 实际上是令r=t，再返回r
	// defer在r=t=0后， return r 前执行，故返回0
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	// 虽然defer在return前操作了r=1，但操作的是参数r，不能改变返回值，仍是1
	return 1
}

func f4() (r int) {
	defer func() {
		r = r + 5
	}()
	// defer操作了返回值r，成功改变
	return 1
}