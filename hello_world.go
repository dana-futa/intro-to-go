package main // required for a standalone executable

import (
	"fmt"
	"math"
	"time"
) // fmt implements formatted IO

// can also write the imports like this, is there a performance improvement for one over the other?
// import "fmt"

// global variables
var g int
var (
	h bool
	i float32
	j string
)

// main.main() is the first function that is run when this program executes
func main() {
	fmt.Println("Hello world!")
	fmt.Println("こんにちは世界！")

	// examples of values
	fmt.Println("go" + "lang")
	fmt.Println("2+3 =", 2+3)
	fmt.Println("9.0/4.0 =", 9.0/4.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	/* QUESTION: is there a ternary operator in Go??
	--> https://gobyexample.com/if-else
	"There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions."
	*/

	// examples of variables
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// variable type is inferred
	/* QUESTION: can variable type change?? */
	var d = true
	fmt.Println(d)
	// d = "false" /* looks like the type can't change, nice */

	/*
		"Variables declared without a corresponding initialization are zero-valued.
		For example, the zero value for an int is 0."
			Very interesting. Wonder why null wasn't chosen.
	*/
	var e int
	fmt.Println(e)
	/* double checking if Go behaves like JavaScript. In JS 0 == false resolves to true. Nice type checking in Go. */
	// fmt.Println(e == false)

	// := syntax is shorthand for declaring and initializing a variable --> var f string = "chicken sandwich"
	/* did the creaters of Go borrow := from PostgreSQL? */
	f := "chicken sandwich"
	fmt.Println(f)

	// intializing global variables
	g = 7
	h, i = false, 3.14
	j = "I am a string"
	fmt.Println(g, h, i, j)

	// examples of constants
	const k = 5000
	fmt.Println(math.Sin(k))
	const l = 3e20 / k
	fmt.Println(l)
	fmt.Println(int64(l))

	// examples of types
	const m int32 = 12        // 32-bit integer
	const n float32 = 20.5    // 32-bit float
	var o complex128 = 1 + 4i // 128-bit complex number, so is "i" a reserved value? I have a variable named "i" with the value of 3.14
	var p uint16 = 14         // 16-bit unsigned integer
	fmt.Println(m, n, o, p)
	fmt.Println("what is the value of 'o'", o)
	var q = i
	fmt.Println("will q print the value of the variable 'i' or the complex number 'i'", q)
	// const r complex128 = i    // this causes type error
	const r complex128 = 1i // this does not cause type error
	fmt.Println(r)

	/*
		To use complex number "i" the type of the variable or constant must be complexX.
		That's good news. For loops can still use "i".
	*/

	s := 42                                       // int
	t := 3.14                                     // float64
	u, v := true, false                           // bool
	w := "Go is cool!"                            // string
	fmt.Printf("%T %T %T %T %T\n", s, t, u, v, w) // log the type of a variable with %T

	// examples of for loops
	x := 1
	for x <= 3 {
		fmt.Println(x)
		x = x + 1
	}

	for y := 9; y <= 12; y++ {
		fmt.Println(y)
	}

	for {
		fmt.Println("in a loop that will break immediately")
		break
	}

	for z := 0; z <= 5; z++ {
		if z%2 == 0 {
			continue
		}
		fmt.Println(z)
	}

	// examples of if/else statements
	if 5%2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}

	if 12%4 == 0 {
		fmt.Println("12 is divisible by 4")
	}

	if 5%2 == 0 || 6%2 == 0 {
		fmt.Println("either 5 or 6 are even")
	}

	// you can have a statement before a condition
	/* this seems pretty handy to not have to declare extra variables outside of this logic
	especially when creating such a variable would make the code more readable, but the variable
	wouldn't be used anywhere else */
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// num = num + 1; // can't access num outside the scope of the conditions

	// examples of switch statements
	someNum := 8
	fmt.Print("Write ", someNum, " as ")
	switch someNum {
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	}

	fmt.Println("time.Now().Weekday() outputs:", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/* QUESTION: is it practical or is there any kind of performance gain from using a switch without an expression
	as an if/else statement? */
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// this is a "type switch"
	whatAmI := func(i interface{}) {
		switch checkType := i.(type) { // statement before condition is used here, nice
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", checkType)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hi")

	// examples of arrays
	var animalsArr = [4]string{"cat", "dog", "chicken", "frog"}
	for iter := 0; iter < len(animalsArr); iter++ {
		fmt.Println(iter, animalsArr[iter])
	}

	// for loop with range
	var birdsArr = [...]string{"sparrow", "robin", "turkey", "ostrich"} // ... NOT JS spread syntax, this lets compiler count the array length
	for index, option := range birdsArr {
		fmt.Println(index, option)
	}

	// examples of slices --> dynamic arrays
	/*
		slice length = "current length"
		slice capacity = "max length the slice can grow to"
	*/

}

/*
    NOTES
go build hello_world.go		compiles source code and dependencies into executable binary
./hello_world				./ followed by binary file name executes the binary
go run hello_world.go		combines the above commands, but doesn't create a binary

    REFERENCES
https://go.dev/doc/tutorial/getting-started
https://gobyexample.com/hello-world
https://fireship.io/lessons/learn-go-in-100-lines/

*/
