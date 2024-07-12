# Dockerfile

# Stage 1: Build the Go binary
FROM golang:1.22 as build_base
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code and build the application
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o tech .

# Stage 2: Create a lightweight image
FROM alpine:latest
EXPOSE 9999

# Copy the necessary files from the builder stage
COPY --from=build_base /app/tech /app/tech
COPY --from=build_base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build_base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Install curl
RUN apk add --no-cache curl

WORKDIR /app
CMD ["./tech"]
