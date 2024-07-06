# Go app with docker

## Description

This is a simple Go app that runs in a Docker container. The app listens on port 3000 and returns a simple message when accessed via a web browser.

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
docker run -d -p 3000:3000 go-app # -d flag runs the container in detached mode and -p flag maps the host port to the container port, 3000:3000 - host:container
```

### Access the app in a web browser

[http://localhost:3000](http://localhost:3000)

### Command to check the running container

```bash
docker ps
```

### Docker commands

```bash

docker ps # List all running containers
docker ps -a # List all containers
docker stop <container_id> # Stop a running container
docker rm <container_id> # Remove a container
docker images # List all images
docker rmi <image_id> # Remove an image
```

### Docker keywords meaning

- **FROM**: The base image to build the new image on top of.
- **WORKDIR**: The working directory inside the container.
- **COPY**: Copy files from the host to the container.
- **RUN**: Run a command inside the container.
- **EXPOSE**: Expose a port to the host machine.
- **CMD**: The command to run when the container starts.
- **ENTRYPOINT**: The entry point for the container.
- **ENV**: Set environment variables inside the container.
- **ARG**: Define build-time variables.
- **VOLUME**: Create a mount point for a volume.
- **USER**: Set the user or UID to run the container.
