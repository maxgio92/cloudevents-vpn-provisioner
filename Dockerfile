# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19 as builder

# Copy local code to the container image.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY main.go main.go
COPY pkg/ pkg/

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -v -o vpn-provisioner

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/vpn-provisioner /vpn-provisioner

# Run the web service on container startup.
CMD ["/vpn-provisioner"]
