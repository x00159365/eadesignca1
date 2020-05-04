package main

import (
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/mediocregopher/radix/v3" // redis api for pub-sub
)

func main() {

	// MG leave as string
	// get the door number from command line arg 1
	var newsSource = os.Args[1]
	//if err != nil {
	//	panic(err)
	//}

	// get the max number of seconds between entries for the random generator
	maxSeconds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	// connect to redis specified by url in the third argument
	conn, err := Dial("tcp", os.Args[3])
	if err != nil {
		panic(err)
	}

	var num int = 1

	for {
		// get the content of message
		var t = time.Now()
		var msgContent = os.Args[4] + ":update time" + strconv.Itoa(num) + "::" + t.Format("2006-01-02 15:04:05")
		num++
		log.Printf(newsSource + "[INFO]: Publishing to channel " + newsSource + " ::msg:" + msgContent)
		// the event contains a dummy string "1" - it is the count of events that matters, not the content
		conn.Do(Cmd(nil, "PUBLISH", newsSource, msgContent))
		time.Sleep(time.Duration(maxSeconds) * time.Second)
	}

}
