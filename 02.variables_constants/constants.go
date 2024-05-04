package main

import "fmt"

/*
1. 상수는 변수와 다르게 한번 선언 수 수정이 불가능하다. const _name_ _type_ = N의 형식으로 선언가능하다.
*/
func part_1() {
	const a int = 12
	const b = "반갑습니다." // 변수와 같이 자동 타입 추론이 적용된다.

	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
}

/*
2. 또한 묶어 다중 선언이 가능하다. 특히 1씩 증가하는 상수를 선언하고 싶을 시 "iota" 키워드를 사용하면 된다고 한다.
*/
func part_2() {
	const (
		c = "다중"
		d = "선언"
		e = 123
	)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("Type of e: %T\n", e)

	const (
		iota1 = iota
		iota2
		iota3
	)
	fmt.Println("[iota examples]")
	fmt.Println(iota1)
	fmt.Println(iota2)
	fmt.Println(iota3)
}

func main() {
	// part 1 result
	fmt.Println("===== Part1 result =====")
	part_1()

	// part 2 result
	fmt.Println("\n===== Part2 result =====")
	part_2()
}
