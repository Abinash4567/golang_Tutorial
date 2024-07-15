package main

import "fmt"

func main(){
    var username string = "John Doe";
    var age int = 30;
    fmt.Println("Hello my name is", username, "and I am", age, "years old.");
    fmt.Printf("Hello my name is %s and I am %d years old.\n", username, age);
    var sentence string  = fmt.Sprintf("Hello my name is %s and I am %d years old.", username, age);
    fmt.Println(sentence);
}