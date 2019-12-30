# Simple Pomodoro Timer

This pomodoro timer has a very simple feature set based on my current needs.
Most use cases will be either:

```bash
./pomodoro.exe # for running a default timer for 25 minutes
```

or

```bash
./pomodoro.exe --break # for taking a 5 minute break after a pomodoro
./pomodoro.exe --break-long # for taking a 15 minute break after 4 pomodoros
```

Another common use case would be to add the `--message` options tp specify the task you are working on during that pomodoro.

---

All current available options can be shown by invoking

```bash
./pomodoro.exe --help

Usage of pomodoro.exe
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
