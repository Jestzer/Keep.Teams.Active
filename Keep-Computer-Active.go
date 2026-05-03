package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/go-vgo/robotgo"
)

func main() {

	// Print version number, if requested.
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-version" || arg == "--version" || arg == "-v" {
			fmt.Println("Version number: 1.5")
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

	// On macOS, synthesized F13 events don't reset HIDIdleTime, so spawn caffeinate
	// to create an IOPMAssertion. -w makes it exit when this process exits.
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("caffeinate", "-d", "-i", "-w", strconv.Itoa(os.Getpid()))
		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, redText("Warning: failed to start caffeinate: "+err.Error()))
		}
	}

	ticker := time.NewTicker(3 * time.Minute)
	defer ticker.Stop()

	for {
		switch runtime.GOOS {
		case "linux":
			// Linux gets angrryyyy when you try to use F13.
			robotgo.KeyTap("scrolllock")
			robotgo.KeyTap("scrolllock")
		case "darwin":
			// caffeinate is doing the work; no per-tick keypress needed.
		default:
			robotgo.KeyTap("f13")
		}

		fmt.Println("Continuing to keep computer active. See you in 3 minutes! The current time is", time.Now().Local().Format("03:04:05 PM."))

		<-ticker.C
	}

}
