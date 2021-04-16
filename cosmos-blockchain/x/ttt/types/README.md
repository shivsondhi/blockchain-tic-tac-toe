# Types

All files here belong to the `types` package. These files define the message types, error types and keeper interface types that will be used by the `ttt` module. 

## Error types (`errors.go`)

This file defines the different kinds of errors that the program may encounter during execution. They are registered with the SDK's `types` package and are used in the Handler. 

## Message types 
A message type file is defined for each type of message - create game, join game, play move and delete game. These files define a constructor function to create a new message and perform some basic validation of the message contents. 

## Game type 
The `Game` struct defines the structure for our game state on the blockchain. The `TypeGame.go` file defines this struct and the keeper uses it to send state information to the blockchain. 
