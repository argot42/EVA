package messagetypes

// message to robot type
//const (
//  INITIALIZE = iota
//  YOUR_NAME
//  YOUR_COLOUR
//  //  BIN_DATA_TO
//  //  ASCII_DATA_TO
//  //  LOAD_DATA_FINISHED
//  GAME_OPTION
//  GAME_STARTS
//  RADAR
//  INFO
//  COORDINATES
//  ROBOT_INFO
//  ROTATION_REACHED
//  ENERGY
//  ROBOTS_LEFT
//  COLLISION
//  WARNING
//  DEAD
//  GAME_FINISHES
//  //  SAVE_DATA
//  EXIT_ROBOT
//  UNKNOWN_MESSAGE_TO_ROBOT = -1
//)
//
//// messages from robot to type
//const (
//  ROBOT_OPTION = iota
//  NAME
//  COLOUR
//  //  LOAD_DATA
//  ROTATE
//  ROTATE_TO
//  ROTATE_AMOUNT
//  SWEEP
//  ACCELERATE
//  BRAKE
//  BREAK
//  SHOOT
//  PRINT
//  DEBUG
//  DEBUG_LINE
//  DEBUG_CIRCLE
//  UNKNOWN_MESSAGE_FROM_ROBOT = -1
//
//  //  BIN_DATA_FROM
//  //  ASCII_DATA_FROM
//  //  SAVE_DATA_FINISHED
//)

// argument type
const (
  NONE = iota
  STRING
  DOUBLE
  INT
  HEX
  //  BINDATA
)

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
const number_of_object_types int = 5
