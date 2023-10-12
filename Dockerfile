FROM golang:alpine as builder

ENV GO111MODULE=on

# define timezone
ENV TZ Asia/Jakarta

# Install git.
RUN apk update && apk add --no-cache git  && apk add build-base

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

#Command to run the executable
CMD ["make","run-rest"]
