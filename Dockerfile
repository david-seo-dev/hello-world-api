# Adding in multi-stage build to reduce size of image

# === STAGE : Building golang application

# Smaller base image to "golang:latest"
FROM golang:alpine as builder

# Default working directory for golang
WORK $HOME/go

# Copying over binaries
COPY . .

# To remove dependency issues copying up 3rd party dependencies
COPY ./src ./src/

# Building with cmd below to allow golang binary to work in normal linux environments
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o main .

# === STAGE : RUN GOLANG EXE

# Scratch base image to reduce size
FROM scratch

# Set the Current Working Directory to go's default location
WORKDIR ~/app

# Copying over the compiled executable from previous stage step
COPY --from=builder $HOME/go/main ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

# === FINISHED

# If you wanted to build the go app inside the container then do the following:
# Set the Current Working Directory to go's default location
# WORKDIR $HOME/go
# Copy over the .go file
# COPY . .
# Copy over the dependencies to build
# COPY ./src ./src/
# Build the Go app
# RUN go build -o main .
# Then expose 8080 and run CMD ["./main"] as above
 
