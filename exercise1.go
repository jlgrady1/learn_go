package main

import (
	"fmt"
	"github.com/jlgrady1/learn_go/util"
	"os"
	"strings"
)

func exercise1_1() {
	util.PrintExercise("1.1")
	echo := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s\n", echo)
	fmt.Printf("invoked by %s\n", os.Args[0])
}

func exercise1_2() {
	util.PrintExercise("1.2")
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d - %s\n", i, arg)
	}
}

func exercise1_3() {
	util.PrintExercise("1.3")

}

func main() {
	exercise1_1()
	exercise1_2()
}
