package main

import (
	"bufio"
	"fmt"
	"os"
)


/*
$ go run main.go 
first line
first line
second line
second line
...
...
^Csignal: interrupt
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
