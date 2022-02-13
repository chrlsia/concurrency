package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		execute the function twice , sequentially
		with arguments "chris" and "tasia"
	*/
	count("chris")
	count("tasia")
	/*
		execute the function:
			first time as goroutine, argument "chris"
			second time not as goroutine, argument "tasia"
	*/
}

/*
	create a function which:
		get a string as parameter
		iterates 3 times
			in each iteration prints the iteration number,string
		sleeps for 1 sec
*/
func count(str string) {
	for i := 1; i < 4; i++ {
		fmt.Printf("%v %v\t", i, str)
		time.Sleep(time.Second * 1)
	}
	fmt.Println()
}
