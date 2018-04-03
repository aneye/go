package main

import (
	"fmt"
	"github.com/aneye/go/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("Hello, world"));
	for i := 0; i< 5; i++ {
		fmt.Printf("hello ");
	}
	i := 2;
	for i < 4 {
		fmt.Printf("hello 2");
		i = i + 1 ;
	}
	for {
		fmt.Printf("Bye 3");
		i++;
		if i > 10 {
			break;
		}
	}
}

