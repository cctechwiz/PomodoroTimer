package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	fh "github.com/cctechwiz/utils/flagHelpers"
)

var (
	duration    time.Duration
	message     string
	shortBreak  bool
	longBreak   bool
	logFilePath string
	disableLog  bool
)

func init() {
	flag.DurationVar(&duration, "duration", (25 * time.Minute), "set timer duration")
	flag.StringVar(&message, "message", "", "message for this timer")
	flag.BoolVar(&shortBreak, "break", false, "take a short break (default 5m0s)")
	flag.BoolVar(&longBreak, "break-long", false, "take a long break (default 15m0s)")
	flag.StringVar(&logFilePath, "log", "./completedPomodoros.csv", "location to save logs")
	flag.BoolVar(&disableLog, "log-disable", false, "disable logging")

	flag.Parse()

	if duration < (1 * time.Second) {
		fmt.Println("Must specify a minimum of 1 second.")
		return
	}

	if shortBreak && longBreak {
		fmt.Println("Can only specify one type of break.")
		return
	}

	if shortBreak {
		if !fh.IsFlagPassed("duration") {
			duration = 5 * time.Minute
		}
		message += " (short break)"
	} else if longBreak {
		if !fh.IsFlagPassed("duration") {
			duration = 15 * time.Minute
		}
		message += " (long break)"
	}
}

func main() {
	fmt.Printf("Starting %v timer", duration)
	if message != "" {
		fmt.Printf(": %v", message)
	}
	fmt.Println()

	startTime := time.Now().Format(time.RFC850)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(duration)
		done <- true
	}()

	timeRemaining := duration
Loop:
	for {
		select {
		case <-done:
			fmt.Println("\nDone!")
			break Loop
		case <-ticker.C:
			//TODO: Look into replacing this hack with a TUI or something for nicer output
			fmt.Printf("\r                 ")
			fmt.Printf("\r%v", timeRemaining)
			timeRemaining = timeRemaining - (1 * time.Second)
		}
	}

	stopTime := time.Now().Format(time.RFC850)

	if !disableLog {

		logHeader := []string{"startTime", "stopTime", "duration", "message"}
		logRecord := []string{startTime, stopTime, duration.String(), message}

		f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()

		fi, err := f.Stat()
		check(err)

		w := csv.NewWriter(f)

		if fi.Size() == 0 {
			werr := w.Write(logHeader)
			check(werr)
		}

		werr := w.Write(logRecord)
		check(werr)
		w.Flush()
	}
}

//TODO: Move to utils package once it is all flushed out and useful
func check(err error) {
	if err != nil {
		fmt.Println(err)
		//TODO: add enum arument that decides what program does if there is an error (e.g. abort, resume, etc.)
	}
}
