package main

import (
	
	"fmt"
	"encoding/base64"
	"encoding/hex"
)

var x int
var y int
var a int
var b int
var c int
var d int

var option int
var obj string
var message string

func xorEncrypt(text string, key byte) string {
	// turn the text to byte
	data := []byte(text)
	// find the length of the byte
	result := make([]byte, len(data))
	// loop and xor them char by char
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key
	}
	return string(result)
}

func main() {

	//Lab 1:
	
	// initializing the value of x to 4
	x  = 4
	// adding 5
	x += 5
	// subtracting 3
	x -= 3
	// multiplying by 2
	x *= 2
	// dividing by 3
	x /= 3
	// modulus by 2
	x %= 2
	fmt.Println(x)

	//Lab 2:

	y = 1

	since y is less than 5 and less than 3, the result is true
	fmt.Println(y < 5 && y < 3)
	since y is less than 5 but more than 0, the result is true
	fmt.Println(y < 5 || y < 0)
	since y is not more than 5 and more than 3, but since we use the ! operator, it inverts the result making it true
	fmt.Println(!(y > 5 && y > 3))

	//Lab 3:

	x = 9
	y = 10

	c  = ^4
	d = ^5
	
	%b it used to convert the int to binry
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("c = %b\n", c)
	fmt.Printf("d = %b\n", d)


	fmt.Printf("x & y is %b\n", x & y)


	fmt.Printf("x | y is %b\n", x | y)


	fmt.Printf("x ^ y is %b\n", x ^ y)

	fmt.Printf("NOT c is %b\n", c)
	fmt.Printf("NOT d is %b\n", d)

	fmt.Printf("x << 2 is %b\n", x << 2)
	fmt.Printf("y >> 2 is %b\n", y >> 2)

	//Lab 4:

	for {

		fmt.Println("===== Mini Calculator =====")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Modulus")
		fmt.Println("6. Exit")
	
		fmt.Print("Choose: ")
		fmt.Scan(&option)
	
		if option == 6 {
			fmt.Println("Exiting...")
			break
		}
	
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		
		fmt.Print("Enter b: ")
		fmt.Scan(&b)
	
		switch option {
		case 1:
			fmt.Println(a + b)
		case 2:
			fmt.Println(a - b)
		case 3:
			fmt.Println(a * b)
		case 4: 
			fmt.Println(a / b)
		case 5:
			fmt.Println(a % b)
		default:
			fmt.Println("Error.")
		}
		
	}

	//Lab 5

	fmt.Println("===== Binary, Hex and Base64 Encoding")
	fmt.Print("Enter a string: ")
	fmt.Scan(&obj)
	data := []byte(obj)

	fmt.Println("Your string is: ", obj)
	fmt.Printf("Binary version: %b\n", data)
	fmt.Println("Hex: ", hex.EncodeToString(data))
	fmt.Println("Base64: ", base64.StdEncoding.EncodeToString(data))

	// Lab 6

	fmt.Println("Enter your message: ")
	fmt.Scan(&message)

	key := byte('K')

	cipher := xorEncrypt(message, key)
	plain := xorEncrypt(cipher, key)

	fmt.Println("Your message: ", message)
	fmt.Printf("Cipher text: %x\n", cipher)
	fmt.Println("Decrypted: ", plain)



}