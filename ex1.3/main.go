package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func join(strs []string, sep string) string {
	result := ""
	for i, strs := range strs {
		result += strs
		if i+1 < len(strs) {
			result += sep
		}
	}
	return result
}

func main() {
	// implemented Join
	myStartTime := time.Now()
	myJoinResult := join(os.Args, " ")
	myEndTime := time.Now()

	// strings Join
	startTime := time.Now()
	joinResult := strings.Join(os.Args, " ")
	endTime := time.Now()

	fmt.Println(myJoinResult, myEndTime.Sub(myStartTime))
	fmt.Println(joinResult, endTime.Sub(startTime))
}
