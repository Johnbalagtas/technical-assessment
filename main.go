package main

// import fmt package for input/output
import (
	"fmt"

	"technical.com/exercise/utils"
)

// main function (entry point of the program)
func main() {

	// variable input type: int (storing the input data from a user)
	var input int

	// for loop infinite loop until the input == 0; break
	for {
		// print out the message to the user (asking for a number)
		fmt.Print("Enter a number: ")

		// scan the input from the user and store it in the input variable (address)
		fmt.Scan(&input)

		// condition check if the input is a prime number with the help of the function numIsPrime
		// the function numIsPrime will return a boolean value

		if utils.NumIsPrime(input) {
			// if the condition is true print out the message to the user that the input is a prime number
			fmt.Println(input, "IS A PRIME NUMBER")
			factorial := utils.Factorial(input)
			fmt.Printf("Factorial of %d is %d\n", input, factorial)
		} else {
			// if the condition is false print out the message to the user that the input is not a prime number
			fmt.Println(input, "IS NOT A PRIME NUMBER")
		}

		// condition check if the input is 0
		// if the condition is true print out the message to the user that the program has ended
		if input == 0 {
			fmt.Print("Ended")
			break
		}
	}

}
