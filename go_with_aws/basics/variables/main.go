package main

import "fmt"

func main() {
	var myName string = "Shivani"
	age := 21
	myFloat := 78.1

	fmt.Printf("Hello, I am %s, age is %d, float is %f\n", myName, age, myFloat)

	// go has concept of zero-value. we can declare a var, but haven't mentioned value, we get default vals - false (bool) 0 (int)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "Prime"

	fmt.Printf("my friends name is %s my bool is %t and my other int is %d\n", myFriendsName, myBool, myOtherInt)
}
