package main

import "fmt"

func main() {
	fnFor();
	fnWhile();
}

func fnFor() {
	for index := 0; index <=  10; index++ {
		fmt.Printf("For index %d\n", index);
	}
}

func fnWhile() {
	index := 0;
	for index < 5 {
		index++;
		fmt.Printf("While index %d\n", index);
	}
}