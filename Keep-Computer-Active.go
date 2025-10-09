package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	// Print version number, if requested.
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-version" {
			fmt.Println("Version number: 1.3")
			os.Exit(0)
		}
	}

	for {
		// Linux gets angrryyyy when you try to use F13.
		if runtime.GOOS == "linux" {
			robotgo.KeyTap("scrolllock")
			robotgo.KeyTap("scrolllock")
		} else {
			robotgo.KeyTap("f13")
		}

		now := time.Now()
		next3Minutes := now.Truncate(3 * time.Minute).Add(3 * time.Minute)

		// After using other languages'
		fmt.Println("Continuing to keep computer active. See you in 3 minutes! Current time is", time.Now().Local().Format("03:04:05 PM."))

		// Wait for 3 minutes.
		time.Sleep(time.Until(next3Minutes))
	}
}
