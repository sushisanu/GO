package main

import "fmt"

func main() {
	msg := "some message";
	var msgPointer *string = &msg;

	fmt.Println(msg)
	fmt.Println(*msgPointer)

	changeMessage(&msg, "new maseeage 1")
	fmt.Println(msg)

	changeMessage(msgPointer, "new maseeage 2")
	fmt.Println(msg)

	changeMessage(msgPointer, "new maseeage 3")
	fmt.Println(*msgPointer)

}

func changeMessage(aPointer *string, newMessage string) {
    *aPointer = newMessage 
}