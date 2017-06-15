package main

import (
    "os"
    "os/signal"
    "syscall"
    "./automaton"
)

func main() {
    sig := make(chan os.Signal, 5)
    // register signal
    signal.Notify(sig, syscall.SIGUSR1)

    automaton.Get_msgs(sig, "eva")
}
