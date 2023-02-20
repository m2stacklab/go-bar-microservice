FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-bar-microservice

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-bar-microservice .

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-bar-microservice/out/go-bar-microservice /app/go-bar-microservice

EXPOSE 3001

# Run the binary program produced by `go install`
CMD ["/app/go-bar-microservice"]