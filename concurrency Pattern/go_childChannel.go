package main

import (
	"fmt"
	"time"
)

func calling(ch <-chan bool){
	for{
		select{
		case <-ch:
			return;
		default:
			fmt.Println("Working!!!");
		}
	}
}

func main() {
	ch := make(chan bool);
	go calling(ch);
	time.Sleep(time.Second * 4);
	// ch<-true;
	close(ch);
}
