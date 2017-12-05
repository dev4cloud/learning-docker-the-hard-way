# Basic container administration

This fourth part of our Docker guide gives an overview of some of the most frequently used commands when managing Docker containers.

## Outline

 - The `docker run` command
 - Keeping the overview of your containers
 - Inspecting containers

<br/>

## The `docker run` command

### Introduction

The probably most important Docker command for managing the life cycle of a container is the `run` command. We already met it in the previous section when we launched our very first container:

```
$ docker run hello-world
```

The generalized structure of the `docker run` command is defined as follows: <br/> _docker run [OPTIONS] IMAGE [COMMAND] [ARG...]_

What it does is running a command inside a new container which is derived from the Docker image denoted by _IMAGE_. The _COMMAND_ that shall be executed is optional as images are free to specify an executable which shall be invoked by default. This is also the case for the `hello-world` image as shown above. Assumed that IMAGE is not present on disk it first gets downloaded (or "pulled" in Docker-speak) from Docker Hub.

#### Specifying a custom command   

As a next step, we will try to run a custom command inside a container. Paste the following Docker command into your terminal and give it a whirl:

```
$ docker run debian echo "Hello Docker!"
Hello Docker!
```

Here, we use the Debian image and instruct Docker to create a container that executes the `echo` shell command and prints the result to STDOUT.   


#### Interactive containers

Up to now, the containers we launched were slightly boring: All they did was executing a single command and exiting afterwards. Thus, let's make on step forward to see how we can run a shell inside a container:

```
$ docker run --interactive --tty debian bash
root@431ebce78b63:/#
```

The Docker command from above starts bash in a container and waits for input instead of exiting.

```
root@431ebce78b63:/# echo "Hello Docker!"
Hello Docker!
```

Note that we additionaly need to append two options to the `run` command for this to work:

 - `--tty` : Allocate a pseudo-terminal (tty) and attach it to the container process.
 - `--interactive` : Keeps the container process's STDIN open even if no tty is attached.  
