# go-docker-api

### hello world api in golang with docker

If you would like to run this maunal, these are the pre-steps to run.

`go get` *//Make sure your GOPATH is properly configured or you have the .go files in the default location $HOME/go*

`go build -o main go-docker-api.go`

`docker build -t go-docker .` *//Make sure you have the Dockerfile in the same directory*

`docker run -d -p 8080:8080 go-docker`

### To do a warm up run the following curl commands

`curl localhost:8080`

  > Hello, World

`curl localhost:8080/status`

  > {"description":" your description in metadata ","version":"0.0.1","lastcommitsha":"github's commit hash"}
 
 **To stop docker container running**
 `docker ps`  *(grab the container ID)*
 
 `docker container stop` *{container ID}*
 
 **Clean up for docker if multiple unlabelled images/containers are found**
 
 `docker container prune`
 
 `docker rmi -f $(docker images --quiet --filter "dangling=true")`

### Word on unit tests
**Todo** go_test.go currently does not work -- In future required to be investigated further - temp measure - added shell script to curl each endpoint with an expected result.
