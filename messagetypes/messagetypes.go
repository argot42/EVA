package messagetypes

// warning type
const (
  UNKNOWN_MESSAGE = iota
  PROCESS_TIME_LOW
  //  ENERGY_LOW
  //VARIABLE_OUT_OF_RANGE
  MESSAGE_SENT_IN_ILLEGAL_STATE
  UNKNOWN_OPTION
  OBSOLETE_KEYWORD
  NAME_NOT_GIVEN
  COLOUR_NOT_GIVEN
)

// game option type
const (
  ROBOT_MAX_ROTATE = iota
  ROBOT_CANNON_MAX_ROTATE
  ROBOT_RADAR_MAX_ROTATE

  ROBOT_MAX_ACCELERATION
  ROBOT_MIN_ACCELERATION

  ROBOT_START_ENERGY
  ROBOT_MAX_ENERGY
  ROBOT_ENERGY_LEVELS

  SHOT_SPEED
  SHOT_MIN_ENERGY
  SHOT_MAX_ENERGY
  SHOT_ENERGY_INCREASE_SPEED

  TIMEOUT

  DEBUG_LEVEL            // 0 - no debug 5 - highest debug level

  SEND_ROBOT_COORDINATES    // 0 - no coordinates 
                                // 1 - coordniates are given relative the starting position
                                // 2 - absolute coordinates 
)

// robot option type
const (
  SIGNAL=2                   // 0 - no signal > 1 - signal to send (e.g. SIGUSR1 or SIGUSR2) 
  SEND_SIGNAL=0              // 0 - false 1 - true
  SEND_ROTATION_REACHED=1    // 0 - no messages
                              // 1 - messages when RotateTo and RotateAmount finished
                              // 2 - messages also when sweep direction is changed

  USE_NON_BLOCKING=3          // 0 - false 1 - true 
                              // This option should always be sent as soon as possible
)

// object type
const (
  ROBOT = iota
  SHOT
  WALL
  COOKIE
  MINE
  LAST_OBJECT_TYPE
  NOOBJECT = -1
)

// number of object types
//const number_of_object_types int = 5

// part of robot number
const (
    // the sum of these rotate the combination of those objects
    ROBOT_ROTATE = 1
    CANNON_ROTATE = 2
    RADAR_ROTATE = 4
)
