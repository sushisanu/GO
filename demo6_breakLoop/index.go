package main

import "fmt"

func main() {
	fnFor();
	fnWhileUsingBreak();
}

func fnFor() {
	for index := 0; index <=  10; index++ {
		fmt.Printf("For index %d\n", index);
	}
}

func fnWhileUsingBreak() {
	index := 0;
	for true {
		if index > 5{
			break
		} 
			
		fmt.Printf("While break index %d\n", index);
		index++;

	}
}