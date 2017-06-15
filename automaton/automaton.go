package automaton

import (
    "fmt"
    "os"
    "time"
    "syscall"
    "../messagetypes"
)

func Get_msgs (sig <-chan os.Signal, name string) {
    var msg_name string

    // set robot options
    fmt.Printf("RobotOption %d %d\n", messagetypes.SIGNAL, syscall.SIGUSR1)

    for {
        // non-blocking signal checking
        select {
        case <-sig:
            // get message name
            fmt.Scanf("%s", &msg_name)
            fmt.Println(msg_name)

            // robot response
            switch msg_name {
            case "Initialize":
                var init int
                fmt.Scanf("%d", &init)

                if (init == 1) {
                    fmt.Printf("Name %s\n", name)
                    fmt.Println("Colour dede11 de5500")
                }

            case "GAME_STARTS":
                fmt.Printf("Sweep 6 %f %f %f\n", 1.0, 2.0, 3.0)
                fmt.Printf("Accelerate %f\n", 0.4)
            }

        default:
            fmt.Println("Print no msg")
        }

        time.Sleep(3000 * time.Millisecond)
        fmt.Println("Print return")
    }
}
