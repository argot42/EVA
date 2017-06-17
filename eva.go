package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "./automaton"
    "./messagetypes"
)

func main() {
    // set robot options 
    // tell rtb to use non-blocking communication
    //fmt.Printf("RobotOption %d %d\n", messagetypes.USE_NON_BLOCKING, 1)
    // tell rtb to use signals
    fmt.Printf("RobotOption %d %d\n", messagetypes.SIGNAL, syscall.SIGUSR1)

    // register signal
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGUSR1)


    var end bool = false
    for !end {
        <-sig
        end = automaton.Get_msg()
        fmt.Printf("Print ---------\n")
    }
}
