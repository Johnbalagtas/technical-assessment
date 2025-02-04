package utils

// function numIsPrime to check if the number is a prime number
// @return boolean
// @param num int
func NumIsPrime(num int) bool {
	if num < 2 {
		return false
	}

	// Loop through the numbers from 2 to the square root of the number
	// If the number is divisible by any of the numbers, it is not a prime number
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	// If the number is not divisible by any of the numbers, it is a prime number
	return true
}
