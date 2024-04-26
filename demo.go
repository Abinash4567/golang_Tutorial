package main

import (
	"errors"
	"fmt"
)

func returned(age int) (int, error) {
    if age > 50{
        return 0, errors.New("cannot evaluate");
    }
    return 1, nil
}

func main(){
    x, y := returned(89)
    fmt.Printf("%d %d", x, y);
}