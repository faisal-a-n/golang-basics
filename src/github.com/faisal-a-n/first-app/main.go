package main

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
	// fmt.Println(a, len(a), cap(a))

	/*
		Arrays and slices ends
	*/
}
