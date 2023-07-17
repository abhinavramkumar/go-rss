# Start from the latest golang base image
FROM golang:1.20.6-alpine3.18

# Add Maintainer Info
LABEL maintainer="Abhinav Ramkumar<abhinavramkumar91@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN go install github.com/cosmtrek/air@latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN apk --no-cache add ca-certificates

RUN apk add --no-cache tzdata

# Expose port 3333 to the outside world
EXPOSE 3333

EXPOSE 2345

# Build Args
ARG LOG_DIR=/app/logs

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables

ENV LOG_DIR=${LOG_DIR}/app.log

CMD ["air"]
