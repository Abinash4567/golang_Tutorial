package main

import "fmt"

func main() {



	var name string = "Abc";
	fmt.Println(name);

	var logggedin bool = false;
	fmt.Println(logggedin);

	var value float64 = 12.39483;
	fmt.Println(value);

	var bb uint32 = 124;
	fmt.Println(bb);

	//implicit
	var website = "Hello lexical";
	fmt.Println(website);

	numberofUser := 123;
	fmt.Println(numberofUser);

	fmt.Println(cores);
	fmt.Printf("It's type is: %T",cores);
	
	
	

	roll_num, sentence := 12, "My name is Abinash"
	var registration int = 20214540
	fmt.Printf("%s and my roll number is %d.\nMy registration number is %d.\n", sentence, roll_num, registration)

	if roll_num > 0 {
		fmt.Println("Registration Number is Positive.")
	} else {
		fmt.Println("Registration Number is Negative.")
	}
}
