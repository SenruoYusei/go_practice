package main

import (
	"fmt"
)

// #1
// type Vertex struct {
// 	X int
// 	Y int
// }

// var (
// 	v1 = Vertex{1, 2}
// 	v2 = Vertex{X: 1}
// 	v3 = Vertex{}
// 	p  = &Vertex{1, 2}
// )

// #5
// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// #2, #4
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// #3
// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
// }

// #5
// func Pic(dx, dy int) [][]uint8 {
// 	result := make([][]uint8, dy)
// 	for i := 0; i < dy; i++ {
// 		result[i] = make([]uint8, dx)
// 	}
// 	for i := 0; i < dx; i ++ {
// 		for j := 0; j < dy; j++ {
// 			result[i][j] = uint8(i * j)
// 		}
// 	}
// 	return result
// }

// #6 #7 #8
type Vertex struct {
	Lat, Long float64
}

// #6
// var m map[string]Vertex
// #7
//
//	var m = map[string]Vertex{
//		"Bell Labs": Vertex{40.68433, -74.39967},
//		"Google":    Vertex{37.42202, -122.08408},
//	}
//
// #8
// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37.42202, -122.08408},
// }

// #9
// func WordCount(s string) map[string]int {
// 	fields := strings.Fields(s)
// 	result := make(map[string]int)
// 	for _, field := range fields {
// 		_, ok := result[field]
// 		switch ok {
// 			case false:
// 				result[field] = 1
// 			case true:
// 				result[field] = result[field] + 1
// 		}
// 	}
// 	return result
// }

// #10
// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// #11
// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

func fibonacci() func() int {
	prev := 0
	fmt.Println(prev)
	cur := 1
	fmt.Println(cur)
	return func() int {
		result := prev + cur
		prev = cur
		cur = result
		return result
	}
}

func main() {
	// i, j := 42, 2701

	// p := &i
	// fmt.Println(*p)
	// *p = 21
	// fmt.Println(i)

	// p = &j
	// *p = *p / 37
	// fmt.Println(j)

	// #1
	// fmt.Println(Vertex{1, 2})

	// v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v)

	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)

	// fmt.Println(v1, p, v2, v3)

	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// primes := [6]int{2, 3, 5, 7, 11, 13}

	// var s []int = primes[1:4]
	// fmt.Println(s)

	// # スライスは既存の配列の参照をしている、変更すると元の配列も変更される
	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)
	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)

	// #2
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// s = s[:0]
	// printSlice(s)

	// s = s[2:4] // 削除したら前半の値は。それ以降、取ってこれない
	// printSlice(s)

	// s = s[:4] // 直前のスライスで持っていない後ろの値も取ってこれる
	// printSlice(s)

	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }

	// #3
	// a := make([]int, 5)
	// printSlice("a", a)

	// b := make([]int, 0, 5)
	// printSlice("b", b)

	// c := b[:2]
	// printSlice("c", c)

	// d := c[2:5]
	// printSlice("d", d)

	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	// #4
	// var s []int
	// printSlice(s)

	// s = append(s, 0)
	// printSlice(s)

	// s = append(s, 1)
	// printSlice(s)

	// s = append(s, 2, 3, 4)
	// printSlice(s)

	// #5
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// pow := make([]int, 10)
	// for i := range pow {
	// 	pow[i] = 1 << uint(i)
	// }
	// for _, value := range pow {
	// 	fmt.Printf("%d\n", value)
	// }

	// #5
	// pic.Show(Pic)

	// #6
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{40.68433, -74.39967}
	// fmt.Println(m["Bell Labs"])

	// #7 #8
	// fmt.Println(m)

	// m := make(map[string]int)

	// m["Answer"] = 42
	// fmt.Println("The value: ", m["Answer"])

	// m["Answer"] = 48
	// fmt.Println("The value: ", m["Answer"])

	// delete(m, "Answer")
	// fmt.Println("The value: ", m["Answer"])

	// v, ok := m["Answer"]
	// fmt.Println("The value: ", v, "Present", ok)

	// #9
	// wc.Test(WordCount)

	// #10
	// hypot := func(x, y float64) float64 {
	// 	return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))

	// fmt.Println(compute(hypot))
	// fmt.Println(compute(math.Pow))

	// #11
	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
