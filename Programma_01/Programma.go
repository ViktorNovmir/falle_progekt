package main

import "fmt"

func main() {
	var i int = 10
	число := i&1 == 0
	результат := число
	fmt.Println("Число чётное", результат, i)

	Число1 := 10
	Число2 := 5
	РезультатИНЕ := Число1 &^ Число2
	fmt.Printf("Число1: %08b\n", Число1) // 00000101
	fmt.Printf("Число2: %08b\n", Число2) // 00000011
	fmt.Printf("РезультатИНЕ: %08b\n", РезультатИНЕ)
	fmt.Println(РезультатИНЕ)

	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	z := copy(y, x)
	fmt.Println("скопировал в x 4 елемента", z)
	fmt.Println("копия в y", y)

	var s string = "hello@!"
	runs := []rune(s)
	yy := make([]rune, 7)
	zz := copy(yy, runs)
	fmt.Println("скопировал", zz)

	fmt.Println("Я зделал так", runs)

}
