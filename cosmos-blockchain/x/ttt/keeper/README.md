# Keeper

The keeper interacts with the blockchain through the ABCI. Each module in the Cosmos environment has its own keeper and keeper interfaces can be used to share functionality accross modules. This keeps the module functions decoupled but while allowing for shared functionality. 

## `keeper.go` 

This file defines the `ttt` module's keeper. It shares an interface with the banking module's keeper as well as the auth module's keeper. The `storeKey` is unique for each keeper and is kept private. This ensures that only the relevant keeper can interact with the blockchain in each module. 

## `game.go`

This file defines a number of functions that allow the application module to interact with the blockchain. Some important functions are `Set`, `Get`, `GameExists` and the querier functions like `getBoard` and `listShort`. These functions are called by the handler and they access the blockchain state from within the keeper. 

## `querier.go`

This file defines the mapping of the keeper functions defined in `game.go` to query types. 
