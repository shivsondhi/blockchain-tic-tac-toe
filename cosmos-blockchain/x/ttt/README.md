# TTT Module 

The module has a core of four types of files: 

-	Client side 
-	Types 
-	Handler 
-	Keeper 

The client side files get information from the user (frontend) and pass it to the `types` package files. This triggers an event which invokes the handler. The handler handles the event and calls the keeper which makes changes to the blockchain state. 

## Default files 

The `module.go` file sets up the ttt module and provides functions for tasks like initialising and exporting the genesis block, etc. The `abci.go` and `genesis.go` files allow you to define certain functions like what the module should do when a block is created or when a block is committed; when to initialise a new genesis block, etc. These files have not been modified because the Tic Tac Toe application doesn't require these functionalities. 

## Handler files

The handler is event-driven. When the main `type` message structs are updated or modified different handlers are invoked based on the message type. The `handler.go` file defines this mapping from message type to handler. The other handler files define the functions that are invoked by `handler.go`. 

### `handlerMsgCreateGame.go`

This file receives the `CreateGame` message struct which is updated when a new game is created. The handler sends the message details to the keeper which then updates the blockchain state via the Application Blockchain Interface (ABCI). The Tendermint consensus protocol is executed by the consensus layer and the blockchain state is updated. 

### `handlerMsgDeleteGame.go`

This file receives the game ID of the game to be deleted and sends the information to the keeper for deletion. 

### `handlerMsgJoinGame.go`

This file receives the `JoinGame` message struct which contains information about the opponent and new status of the game. It performs some checks for invalid game ID, incorrect game status, etc. before deciding the player roles (X and O). This is done by getting the user's public keys (using the Auth module), concatenating them and calculating its hash. The hashing algorithms used is SHA1. Once the roles have been decided, all of this information is compiled and sent to the keeper. 

### `handlerMsgPlayMove.go`

This file receives the `PlayMove` message struct which has information about a player's move and the game ID. The handler checks the move format, whether the player is out of turn or messaging the wrong game and once everything is validated, the game state is updated and sent to the keeper. This handler also checks whether the current move ends the game. If so, the winner field is also updated. 
