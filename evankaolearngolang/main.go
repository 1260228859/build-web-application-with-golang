package main

import (
	// "encoding/json"
	"fmt"
	// "reflect"
	// "mymath"
	// "strconv"

)

// import "github.com/astaxie/beedb"



func main1(a int) int {
	b := a + 20
	return b
}

func main2() {
	// a := "ssssss"
	// var b = "aaaaaa"
	// var c string = "ddddddd"
	// fmt.Println(a, b, c)

	// a = "1" + a[:1]
	// s := []byte(a)
	// fmt.Println(a, s)

	// s1 := string(s)
	// fmt.Println(s1, s[0])

	// s2 := strconv.Itoa(49)
	// fmt.Println(s2)

	// a1 := 16
	// resu := strconv.Itoa(a1)
	// fmt.Println(resu, reflect.TypeOf(resu), reflect.TypeOf(a1))

	// 高位转低位可能会有问题
	// var a2 int8 = 16
	// a3 := int8(a1)
	// fmt.Println(a2, a3)

	// resu := "11100200"
	// a4, err := strconv.Atoi(resu)
	// if err == nil {
	// 	fmt.Println("没有错误", a4, reflect.TypeOf(a4))
	// } else {
	// 	fmt.Println("错误", nil)
	// }
	// fmt.Println(a4)

	// v1 := 99
	// v2 := strconv.FormatInt(int64(v1), 16)
	// fmt.Println(v2)

	// fmt.Println(v2, reflect.TypeOf(v2))

	// data := "1001000101"
	// data1, _ := strconv.ParseInt(data, 2, 0)

	// var data1 int
	// fmt.Println(data1)

	// var v2 *int
	// v3 := new(int)
	// fmt.Println(v2, v3, *v3, reflect.TypeOf(v3))
	// m := `sdf
	// sdf
	// `
}

func main3()  {
	// 字符串
	s := "12323高件套"
	a := []byte(s)
	b := []rune(s)
	fmt.Println(a, b)
	
	// 
	type Person struct {
		name string
		age int
	}

	person := Person{name: "wwww", age: 23}
	fmt.Println(person, person.age)

	//
	astr := make([]int, 10, 20)
	fmt.Println(astr)
}

func main() {
	// a := main1(2)
	// // fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	// fmt.Printf("hollo word", a)
	main3()
}
