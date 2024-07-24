package main

import (
	"fmt"
	"time"
)

func callfunction(m string){
	fmt.Println(m);
}

func main() {
	go callfunction("Child Thread");
	fmt.Println("Waiting!!!")
	time.Sleep(time.Second * 3);
	fmt.Println("Main Thread")
}
