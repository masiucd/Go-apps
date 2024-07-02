# Go app with docker

## Description

This is a simple Go app that runs in a Docker container. The app listens on port 8080 and returns a simple message when accessed via a web browser.

## Why Docker

Docker is a platform for developers and sysadmins to develop, deploy, and run applications with containers. The use of Linux containers to deploy applications is called containerization. Containers are not new, but their use for easily deploying applications is.

## How to run the app

1. Clone the repository
2. Build the Docker image
3. Run the Docker container
4. Access the app in a web browser

### Command to clone the repository

```bash
git clone
```

### Command to build the Docker image

```bash

docker build -t go-app .
```

### Command to run the Docker container

```bash
docker run -d -p 8080:8080 go-app
```

### Access the app in a web browser

[http://localhost:8080](http://localhost:8080)
