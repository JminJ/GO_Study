package main

import "fmt"

/*
1.  var를 통해 변수를 선언한다. var _name_ _type_ = N의 형식으로 선언이 가능하며, 자동 타입 추적이 가능하므로 var _name_ = N의 형식으로 선언 또한 가능하다.
*/
func part_1() {
	var a int = 15
	var b float32 = 1.
	var c = 12

	fmt.Printf("type of a: %T \n", a)
	fmt.Printf("type of b: %T \n", b)
	fmt.Printf("type of c: %T \n", c)
}

/*
2. 다중 선언 또한 가능하다.
*/
func part_2() {
	var d, e, f string
	// var d, e, f string = "반갑습니다", "어떻게 지내셨나요", "저는 잘 지냈죠" <== 해당 방식으로 다중 선언 가능
	d = "반갑습니다"
	e = "어떻게 지내셨나요"
	f = "저는 잘 지냈죠"

	fmt.Printf("%s \n", d)
	fmt.Printf("%s \n", e)
	fmt.Printf("%s \n", f)
}

func main() {
	// part 1 result
	fmt.Println("===== Part1 result =====")
	part_1()

	// part 2 result
	fmt.Println("\n===== Part2 result =====")
	part_2()
}
