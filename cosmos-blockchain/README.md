# Tic Tac Toe on the blockchain

**ttt** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

You will first need to install the Go programming language and the starport files. Once installed go into the application folder and run the following command:  

```
starport serve
```

The `serve` command installs dependencies, initializes and runs the application.

## Accounts

Initialization parameters of the app are stored in `config.yml`.
This contains all the user accounts with their account information as well as all the validators. The list of current user accounts is: 
-	alice 
-	bob
-	charlie
-	dave
-	eve
-	frank

Currently all users are validators.

## User Guide

Once you have run `starport serve` from the command line, open up a new terminal and now you can interact with the blockchain from here. Note that you may be prompted to add `GOPATH/bin` to your `PATH` variable which can be done by running `export PATH=$PATH:$(go env GOPATH)/bin` in a linux terminal or by adding the path to your `~/.bashrc file`. 

Here are some useful commands to interact with the application: 


| Command                                                  | Use                                                                     |
| -------------------------------------------------------- | ----------------------------------------------------------------------- |
| `tttcli tx ttt create-game --from alice`                 | Will create a new game with Alice as the creator                        |
| `tttcli tx ttt join-game $(gameID) --from bob`           | Will join the game with the specified gameID and Bob as the opponent    |
| `tttcli tx ttt play-move $(gameID) $(move) --from alice` | Will play the specified move in the specified gameID as Alice           |
| `tttcli query ttt list-short`                            | Will list short state of all games                                      |
| `tttcli query ttt get-board $(gameID)`                   | Will list the board for game with the specified gameID                  |
| `tttcli query ttt list-game`                             | Will list full state of all games                                       |
| `tttcli query ttt get-game $(gameID)`                    | Will list full state of the game with the specified gameID              |


Sometimes the state may not update immediately on querying the blockchain. If you wait for about 10 seconds and try again it should be updated. 

## The Game

### Game state

The game state includes the following fields:
-	Creator address (included in short state)
-	Game ID (included in short state)
-	Opponent address
-	Game status [open, inprogress or finished] (included in short state)
-	Game state [a string that represents the game board] 
-	Address of player playing as X 
-	Address of player playing as O 
-	Next move [X or O] 
-	Winner [null, X, O or draw] (included in short state)

### Flow of the game

- 	A player sends a transaction to the app to create a game. 
-	A game is added to the blockchain with the 'open' status. 
-	Another player can then send a transaction to accept any open game using its gameID. 
-	The game's state now changes to 'inprogress'. 
-	The public keys of both players are concatenated and hashed. If the first bit is 0, the creator takes 'O' and the opponent takes 'X' and vice versa. 
-	Both users submit transactions to the app to play their moves. 
-	The game state changes to 'finished' when the game is over and the winner is updated. If nobody wins, the match ends in a draw. 
-	Multiple games can be played concurrently. 

### Rules of the game 

-	Always pass the `--from` flag with the three commands as specified in the guide above. The application will throw an error if the sender is not specified.
-	X starts.
-	To play a move pass a number from 1-9. The mapping for the numbers to the position on the board is provided below.
-	A user cannot join a game that is `inprogress` or `finished`. 
-	A user cannot play a move in a game that is not `inprogress`.
-	A user cannot play a move in a game where he/she is not listed as either creator or opponent. 
-	A user cannot play a move in a game out of turn. 


| Tic   | Tac   | Toe   |
| :---: | :---: | :---: |
|  1    |  2    |  3    |
|  4    |  5    |  6    |
|  7    |  8    |  9    |


## Helpful tips

-	Use the `list-short` query command to check for open games.
-	Use the `get-board` query command to check the board for a particular game.
-	Use the command `tttcli keys show alice -a` to show Alice's account address. You can replace `alice` with any other user name.
-	Use the command `tttcli query account $(tttcli keys show bob -a)` to show Bob's user information. 


<!-- ## Other Information (Provided by Starport)

### UI on Github Pages

Click the link below, and scroll down until you see it get her pages. Then, select the branch gh-pages.

[Github Pages Setings](https://github.com/(shivsondhi/ttt/settings/)

After you do that you can visit your chain's UI at:

https://shivsondhi.github.io/ttt

This is especially useful when you would like to rapidly iterate on a live user interface. Remember, each community member can have their own github pages instance, allowing your community to mix-and-match front ends.


### Docker Images And Pi Images

In order for Docker images and Raspberry Pi images to build successfully, please add your docker hub credentials as [secrets](https://github.com/shivsondhi/ttt/settings/secrets/actions)

Add these:

DOCKERHUB_USERNAME
DOCKERHUB_TOKEN

You can get the token [here](https://hub.docker.com/settings/security) -->

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
