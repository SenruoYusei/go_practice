package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// #1
// type Vertex struct {
// 	X, Y float64
// }

// #1
// func (v Vertex) Abs() float64 { // Vertex 型にメソッドを定義、 Vertex 型のレシーバ v を持つ
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// #2
// func Abs(v Vertex) float64 { // 通常の関数、Vertex 型のレシーバ引数を持つ
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// #3
// type MyFloat float64

// func (f MyFloat) Abs() float64 { // これは MyFloat を定義しているパッケージでのみ定義可能
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// #4
type Vertex struct {
	X, Y float64
}

// #4
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y + v.Y)
// }

// #4
// func (v *Vertex) Scale(f float64) { // *Vertex に定義されている、レシーバが指す変数を変更できる
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// #5
// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// #6
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// #7
// func (v Vertex) Abs() float64 { // 特定の変数の型を与えないといけない
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func AbsFunc(v Vertex) float64 { // 変数、またはポインタを与えれば良い
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// #8
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// #9
// type Abser interface {
// 	Abs() float64
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// #10 #11 #12
type I interface {
	M()
}

type T struct {
	S string
}

// #10
// func (t T) M() {
// 	fmt.Println(t.S)
// }

// #11
// func (t *T) M() {
// 	fmt.Println(t.S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func describe(i I) { // interface を実装している型であれば使用できる
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// #12
// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}
// 	fmt.Println(t.S)
// }

// #13
// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// #15
// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

// #16
// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// # interface の中に組み込みで
// # String() string
// # Error() string
// # などがある

// #17
// type IPAddr [4]byte

// func (ipa IPAddr) String() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
// }

// #18
// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s", e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// #19
// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 1 {
// 		return x, ErrNegativeSqrt(x)
// 	}
// 	z := float64(1)
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z, nil
// }

// const (
// 	scKey    = 13
// 	alphabet = "abcdefghijklmnopqrstuvwxyz"
// 	table    = alphabet + alphabet
// )

// type rot13Reader struct {
// 	r io.Reader
// }

// func (r rot13Reader) Read(p []byte) (int, error) {
// 	n, err := r.r.Read(p)
// 	if err != nil {
// 		return n, err
// 	}
// 	for i := range p {
// 		p[i] = r.rotate(p[i], scKey)
// 	}
// 	return n, err
// }

// func (r *rot13Reader) rotate(b byte, key int) byte {
// 	ch := rune(b)
// 	if !unicode.IsLetter(ch) {
// 		return b
// 	}

// 	key = key % len(alphabet)

// 	isUpper := false
// 	if unicode.IsUpper(ch) {
// 		isUpper = true
// 		ch = unicode.ToLower(ch)
// 	}

// 	idx := byte(ch) - 'a' + byte(key)
// 	rotatedCh := rune(table[idx])
// 	if isUpper {
// 		rotatedCh = unicode.ToUpper(rotatedCh)
// 	}
// 	return byte(rotatedCh)
// }

type Image struct {
	w int
	h int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// #1
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())

	// #2
	// v := Vertex{3, 4}
	// fmt.Println(Abs(v))

	// #3
	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f)
	// fmt.Println(f.Abs())

	// #4
	// v := Vertex{3, 4}
	// v.Scale(10)
	// fmt.Println(v)
	// fmt.Println(v.Abs())

	// #5
	// v := Vertex{3, 4}
	// fmt.Println(v)
	// Scale(&v, 10)
	// fmt.Println(v)
	// fmt.Println(Abs(v))

	// #6
	// v := Vertex{3, 4}
	// v.Scale(2) // &v.Scale(2) として呼び出す
	// ScaleFunc(&v, 10)

	// p := &Vertex{4, 3}
	// p.Scale(3)
	// ScaleFunc(p, 8)

	// fmt.Println(v, p)

	// #7
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())
	// fmt.Println(AbsFunc(v)) // 特定の変数の型を与えないといけない

	// p := &Vertex{4, 3}
	// fmt.Println(p.Abs())
	// fmt.Println(AbsFunc(*p))

	// #8
	// v := &Vertex{3, 4}
	// fmt.Printf("Before scallig: %+v, Abs: %v\n", v, v.Abs())
	// v.Scale(5)
	// fmt.Printf("After scallig: %+v, Abs: %v\n", v, v.Abs())

	// #9
	// var a Abser
	// f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	// a = f
	// a = &v

	// a = v
	// fmt.Println(a.Abs())

	// #10
	// var i I = T{"Hello"}
	// i.M()

	// #11
	// var i I
	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// i = F(math.Pi)
	// describe(i)
	// i.M()

	// #12
	// var i I

	// var t *T
	// i = t
	// describe(i)
	// i.M()

	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// #13
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "hello"
	// describe(i)

	// #14
	// var i interface{} = "hello"

	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// // f = i.(float64)
	// // fmt.Println(f)

	// #15
	// do(21)
	// do("hello")
	// do(true)

	// #16
	// a := Person{"Arthur Dent", 42}
	// z := Person{"Zaphod Beeblebrox", 9001}
	// fmt.Println(a, z)

	// #17
	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }

	// #18
	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// }

	// #19
	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))

	// #20
	// r := strings.NewReader("Hello, Reader!")

	// b := make([]byte, 8)
	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	// reader.Validate(MyReader{})

	// s := strings.NewReader(("Lbh penpxrq gur pbqr!"))
	// r := rot13Reader{s}
	// io.Copy(os.Stdout, &r)

	// m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(0, 0).RGBA())

	m := Image{100, 200}
	pic.ShowImage(m)
}
