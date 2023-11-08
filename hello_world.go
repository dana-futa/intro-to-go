package main // required for a standalone executable

import "fmt" // fmt implements formatted IO

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
	/* is there a ternary operator in Go?? */

	// examples of variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// variable type is inferred
	/* can variable type change?? */
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
