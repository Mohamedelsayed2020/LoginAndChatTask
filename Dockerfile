FROM golang:1.13


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8090 to the outside world
EXPOSE 8090

# Run the executable
CMD ["LoginAndChatTask-app"]