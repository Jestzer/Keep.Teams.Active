package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/fatih/color"
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

	// Setup for better Ctrl+C messaging, specifically. This is a channel to receive OS signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	redText := color.New(color.FgRed).SprintFunc()

	go func() {

		<-signalChan

		fmt.Println(redText("\nExiting from user input."))
		os.Exit(0)
	}()

	for {
		// Linux gets angrryyyy when you try to use F13.
		if runtime.GOOS == "linux" {
			robotgo.KeyTap("scrolllock")
			robotgo.KeyTap("scrolllock")
		} else {
			robotgo.KeyTap("f13")
		}

		fmt.Println("Continuing to keep computer active. See you in roughly 3 minutes! The current time is", time.Now().Local().Format("03:04:05 PM."))

		// Wait for 3 minutes.
		// It will seem like the time is slowly drifting. That's because this program and the for loop don't run instantly...
		// ... and the 3 minute timer doesn't take this into account. I could make it /seemingly/ run without any drifting...
		// ... but there's really no point in adding all that code for something that really doesn't make any difference.
		time.Sleep(3 * time.Minute)
	}

}
