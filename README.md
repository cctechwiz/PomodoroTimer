# Simple Pomodoro Timer

This pomodoro timer has a very simple feature set based on my current needs.

### Installing

[Golang](https://golang.org/) is require to run / install this program.

This program can be run direclty with `go run PomodoroTimer.go` after cloning or downloading the code.

It can also be installed with
```bash
go get github.com/cctechwiz/PomodoroTimer
```

### Usage

Most use cases will be either:

```bash
go run PomodoroTimer.go # for running a default timer for 25 minutes
```

or

```bash
go run PomodoroTimer.go --break # for taking a 5 minute break after a pomodoro
go run PomodoroTimer.go --break-long # for taking a 15 minute break after 4 pomodoros
```

Another common use case would be to add the `--message` options to specify the task you are working on during that pomodoro, or how you spent the break.

---

All current available options can be shown by invoking

```bash
go run PomodoroTimer.go --help

Usage of PomodoroTimer
  -break
        take a short break (default 5m0s)
  -break-long
        take a long break (default 15m0s)
  -duration duration
        set timer duration (default 25m0s)
  -log string
        location to save logs (default "./completedPomodoros.csv")
  -log-disable
        disable logging
  -message string
        message for this timer
```
