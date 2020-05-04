package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	var msg = ""
	var num int = 1
	msg = "test"+strconv.Itoa(num)+": " + t.String()
	fmt.Println(msg)
	num++
	msg = "test"+ strconv.Itoa(num) + ": " + t.String()
	fmt.Println(msg)
	num++
	msg = "test"+string(num)+": " + t.String()
	fmt.Println(msg)
	num++
	msg = "test"+string(num)+": " + t.String()
	fmt.Println(msg)
	num++
	
}