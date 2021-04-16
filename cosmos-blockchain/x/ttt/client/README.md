# Client Files

The client files deal with a user at two interfaces - a rest API (web application) and the command line. Both produce similar results. The files define transactions and queries i.e. GET and POST/PUT requests. 

The main command for the module (`tttcli`) and the daemon (`tttd`) are defined in the `app` folder in the project's home directory. The subcommands `tx` and `query` are also defined there. The subcommands for `tx` and `query` are defined here. 

## Command Line (cli/)

### `query.go` and `tx.go`

These files define the subcommands for the `query ttt` and `tx ttt` commands respectively. The required subcommand functions are added to their parents here. 

### `queryGame.go` and `txGame.go`

These files define the functions that were added in `query.go` and `tx.go`. They extract information that the user provides on the command line and create a new message of the required type. 


## Rest API (rest/)

### `rest.go`

This file defines what function is called when a user interacts with the given endpoints in different ways. 

### `queryGame.go` and `txGame.go`

The functions that are mapped in `rest.go` are defined in these files. Similarly to their corresponding files in the previous section, they extract the information provided by the user requests and send create new messages using it. 
