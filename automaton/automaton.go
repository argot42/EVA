package automaton

import (
    "fmt"
    "../configuration"
)

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
            fmt.Printf("Colour 3399ff 777777\n")
        }

    case "GameOption":
        var optionnr int
        var value float32

        _,err = fmt.Scan(&optionnr)
        if (err != nil) { fmt.Printf("Print Error getting optionnr: %s\n", err) }
        _,err = fmt.Scan(&value)
        if (err != nil) { fmt.Printf("Print Error getting value: %s\n", err) }

        // $$
        fmt.Printf("Print optionnr: %d - value: %f\n", optionnr, value)
        // $$

    case "GameStarts":
        fmt.Printf("Print I'm ready\n")

    case "GameFinishes":
        fmt.Printf("Print Goodbye\n")

    case "ExitRobot":
        return true

    case "RobotsLeft":
        var robots_left int
        _,err = fmt.Scan(&robots_left)
        if (err != nil) { fmt.Printf("Print Error getting last robots left: %s\n", err) }

        // $$
        fmt.Printf("Print robots left: %d\n", robots_left)
        // $$

    case "Radar":
        /* This message gives information from the radar each turn. The radar angle is relative to the robot front;
         * it is given in radians */

        var distance, angle float32
        var observed_obj_type int

        _,err = fmt.Scan(&distance)
        if (err != nil) { fmt.Printf("Print Error getting distance: %s\n", err) }
        _,err = fmt.Scan(&angle)
        if (err != nil) { fmt.Printf("Print Error getting angle: %s\n", err) }
        _,err = fmt.Scan(&observed_obj_type)
        if (err != nil) { fmt.Printf("Print Error getting observerd object type: %s\n", err) }

        // $$
        fmt.Printf("Print distance: %f - angle: %f - type: %d\n", distance, angle, observed_obj_type)
        // $$

    case "Info":
        /* The Info message does always follow the Radar message. It gvies more general information on the state of the robot.
         * The time is the game-time elapsed since the start of the game. This is not necessarily the same as the real time elapsed,
         * due to time scale and maxe timestep. */

        var t, speed, cannon_angle float32
        _,err = fmt.Scan(&t)
        if (err != nil) { fmt.Printf("Print Error getting time: %s\n", err) }
        _,err = fmt.Scan(&speed)
        if (err != nil) { fmt.Printf("Print Error getting speed: %s\n", err) }
        _,err = fmt.Scan(&cannon_angle)
        if (err != nil) { fmt.Printf("Print Error getting cannon angle: %s\n", err) }

        // $$
        fmt.Printf("Print time: %f - speed: %f - cannon angle: %f\n", t, speed, cannon_angle)
        // $$

    case "Coordinates":
        /* Tells you the current robot position. It is only sent if the option 'Send robot coordinates is 1 or 2'. If it is 1
         * the coordinates are snet relative the starting position, which has the effect that the robot doesn't know where it is starting,
         * but only where it has moved since. */

        var x, y, angle float32
        _,err = fmt.Scan(&x)
        if (err != nil) { fmt.Printf("Print Error getting x axis value: %s\n", err) }
        _,err = fmt.Scan(&y)
        if (err != nil) { fmt.Printf("Print Error getting y axis value: %s\n", err) }
        _,err = fmt.Scan(&angle)
        if (err != nil) { fmt.Printf("Print Error getting angle value: %s\n", err) }

        // $$
        fmt.Printf("Print x: %f - y: %f - angle: %f\n", x, y, angle)
        // $$

    case "RobotInfo":
        /* If you detect a robot with your radar, this message will follow, giving some information on the robot. The opponents energy level
         * will be given in the same manner as your own energy. The second argument is only interesting in team-mode, 1 means a
         * teammate and 0 an enemy */

        var energy_level float32
        var teammate int

        _,err = fmt.Scan(&energy_level)
        if (err != nil) { fmt.Printf("Print Error getting energy level: %s\n", err) }
        _,err = fmt.Scan(&teammate)
        if (err != nil) { fmt.Printf("Print Error getting teammate/enemy flag: %s\n", err) }

        // $$
        fmt.Printf("Print energy_level: %f - teammate: %d\n", energy_level, teammate)
        // $$

    case "RotationReached":
        /* When the robot option SEND_ROTATION_REACHED is set appropriately, this message is sent when a rotation 
         * (with RotateTo or RotateAmount) has finished or the direction has changed (when sweeping). 
         * The argument corresponds to 'what to rotate' in e.g. Rotate. */

        var reached int
        _,err = fmt.Scan(&reached)
        if (err != nil) { fmt.Printf("Print Error getting rotation reached: %s\n", err) }

        // $$
        fmt.Printf("Print rotation_reached: %d\n", reached)
        // $$

     case "Energy":
         /* The end of each round the robot will get to know its energy level. It will not, however, get the exact energy, instead it is a
          * discretized into a number of "energy levels" */

        var level float32

        _,err = fmt.Scan(&level)
        if (err != nil) { fmt.Printf("Print Error getting own energy level: %s\n", err) }

        // $$
        fmt.Printf("Print own energy level: %f\n", level)
        // $$

    case "Collision":
        /* When a robot hits (or is hit by) something it gets this message. You get the angle from where the collision ocurred
         * (the angle relative the robot) and the type of object hitting you, but not how severe the collision was. This can
         * however, be determined indirectly (approximately) by the loss of energy. */

        var obj_type int
        var angle_relative_robot float32

        _,err = fmt.Scan(&obj_type)
        if (err != nil) { fmt.Printf("Print Error getting colliding object type: %s\n", err) }
        _,err = fmt.Scan(&angle_relative_robot)
        if (err != nil) { fmt.Printf("Print Error getting angle relative to robot: %s\n", err) }

        // $$
        fmt.Printf("Print Colliding object type: %d - angle relative robot: %f\n", obj_type, angle_relative_robot)
        // $$

    case "Warning":
        /* A warning message can be sent when robot has to be norified on different problems which have ocurred. Currently seven
         * different warning messages can be sent */

        var warn_type int
        var warn_msg string

        _,err = fmt.Scan(&warn_type)
        if (err != nil) { fmt.Printf("Print Error getting warning type: %s\n", err) }
        _,err = fmt.Scan(&warn_msg)
        if (err != nil) { fmt.Printf("Print Error getting warning message: %s\n", err) }

        // $$
        fmt.Printf("Print Warning type: %d - Warning message: %s\n", warn_type, warn_msg)
        // $$


    default:
        fmt.Printf("Print default -> %s\n", msg_name)
        return true
    }

    return false
}
