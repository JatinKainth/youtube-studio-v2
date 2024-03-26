## Setting up go
Refer this [website](https://go.dev/doc/install).

## Project structure
It is as follows:


    | - api/
    | - clients/
    | - - youtube/
    | - config/
    | - db/
    | - dtype/
    | - model/
    | - .gitignore
    | - credentials.toml
    | - go.mod
    | - go.sum
    | - main.go
    | - Makefile
    | - Dockerfile

### api
contain the handler function for running the application logic

### clients
containg external 3rd party clients and function to communicate with them

### config
contain the configration logic for the application like initializing the db and credentails

### db
containg function to communicate with the db

### model
contain struct to convert db types to golang structs

### main.go
entry point of the application, this contains the server and cron startup logic and all esentials needed like credentails, db, etc.

## Running the application.
Run the application 
```bash
make run
```

