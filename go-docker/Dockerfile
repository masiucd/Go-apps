# Description: Dockerfile for building the Go application
FROM golang:1.21-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod init go-apps/go-docker.com
RUN go mod tidy


# Build the Go app
RUN go build -o main .


# Expose port 8080 to the outside world
EXPOSE 8080


# Command to run the executable
CMD ["./main"]