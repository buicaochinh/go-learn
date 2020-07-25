package main

// func main() {
// 	fmt.Println("Hello Golang")
// 	var number int = 10
// 	fmt.Println(number)

// 	var email = "buicaochinh@gmail.com"
// 	fmt.Println(email)

// 	var (
// 		name    string = "name1"
// 		address string = "address1"
// 		age     int    = 25
// 	)

// 	fmt.Println(name, address, age)

// 	language := "golang"

// 	fmt.Println(language)
// }

// func main() {
// 	//Bool
// 	var myBool bool = true
// 	fmt.Println(myBool)

// 	var mySecondBool bool = false

// 	fmt.Println(mySecondBool)

// 	//String

// 	var myString string = "hello"
// 	fmt.Println(myString)

// 	//int
// 	var myInt int = 69
// 	fmt.Println(myInt)

// 	// int8, 16, 32, 64
// 	// 1. Range
// 	// 2. Bits

// 	// 1. Range
// 	fmt.Println(math.MinInt8)
// 	fmt.Println(math.MaxInt8)
// 	fmt.Println(math.MinInt16)
// 	fmt.Println(math.MaxInt16)
// 	fmt.Println(math.MinInt32)
// 	fmt.Println(math.MaxInt32)
// 	fmt.Println(math.MinInt64)
// 	fmt.Println(math.MaxInt64)

// 	// 2. Bits
// 	fmt.Println("------------------------")
// 	fmt.Println(bits.OnesCount8(math.MaxUint8))

// 	// Bytes

// 	var myByte byte = 1
// 	fmt.Println(myByte)
// 	fmt.Printf("%T\n", myByte)

// 	var a byte = 'A' // byte thuộc kiểu Alias: bí danh, có thể chuyển đổi ASCII
// 	fmt.Printf("%X\n", a)

// 	//float
// 	var myFloat float64 = 10.01
// 	fmt.Println(myFloat)

// 	//complex z= a+bi
// 	var z1 complex64 = 10 + 1i
// 	var z2 complex64 = 10 + 1i
// 	fmt.Println(z1 + z2)

// 	//Rune
// 	var myStr string = "Nhẫn"
// 	runes := []rune(myStr)
// 	for i := 0; i < len(runes); i++ {
// 		fmt.Printf("%c\n", runes[i])
// 	}

//}

// func main() {
// 	var a int = 1
// 	var b float64 = 2.1

// 	fmt.Println(float64(a) + b)
// }

// func main() {
// 	// if else

// 	number := 10
// 	if number == 10 {
// 		fmt.Println("number = 10")
// 	} else {
// 		fmt.Println("number != 10")
// 	}

// 	// if statement; condition {//code}
// 	if a := 99; a > 100 {
// 		fmt.Println("a > 100")
// 	} else if a == 100 {
// 		fmt.Println("a = 100")
// 	} else {
// 		fmt.Println("a < 100")
// 	}
// 	// Fallthrough
// 	switch number {
// 	case 1:
// 		fmt.Println("number = 1")
// 		fallthrough
// 	case 10:
// 		fmt.Println("number = 10")
// 		fallthrough
// 	case 2:
// 		fmt.Println("number = 2")
// 		fallthrough
// 	case 3:
// 		fmt.Println("number = 3")
// 	}

// 	// break, goto
// 	switch number {
// 	case 1:
// 		fmt.Println("number = 1")
// 		fallthrough
// 	case 10:
// 		if number == 10 {
// 			goto handleNumberEqual10
// 		}
// 	handleNumberEqual10:
// 		fmt.Println("handleNumber = 10")
// 	case 2:
// 		fmt.Println("number = 2")
// 		fallthrough
// 	case 3:
// 		fmt.Println("number = 3")
// 	}
// }

// func main() {
// 	// loops, break, continue
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}

// 	// while go khong ho tro while
// 	j := 0
// 	for j < 10 {
// 		fmt.Println(j)
// 		j++
// 	}

// 	// infinite loop
// 	/*
// 		for {
// 			fmt.Println("infinite loop")
// 		}
// 	*/
// }

// func func_name(params) return type {//code}

// func hello() {
// 	fmt.Println("Hello")
// }

// func sayHello(name string) {
// 	fmt.Println("Hello", name)
// }

// func greeting(name string) string {
// 	result := fmt.Sprintf("Hello %s", name)
// 	return result
// }

// //Multiple return values
// func rectInfo(w, h int) (int, int, int) {
// 	area := w * h
// 	return w, h, area
// }

// // Named return values
// func rectInfo2(w, h int) (width int, height int, isSquare bool) {
// 	isSquare = w == h
// 	return w, h, isSquare
// }

// func main() {
// 	hello()
// 	sayHello("Charlie")
// 	fmt.Println(greeting("Alice"))
// 	w, h, area := rectInfo(100, 200)
// 	fmt.Println("Width =", w, "Height =", h, "area =", area)
// 	fmt.Println(rectInfo2(300, 200))
// }

// func main() {
// 	var slice []int
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))

// 	array := [...]int{1, 2, 3, 4, 5}
// 	slice = array[1:4]
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))

// 	//make, copy, append
// 	slice1 := make([]int /*len*/, 2 /*cap*/, 5)
// 	fmt.Println(slice1)

// 	var slice2 []int
// 	slice2 = append(slice2, 100)
// 	fmt.Println(slice2)

// 	// copy

// 	src := []string{"A", "B", "C", "D"}
// 	dest := make([]string, 2)

// 	number := copy(dest, src)
// 	fmt.Println(number)
// 	fmt.Println(dest)

// 	//delete item with index=1
// 	src = append(src[:1], src[2:]...)
// 	fmt.Println(src)
// }

/*
	Variadic function
	1. Khai bao valiadic function
	2. pass 1 slice vao 1 valiadic function
	3. change slice
	append(slice []Type, elems ...Type) []Type
	append([]Type, args, arg2, ..., argsN)
*/
// func addItem(item int, list ...int) {
// 	list = append(list, item, 11, 22, 33, 44, 55)
// 	fmt.Println(list)
// }

// func change(list ...int) {
// 	list[0] = 999
// }

// func main() {
// 	addItem(1, 100, 200, 300, 400)
// 	// 100, 200, 300, 400 -> []int {100, 200, 300, 400}

// 	var list = []int{1, 2, 3, 4}
// 	addItem(100, list...)

// 	change(list...)
// 	fmt.Println(list)
// }

// func main() {
// 	var myMap = make(map[string]int)
// 	fmt.Println(myMap)

// 	var myMap1 map[string]int
// 	fmt.Println(myMap1)

// 	if myMap != nil {
// 		fmt.Println("myMap != nil")
// 	} else {
// 		fmt.Println("myMap = nil")
// 	}

// 	//khai bao voi gia tri khoi tao
// 	myMap2 := map[string]int{
// 		"key1": 1,
// 		"key2": 2,
// 		"key3": 3,
// 		"key4": 4,
// 	}
// 	fmt.Println(myMap2)

// 	//them 1 phan tu vao map
// 	myMap2["key5"] = 5
// 	fmt.Println(myMap2)

// 	// xóa một phần tử
// 	delete(myMap2, "key1")
// 	fmt.Println(myMap2)
// 	fmt.Println(len(myMap2))
// 	// map, slice là một reference type
// 	myMap3 := myMap2
// 	myMap3["key5"] = 1000
// 	fmt.Println(myMap2)

// 	//Truy cap 1 phan tu trong map
// 	value := myMap2["key3"]
// 	fmt.Println(value)

// 	value, found := myMap2["key2"]
// 	if found {
// 		fmt.Println(found)
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println(found)
// 		fmt.Println("Not Found")
// 	}

// 	// map khong compare duoc (ngoai tru gia tri nil)
// }

// con tro

// func applyPointer(pointer *int) {
// 	*pointer = 777
// }

// func main() {
// 	a := 100
// 	var pointer *int
// 	pointer = &a
// 	fmt.Println(*pointer)
// 	fmt.Printf("%T\n", pointer)

// 	b := 123
// 	p2 := new(int)
// 	p2 = &b
// 	fmt.Println(p2)
// 	// zero value

// 	var pointer1 *int
// 	fmt.Println(pointer1) // nil

// 	pointer2 := new(int)
// 	fmt.Println(pointer2) // khong phai nil

// 	// pointer -> array

// 	array := [3]int{1, 2, 3}
// 	var pointer3 *[3]int
// 	pointer3 = &array
// 	fmt.Println(pointer3)
// 	c := 30
// 	var pointer4 *int = &c
// 	applyPointer(pointer4)
// 	fmt.Println(c)
// }

// Struct trong Golang

// type Student struct {
// 	id   int
// 	name string
// 	info StudentInfo
// }

// type StudentInfo struct {
// 	class string
// 	email string
// 	age   int
// }

// func main() {
// 	// named
// 	st1 := Student{
// 		id:   123,
// 		name: "Charlie",
// 	}

// 	fmt.Println(st1)
// 	fmt.Println(st1.id)

// 	// st2 := Student{456, "Alice"}
// 	// fmt.Println(st2)

// 	// var st3 Student = struct {
// 	// 	id   int
// 	// 	name string
// 	// }{
// 	// 	id:   777,
// 	// 	name: "bill",
// 	// }

// 	// fmt.Println(st3)

// 	// anonymous struct
// 	var anonymous = struct {
// 		email string
// 		age   int
// 	}{
// 		email: "charlie@gmail.com",
// 		age:   27,
// 	}

// 	fmt.Println(anonymous)

// 	// pointer tro toi mot struct
// 	pointer := &Student{
// 		id:   789,
// 		name: "Henry",
// 	}

// 	fmt.Println(pointer)
// 	fmt.Println(pointer.id)
// 	fmt.Println(pointer.name)

// 	// anonymous fields
// 	type NoName struct {
// 		string
// 		int
// 	}

// 	n := NoName{
// 		"Nguyen Van A",
// 		5555,
// 	}

// 	fmt.Println(n)

// 	// Struct lồng struct
// 	student := Student{
// 		id:   123,
// 		name: "Charlie",
// 		info: StudentInfo{
// 			class: "K63CE",
// 			email: "charlie@gmail.com",
// 			age:   20,
// 		},
// 	}
// 	fmt.Println(student)

// 	// compare 2 struct

// }

/*
	METHOD
	func (t Type) methodName(params) return_type {// code}
	(t Type) +. receiver
	1. Value Receiver
	2. Pointer Receiver

	Không nhất thiết phải là struct, có thể là non-struct
*/

// type Student struct {
// 	name string
// }

// func (s Student) getName() string {
// 	return s.name
// }

// // 1. Value Receiver
// func (s Student) changeName() {
// 	s.name = "Robin"
// }

// //2. Pointer Receiver

// func (s *Student) changeName2() {
// 	s.name = "Robbin"
// }

// func main() {
// 	student := Student{"Ryan"}

// 	student.changeName()
// 	fmt.Println(student.name)
// 	student.changeName2()
// 	fmt.Println(student.name)

// }

/*
	INTERFACE,
	multiple interfaces,
	embed interface
	empty interface
*/

// type Animal interface {
// 	speak()
// }

// type Movement interface {
// 	move()
// }

// type NextAnimal interface {
// 	Movement
// 	Animal
// }

// type Dog struct{}

// func (d Dog) speak() {
// 	fmt.Println("Woawww Woaww")
// }

// func (d Dog) move() {
// 	fmt.Println("Dogs have 4 legs")
// }

// func goout(i interface{}) {
// 	fmt.Println(i)
// }

// func main() {
// 	dog := Dog{}
// 	var na NextAnimal = dog

// 	na.speak()
// 	na.move()
// 	goout(10)
// 	goout(10.1)
// }

func main() {

}
