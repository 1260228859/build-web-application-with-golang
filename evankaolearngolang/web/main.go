package main

import (
	// "encoding/json"
	"fmt"
	"time"
	// "go/importer"
	// "reflect"
	// "mymath"
	// "strconv"
	// "time"
)

// import "github.com/astaxie/beedb"

// func main1(a int) int {
// 	b := a + 20
// 	return b
// }

// func main2() {
// 	// a := "ssssss"
// 	// var b = "aaaaaa"
// 	// var c string = "ddddddd"
// 	// fmt.Println(a, b, c)

// 	// a = "1" + a[:1]
// 	// s := []byte(a)
// 	// fmt.Println(a, s)

// 	// s1 := string(s)
// 	// fmt.Println(s1, s[0])

// 	// s2 := strconv.Itoa(49)
// 	// fmt.Println(s2)

// 	// a1 := 16
// 	// resu := strconv.Itoa(a1)
// 	// fmt.Println(resu, reflect.TypeOf(resu), reflect.TypeOf(a1))

// 	// 高位转低位可能会有问题
// 	// var a2 int8 = 16
// 	// a3 := int8(a1)
// 	// fmt.Println(a2, a3)

// 	// resu := "11100200"
// 	// a4, err := strconv.Atoi(resu)
// 	// if err == nil {
// 	// 	fmt.Println("没有错误", a4, reflect.TypeOf(a4))
// 	// } else {
// 	// 	fmt.Println("错误", nil)
// 	// }
// 	// fmt.Println(a4)

// 	// v1 := 99
// 	// v2 := strconv.FormatInt(int64(v1), 16)
// 	// fmt.Println(v2)

// 	// fmt.Println(v2, reflect.TypeOf(v2))

// 	// data := "1001000101"
// 	// data1, _ := strconv.ParseInt(data, 2, 0)

// 	// var data1 int
// 	// fmt.Println(data1)

// 	// var v2 *int
// 	// v3 := new(int)
// 	// fmt.Println(v2, v3, *v3, reflect.TypeOf(v3))
// 	// m := `sdf
// 	// sdf
// 	// `
// }

// func main3()  {
// 	// 字符串
// 	s := "12323高件套"
// 	a := []byte(s)
// 	b := []rune(s)
// 	fmt.Println(a, b)

// 	//

// 	person := Person{name: "wwww", age: 23}
// 	fmt.Println(person, person.age)

// 	//
// 	astr := make([]int, 10, 20)
// 	fmt.Println(astr)
// 	astr[6] = 9
// 	c := append(astr, 8, 9)
// 	fmt.Println(c)

// }

// // 函数作为参数传递
// type fun1 func(int)int

// func a222(a int) int {
// 	return a + 2
// }

// func aaaa(a int, f fun1) int {
// 	return f(a)
// }

type person struct {
	name string
	age  int
}

type gf []int
type evankao struct {
	person
	money int
	gf
	int
}

func structure(P *evankao) {
	fmt.Println(P.age, *P)

}

func (P evankao) funa() {
	P.age = 999998
	fmt.Println(P)
}

func (P evankao) funb() {
	P.age = 999997
	fmt.Println(P)
}

type fc interface {
	funa()
	funb()
}

var arg interface{}

func isBool(arg interface{}) {

	switch arg_type := arg.(type) {
	case int:
		fmt.Println("sw-int", arg_type)
	case string:
		fmt.Println("sw-string", arg_type)
	}

	val, ok := arg.(int)
	fmt.Println(val, ok)
}

func main() {
	// a := main1(2)
	// // fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	// fmt.Printf("hollo word", a)

	// gf := append(make(gf, 9, 10), 9, 9)
	// Evan := evankao{person{"gao", 99}, 99999, gf, 2}
	// Evan.gf = []int{1,3,4,5}
	// Evan.funa()
	// fmt.Println(Evan.age, Evan.int)
	// structure(&Evan)

	// interface

	// var fc_i fc
	// fc_i = Evan
	// fc_i.funb()
	// fmt.Println(Evan)

	// interface任意参数传递
	// arg = 12
	// isBool(arg)

	// 反射
	// float1 := 5.234
	// fmt.Println(reflect.TypeOf(float1))
	// v := reflect.ValueOf(&float1)
	// fmt.Println(v.Type(), v.Kind(), v.Float())
	// p := v.Elem()
	// p.SetFloat(7.9)
	// fmt.Println(p)
	// go say("a")
	// go say("a")
	// go say("a")
	// go say("a")
	// fmt.Println("aaaaaaaaa")
	// say("w")

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	// ch2 <- 3
	// ch2 <- 3
	// ch1 <- 2
	// ch1 <- 3
	// v:= <-ch1
	// fmt.Println(v)
	// v1 := <-ch1
	// fmt.Println(v1)

	// selsct
	// go say(2999, ch1, ch2)
	for {
		select {
		case x := <-ch1:
			fmt.Println("chann-get", x)
		case x := <-ch2:
			fmt.Println("quit-ch2", x)
			return
		case <-time.After(5 * time.Second):
			fmt.Println("超时退出")
			return
		}
	}

}

func say(s int, ch1 chan<- int, ch2 chan int) {
	// a222 := <- ch2
	// fmt.Println(a222)
	for i := 0; i < 5; i++ {

		fmt.Println(i)
		ch1 <- i
	}
	ch2 <- s
}
