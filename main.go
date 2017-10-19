package main

import (
	"fmt"
	"unsafe"
)

var printl = fmt.Println

func main() {
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