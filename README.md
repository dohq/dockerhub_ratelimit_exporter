# DockerHub pull-image rate limit exporter
An exporter for prometheus to check the pull limit of the DockerHub.

## Usage
```
 -listen string
       The address to listen on for HTTP requests. (default "0.0.0.0:9767")
 -password string
       Password for use in authentication
 -username string
       Username for use in authentication
```
You can also specify a DockerHub username and password.

### Container

Env Var       | Description
:-----------: | :------
`LISTEN`      | Overwrite the default listen address and port (e.g. `:9868` would change the listen port to 9868)
`DH_USR`      | If you want to get the rate limit for a certain user set this to the desired username
`DH_PWD`      | The password of the user set via `DH_USR`, however I recommend using tokens instead and the `*_FILE` envVar
`DH_PWD_FILE` | Mutual exclusive with the `DH_PWD` var, set to the path of a file containing the password/token of the `DH_USR`. This is useful when using file based secrets, where the password is mounted into the container as a file. (The file must only contain the password/token)

## Installation
`make build`
