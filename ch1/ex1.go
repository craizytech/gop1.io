package main

import (
	"fmt"
	"os"
	"strings"
)

func ex1() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
