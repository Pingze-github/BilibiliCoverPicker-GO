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

	arrayTest();

	pointerTest();

	structTest();

	sliceTest();
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

func arrayTest() {
	var arr [10] float32
	arr2 := [5] int {1,2,3,4,5}
	arr[4] = 1
	fmt.Println(arr);
	fmt.Println(arr2[1]);

	var matrix =  [2][3] int {
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])

	init := func (m [9]int) [9]int {
		var i int
		for i = 0; i < 9; i++ {
			m[i] = i;
		}
		return m
	}
	var m2 [9] int // 不定长数组
	fmt.Println("m2", m2)
	var m3 = init(m2)
	fmt.Println("m3", m3)
	fmt.Println("m2", m2)
	// m2作为参数传入函数，是值传递，m2 != m3
}

func pointerTest() {
	var a = 1
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(*p)

	var ptr *int
	fmt.Println(ptr)
	fmt.Println(ptr == nil)
}

func structTest() {
	type Book struct {
		title string
		author string
	}
	var printBook = func (book Book) {
		fmt.Println("title:", book.title)
		fmt.Println("author:", book.author)
	}
	var book1 Book
	book1.title = "挪威的森林"
	book1.author = "村上春树"
	printBook(book1)
}

func sliceTest() {
	// 切片是可变长度的数组。

	// 定义切片只需要不指定数组长度
	var list []string
	fmt.Println(list)

	// 或者使用make，给出更多设置
	var list2 = make([]string, 5, 10) // 初始长度5 最大长度10
	fmt.Println(list2)
	fmt.Println(unsafe.Sizeof(list2))

	// 初始化
	var list3 = []int{1,2,3}  // len = cap = 3
	fmt.Println(list3)

	// 从数组创建
	var arr = [5]int{1, 2, 3, 4, 5}
	var list4 = arr[:]
	var list5 = arr[1:3]
	fmt.Println(list4)
	fmt.Println(list5)
	fmt.Println(list == nil) //未初始化时nil（指针也是）
	fmt.Println(len(list))
	fmt.Println(len(list2))
	fmt.Println(cap(list2))

	// cap 容量 是slice的容积单位，cap值为一个单位，容量增加时，会增加一个cap的大小
	// 默认cap=1

	// append
	//var sl = make([]int, 0, 1)
	var sl [] int
	sl = append(sl, 1)
	sl = append(sl, 2, 3, 4)
	fmt.Println(sl, cap(sl)) // cap 5
	sl = append(sl, 5, 6, 7)
	fmt.Println(sl, cap(sl)) // cap 10

	var sl2 = make([]int, 10)
	copy(sl2, sl) // sl => sl2
	// sl2 不够大，有多大copy多大
	fmt.Println(sl2)

}