package main

import (
	"fmt"
	"os"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// for initialization, condition; post {
// zero or more statements
//}

/*
- the initialization statement is optional
- condition is a boolean
- post stmt is executed after the body of the loop then the condition
  is evaluated again
- Any of the parts may be omitted. If there is no initialization and
  no post the ';' may also be omitted
*/
