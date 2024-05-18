package main

import (
	"fmt"
)

// #1
// var c, python, java bool

// #2
// var i, j int = 1, 2

// #3
// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// #4
// const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// #0
	fmt.Println(split(17))

	// #1
	// var i int
	// fmt.Println(i, c, python, java)

	// #2
	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	// #3
	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, k, c, python, java)

	// #3
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

	// v := 0.867 + 0.5i
	// fmt.Printf("v is of type %T\n", v)

	// #4
	// const World = "世界"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")
	// const Truth = true
	// fmt.Println("Go Rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
