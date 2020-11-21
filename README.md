# DockerHub pull-image rate limit exporter
An exporter for prometheus to check the pull limit of the DockerHub.

## Usage
```
 -listen string
       The address to listen on for HTTP requests. (default "127.0.0.1:9767")
 -password string
       Password for use in authentication
 -username string
       Username for use in authentication
```
You can also specify a DockerHub username and password.

## Installation
`make build`
