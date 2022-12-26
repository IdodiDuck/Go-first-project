// Basic calculator

//Importations of fmt and package main
package main
import "fmt"

func add(num1 int16, num2 int16) {
	//This function is adding  the first number to the second number or the oppisite
	//Parameters:
	//-----------
	//param num1: type int16
	//		The first number the function recieve, the number which is being added with the second number argument
	//		Note: type of int16 --> 16 bytes, used for small numbers in range of around 100-200 numbers, not for larger numbers which requires more bytes.

	//Param num2: type int16
	//		The second number which is the adder of the first number
	//		Note: type of int16 --> 16 bytes like above.

	fmt.Println(num1 + num2)
}

func subtract(num1 int16, num2 int16) {
	//This function is subtracting the first number in the second number
	//Parameters:
	//-----------
	//param num1: type int16
	//		The first number the function recieve, the number which is being subtracted in the second number argument
	//		Note: type of int16 --> 16 bytes, used for small numbers in range of around 100-200 numbers, not for larger numbers which requires more bytes.

	//Param num2: type int16
	//		The second number which is the subtractor of the first number
	//		Note: type of int16 --> 16 bytes like above.

	fmt.Println(num1 - num2)
}

func multiply(num1 int16, num2 int16) {
	//This function is multiplying the first number in the second number
	//Parameters:
	//-----------
	//param num1: type int16
	//		The first number the function recieve, the number which is being multiplied in the second number argument
	//		Note: type of int16 --> 16 bytes, used for small numbers in range of around 100-200 numbers, not for larger numbers which requires more bytes.

	//Param num2: type int16
	//		The second number which is the multiplier of the first number
	//		Note: type of int16 --> 16 bytes like above.

	fmt.Println(num1 * num2)
}

func divide(num1 int16, num2 int16) {
	//This function is dividing the first number in the second number
	//Parameters:
	//-----------
	//param num1: type int16
	//		The first number the function recieve, the number which is being divided in the second number argument
	//		Note: type of int16 --> 16 bytes, used for small numbers in range of around 100-200 numbers, not for larger numbers which requires more bytes.

	//Param num2: type int16
	//		The second number which is the divider of the first number
	//		Note: type of int16 --> 16 bytes like above.
	
	full_result := float64(num1) / float64(num2)
	fmt.Println(full_result)
}

func main() { 
	// The main function of the program, doesn't recieve any argument unlike the other functions
	//Call the mathematical functions in the main function
	
}