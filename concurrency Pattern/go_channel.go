package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string);
	anotherChannel := make(chan string);
	
	go func(){
		myChannel <- "data from first Process";
	}();

	go func(){
		anotherChannel <- "data from second Process";
	}();

	select {
	case messageFromFirst := <- myChannel:
		fmt.Println(messageFromFirst);
	case messageFromSecond := <- anotherChannel:
		fmt.Println(messageFromSecond);
	}

	fmt.Println("Done!!!");
}
