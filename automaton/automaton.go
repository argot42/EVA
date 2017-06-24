package automaton

import (
    "fmt"
    "math"
    "../configuration"
    "../messagetypes"
    "../structures"
)

// Global gameinfo/robotinfo
var spotted_robot structures.Robot
var spotted_object structures.Object//{spotted: false}
var g_status structures.Game_status
var r_status structures.Robot_status

func Get_msg () bool{
    var msg_name string

    // get msg name
    _,err := fmt.Scan(&msg_name)
    if (err != nil) { fmt.Printf("Print Error getting msg name: %s\n", err) }

    switch msg_name {
    case "Initialize":
        var init int
        _,err = fmt.Scan(&init)
        if (err != nil) { fmt.Printf("Print Error getting init number: %s\n", err) }

        // set robot name and colors
        if (init == 1) {
            fmt.Printf("Name %s\n", configuration.NAME)
            fmt.Printf("Colour %s %s\n", configuration.COLOR, configuration.X)
        }
        fmt.Printf("Print initializing\n")

    case "GameOption":
        var option_n int
        var value float32

        _,err = fmt.Scan(&option_n)
        if (err != nil) { fmt.Printf("Print Error getting optionnr: %s\n", err) }
        _,err = fmt.Scan(&value)
        if (err != nil) { fmt.Printf("Print Error getting value: %s\n", err) }

        g_status.Game_options[option_n] = value

    case "GameStarts":
        fmt.Printf("Print I'm ready\n")
        fmt.Printf("Sweep %d %f %f %f\n", 6, math.Pi/2.0, -math.Pi/2.0, math.Pi/2.0)
        fmt.Printf("Accelerate %f\n", 10.0)

    case "GameFinishes":
        fmt.Printf("Print Goodbye\n")

    case "ExitRobot":
        return true

    case "RobotsLeft":
        fmt.Printf("Print robots left\n")
        _,err = fmt.Scan(&g_status.Robots_left)
        if (err != nil) { fmt.Printf("Print Error getting last robots left: %s\n", err) }

    case "Radar":
        /* This message gives information from the radar each turn. The radar angle is relative to the robot front;
         * it is given in radians */
         fmt.Printf("Print observing\n")

        var observed_obj_type int
        var test string 

        _,err = fmt.Scan(&spotted_object.Pos.Distance)
        if (err != nil) { fmt.Printf("Print Error getting distance: %s\n", err) }
        _,err = fmt.Scan(&spotted_object.Pos.Angle)
        if (err != nil) { fmt.Printf("Print Error getting angle: %s\n", err) }
        _,err = fmt.Scan(&observed_obj_type)
        if (err != nil) { fmt.Printf("Print Error getting observed object type: %s\n", err) }

        _,err = fmt.Scan(&test)
        if (err != nil) { fmt.Printf("Print error getting test: %s\n", err) }

        fmt.Printf("Print %f - %f - %d - %s\n", spotted_object.Pos.Distance, spotted_object.Pos.Angle, observed_obj_type, test)

        switch observed_obj_type {
        case messagetypes.ROBOT:
            //resolve_battle()
            fmt.Printf("Print target acquired\n")

        case messagetypes.WALL:
            wallfacer()

        case messagetypes.COOKIE:
            cookie_spotted()

        case messagetypes.SHOT:
            bullet_spotted()

        case messagetypes.MINE:
            mine_spotted()

        case messagetypes.NOOBJECT:
            nothing_spotted()
        }

    case "Info":
        /* The Info message does always follow the Radar message. It gvies more general information on the state of the robot.
         * The time is the game-time elapsed since the start of the game. This is not necessarily the same as the real time elapsed,
         * due to time scale and maxe timestep. */
         fmt.Printf("Print getting info\n")

        _,err = fmt.Scan(&g_status.Elapsed_time)
        if (err != nil) { fmt.Printf("Print Error getting time: %s\n", err) }
        _,err = fmt.Scan(&r_status.Speed)
        if (err != nil) { fmt.Printf("Print Error getting speed: %s\n", err) }
        _,err = fmt.Scan(&r_status.Cannon_angle)
        if (err != nil) { fmt.Printf("Print Error getting cannon angle: %s\n", err) }

    case "Coordinates":
        /* Tells you the current robot position. It is only sent if the option 'Send robot coordinates is 1 or 2'. If it is 1
         * the coordinates are snet relative the starting position, which has the effect that the robot doesn't know where it is starting,
         * but only where it has moved since. */
         fmt.Printf("Print getting coordinates\n")

        _,err = fmt.Scan(&r_status.Pos.X)
        if (err != nil) { fmt.Printf("Print Error getting x axis value: %s\n", err) }
        _,err = fmt.Scan(&r_status.Pos.Y)
        if (err != nil) { fmt.Printf("Print Error getting y axis value: %s\n", err) }
        _,err = fmt.Scan(&r_status.Pos.Angle)
        if (err != nil) { fmt.Printf("Print Error getting angle value: %s\n", err) }

    case "RobotInfo":
        /* If you detect a robot with your radar, this message will follow, giving some information on the robot. The opponents energy level
         * will be given in the same manner as your own energy. The second argument is only interesting in team-mode, 1 means a
         * teammate and 0 an enemy */
         fmt.Printf("Print getting robot info\n")

        _,err = fmt.Scan(&spotted_robot.Energy)
        if (err != nil) { fmt.Printf("Print Error getting energy level: %s\n", err) }
        _,err = fmt.Scan(&spotted_robot.Teammate)
        if (err != nil) { fmt.Printf("Print Error getting teammate/enemy flag: %s\n", err) }

        if (!spotted_robot.Teammate) {
            resolve_battle()
        }

    case "RotationReached":
        /* When the robot option SEND_ROTATION_REACHED is set appropriately, this message is sent when a rotation 
         * (with RotateTo or RotateAmount) has finished or the direction has changed (when sweeping). 
         * The argument corresponds to 'what to rotate' in e.g. Rotate. */
         fmt.Printf("Print rotation reached\n")

        _,err = fmt.Scan(&r_status.Rotation_reached)
        if (err != nil) { fmt.Printf("Print Error getting rotation reached: %s\n", err) }

     case "Energy":
         /* The end of each round the robot will get to know its energy level. It will not, however, get the exact energy, instead it is a
          * discretized into a number of "energy levels" */
          fmt.Printf("Print getting energy\n")

        _,err = fmt.Scan(&r_status.Energy)
        if (err != nil) { fmt.Printf("Print Error getting own energy level: %s\n", err) }

    case "Collision":
        /* When a robot hits (or is hit by) something it gets this message. You get the angle from where the collision ocurred
         * (the angle relative the robot) and the type of object hitting you, but not how severe the collision was. This can
         * however, be determined indirectly (approximately) by the loss of energy. */
         fmt.Printf("Print identifying collition\n")

        var obj_type int
        var angle_relative_robot float32

        _,err = fmt.Scan(&obj_type)
        if (err != nil) { fmt.Printf("Print Error getting colliding object type: %s\n", err) }
        _,err = fmt.Scan(&angle_relative_robot)
        if (err != nil) { fmt.Printf("Print Error getting angle relative to robot: %s\n", err) }

        switch obj_type {
        case messagetypes.ROBOT:
            bump_into_robot()

        case messagetypes.SHOT:
            under_fire()

        case messagetypes.COOKIE:
            cookie_eaten()

        case messagetypes.MINE:
            mine_detoned()

        case messagetypes.WALL:
            bump_into_wall()
        }

    case "Warning":
        /* A warning message can be sent when robot has to be norified on different problems which have ocurred. Currently seven
         * different warning messages can be sent */

        var warn_type int
        var warn_msg string

        _,err = fmt.Scan(&warn_type)
        if (err != nil) { fmt.Printf("Print Error getting warning type: %s\n", err) }
        _,err = fmt.Scan(&warn_msg)
        if (err != nil) { fmt.Printf("Print Error getting warning message: %s\n", err) }

        fmt.Printf("Print WARNING %d: %s\n", warn_type, warn_msg)

    case "Dead":
        return true

    default:
        fmt.Printf("Print Error message '%s' not handled\n", msg_name)
        //return true
    }

    return false
}

func resolve_battle() {
    if (r_status.Energy >= spotted_robot.Energy) {
        attack(10.0)
    } else {
        retire()
    }
}

func attack(b_energy float32) {
    fmt.Printf("Print attack\n")
    fmt.Printf("RotateAmount %d %f %f\n", messagetypes.ROBOT_ROTATE, math.Pi/3.0, math.Pi)
    fmt.Printf("RotateTo %d %f %f\n", messagetypes.CANNON_ROTATE + messagetypes.RADAR_ROTATE, -math.Pi/3.0, 0.0)
    fmt.Printf("Shoot %f\n", b_energy)
}

func retire() {
    fmt.Printf("Print Don't Shoot\n")
}

// something spotted
func wallfacer() {
    fmt.Printf("Print I saw a wall\n")
}

func cookie_spotted() {
    fmt.Printf("Print :)\n")
}

func bullet_spotted() {
    fmt.Printf("Print :o\n")
}

func mine_spotted() {
    fmt.Printf("Print D:\n")
}

func nothing_spotted() {
    fmt.Printf("Print ?\n")
}

// bump into something
func bump_into_robot() {
    fmt.Printf("Print bump into robot\n")
}

func under_fire() {
    fmt.Printf("Print bump into bullet\n")
}

func cookie_eaten() {
    fmt.Printf("Print bump into cookie\n")
}

func mine_detoned() {
    fmt.Printf("Print bump into mine\n")
}

func bump_into_wall() {
    fmt.Printf("Print bump into wall\n")
}
