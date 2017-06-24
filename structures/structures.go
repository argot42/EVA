package structures

import "../configuration"

type Position struct {
    Distance float32
    Angle float32
}

type Map_position struct {
    X float32
    Y float32
    Angle float32
}

type Robot struct {
    Energy float32
    Teammate bool
    //spotted bool
    //pos position
}

type Object struct {
    Pos Position
    Spotted bool
}

type Robot_status struct {
    Energy float32
    Speed float32
    Cannon_angle float32
    Rotation_reached int
    Pos Map_position
}

type Game_status struct {
    Elapsed_time float32
    Robots_left int
    Game_options [configuration.GAME_OPTIONS] float32
}
