package main

import (
	"bytes"
	"fmt"
	"sync"
)

// "strconv"

/*
	VISIBILTIY
		-Package level variable name starts with an uppercase whereas block level starts with a lowercase
		-NO private scope
*/
// var PackageLevelVariable int = 0

var needsDataType string = "Global variable"

func main() {
	/*
		Variables section
	*/

	//Declaring variable type explicit
	// var i int = 42
	//Declaring variable type implicit (int/float64)
	// j := 32.33

	//Type casting
	// i = int(j)
	// j = float64(i)

	//Converting int to string
	// text := "Hello World"
	// text = strconv.Itoa(i)

	//Converting string to int
	// text := "43"
	// i = strconv.ParseInt()

	// fmt.Println(i, j, text)
	/*
		Variable section ends
	*/

	// complextNumber := complex(32, 3)
	// fmt.Printf("%v, %T", complextNumber, complextNumber)

	// runeToByte := 'Z'
	// c := byte(runeToByte)

	// stringToByte := "Convert this to byte"
	// b := []byte(stringToByte)

	// fmt.Println(b, string(c))
	// fmt.Println(runeToByte)

	/*
		Constants
	*/

	// const _THIS_IS_A_CONSTANT = "Doesn't change"
	// fmt.Printf("%T", _THIS_IS_A_CONSTANT)

	//iota auto increments itself in a constant block
	// const (
	// 	_SUCCESSFUL = iota + 200
	// 	_LIMIT_EXCEEDED
	// 	_TIMEOUT
	// 	_SESSION_ENDED
	// )

	// fmt.Println(_SUCCESSFUL)
	// fmt.Println(_LIMIT_EXCEEDED)
	// fmt.Println(_TIMEOUT)
	// fmt.Println(_SESSION_ENDED)

	/*
		Constants ends
	*/

	/*
		Arrays and slices
	*/
	//Fixed length
	// grades := [...]int{1, 2, 3}
	// grades := [3]int{1, 2, 3}

	// var students []string //Slice
	// students := [3]string{"Alice", "Bob", "Charlie"}

	//Adding items to array
	// students = append(students, "Alice", "Bob", "Charlie")

	//Only works on fixed size arrays
	// studentsCopy := &students //Type is *[3]string
	// studentsCopy[0] = "New Student"

	/*
		Slices are reference type data when copied and
	*/
	// fmt.Println(students, studentsCopy)

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]   //1 slice of all elements
	// c := a[3:]  // slice from 4th element to end
	// d := a[:6]  // slice first 6 elements
	// e := a[3:6] // slice the 4th, 5th, and 6th elements

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// a := make([]int, 3, 4)
	// a = append(a, 1, 2, 3, 4, 5, 6)
	//Spread operator
	// a = append(a, []int{1, 2, 3, 4, 5, 6, 7, 8}...)
	// fmt.Println(a, len(a), cap(a))

	// var s = []int{2, 3}
	// var s1 = append(s, 4)
	// var a = append(s1, 5)
	// fmt.Println(a)
	// fmt.Println(s1)

	// var b = append(s1, 6)
	// fmt.Println(b)
	// fmt.Println(s)
	// fmt.Println(s1)
	// fmt.Println(a)
	// fmt.Println(b)

	/*
		Arrays and slices ends
	*/
	/*
		Maps and structs
	*/
	// stateMap := map[string]int{
	// 	"A": 100000,
	// 	"B": 32324343,
	// 	"C": 34232321,
	// }
	// stateMap["D"] = 5343422
	// delete(stateMap, "B")
	// _, exists := stateMap["B"]
	// fmt.Println(stateMap, exists)

	// arrayMap := map[string][]string{}
	// arrayMap["days"] = []string{"Monday", "Tuesday", "Wednesday"}
	// arrayMap["days"] = append(arrayMap["days"], "Thursday")
	// arrayMap["Month"] = make([]string, 1, 1)
	// fmt.Println(arrayMap)

	// type Person struct {
	// 	name  string
	// 	age   int
	// 	array []int
	// }

	// aPerson := Person{
	// 	name:  "John Doe",
	// 	age:   99,
	// 	array: []int{2, 31, 4, 3},
	// }
	// bPerson := aPerson
	// bPerson.array[0] = 32
	// fmt.Printf("%p, %p", &aPerson.array[0], &bPerson.array[0])

	//Anonymous struct
	// aPerson := struct{ name string }{"John Doe"}
	// fmt.Println(aPerson)

	//Embedding
	// type Company struct {
	// 	name  string
	// 	model string
	// 	price int
	// }

	// type Car struct {
	// 	Company
	// 	engine     string
	// 	fuel_type  string
	// 	horsepower int
	// }

	// aCar := Car{
	// 	engine:     "engine",
	// 	fuel_type:  "petrol",
	// 	horsepower: 600,
	// }
	// aCar.model = "C#2343"
	// aCar.name = "Company name"
	// aCar.price = 324233

	// fmt.Println(aCar)

	//TAGS
	// type Person struct {
	// 	name string `required`
	// }

	// tagType := reflect.TypeOf(Person{})
	// field, _ := tagType.FieldByName("name")
	// fmt.Println(field.Tag)

	/*
		Maps and structs ends
	*/
	/*
		Control statements
	*/

	// mapObject := map[string]int{
	// 	"a": 2,
	// 	"b": 3,
	// 	"c": 4,
	// }
	// // if item, ok := mapObject["a "]; ok {
	// // 	fmt.Println(item)
	// // } else {
	// // 	fmt.Println("Key doesn't exist")
	// // }

	// switch value := "Monday"; value {
	// case "Monday", "Thursday", "Friday":
	// 	fmt.Println("Monday")
	// 	// fallthrough
	// case "Tuesday":
	// 	fmt.Println("Tuesday")
	// case "Wednesday":
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Sunday", value)
	// }

	// value := 42
	// switch {
	// case value <= 30:
	// 	fmt.Println("First case")
	// case value <= 43:
	// 	fmt.Println("Second case")
	// default:
	// 	fmt.Println("Third case")
	// }

	/*
		Control statements ends
	*/

	/*
		Loops
	*/
	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// var numbers []int = []int{1, 2, 3, 4}
	// numberMap := map[string]int{
	// 	"a": 23,
	// 	"b": 24,
	// 	"c": 25,
	// }
	// for v, i := range numberMap {
	// 	fmt.Println(v, i)
	// }

	// Label
	// 	i := 0
	// start:
	// 	fmt.Println("Inside start")
	// 	if i < 4 {
	// 		i++
	// 		goto start
	// 	}
	/*
		Loops end
	*/

	/*
		Control flow
		Defer, panic and recovery
	*/
	// fmt.Println("Start")
	// panicker()
	// fmt.Println("End")
	/*
		Control flow ends
	*/

	/*Pointers*/
	// a := 10
	// b := &a
	// fmt.Println(a, *b)

	// type Person struc

	/*
		Functions
	*/
	// fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 9))
	// quotient, err := divide(42, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(quotient)

	// add := func(a, b int) int { return a + b }
	// fmt.Println(add(1, 2))
	// array := base_array{1, 2, 3, 2}

	// result := array.reduce(func(a, b int) int {
	// 	return a * b
	// })
	// fmt.Println(result)

	/*
		Interface
	*/

	// var writer Writer = ConsoleWriter{}
	// if _, e := writer.Write([]byte("Print this")); e != nil {
	// 	fmt.Println("There was an error printing the message")
	// }

	// fmt.Printf("%T", writer)

	/* Buffer writer */
	// var wc WriterCloser = MakeBufferWriterCloser()
	// _, err := wc.Write([]byte("Hello there, coming right out of the buffer"))
	// if err != nil {
	// 	panic(err)
	// }
	// err = wc.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// bwc := wc.(*BufferWriterCloser)
	// fmt.Printf("%T", bwc)
	/* Buffer writer ends */

	/*
		Go routines
	*/
	// waitGroup.Add(1)
	// go func() {
	// 	fmt.Println("Hello from routine")
	// 	waitGroup.Done()
	// }()
	// waitGroup.Wait()

	// fmt.Printf("Threads %v\n", runtime.GOMAXPROCS(-1))
	/*
		Go routines end
	*/

	/*Channels*/
	channnel := make(chan interface{})

	waitGroup.Add(1)

	go func(data <-chan interface{}) {
		for i := range channnel {
			fmt.Println(i, "acs")
		}
	}(channnel)

	go func(channel chan<- interface{}) {
		channnel <- 32
		channnel <- 322
		channnel <- "Here"
		close(channel)

		waitGroup.Done()
	}(channnel)
	waitGroup.Wait()
	/*Channels end*/

}

var mutex = sync.RWMutex{}
var waitGroup = sync.WaitGroup{}

/* Buffer writer */

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err = bwc.buffer.Read(v)
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, err
}

func (bwc *BufferWriterCloser) Close() (err error) {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err = fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func MakeBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

/*Buffer writer ends*/

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

//JS implementation of reduce
// type base_array []int

// func (array base_array) reduce(function func(a, b int) int) int {
// 	result := array[0]
// 	for i := 0; i < len(array)-1; i++ {
// 		result = function(result, array[i+1])
// 	}
// 	return result
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0.0, fmt.Errorf("You animal, you canNOT divide this number by 0")
// 	}
// 	return a / b, nil
// }

// func sum(numbers ...int) (result int) {
// 	for _, v := range numbers {
// 		result += v
// 	}
// 	return result
// }

// func panicker() {
// 	fmt.Println("About to panic")
// 	// defer recover()
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("Error: ", err)
// 		}
// 	}()
// 	panic("Under attack bitches")
// 	fmt.Println("End of panicker")
// }
