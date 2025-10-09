package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
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
	// Handle exiting more gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		// Linux gets angrryyyy when you try to use F13.
		if runtime.GOOS == "linux" {
			robotgo.KeyTap("scrolllock")
			robotgo.KeyTap("scrolllock")
		} else {
			robotgo.KeyTap("f13")
		}

		fmt.Println("Continuing to keep computer active. See you in 3 minutes! Current time is", time.Now().Local().Format("03:04:05 PM."))

		// Wait for 3 minutes. This has drifting, but idfk how to fix it, okay?
		time.Sleep(3 * time.Minute)

		ctx.Done()

	}
}
