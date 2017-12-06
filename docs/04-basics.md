# Basic container administration

This fourth part of our Docker guide gives an overview of some of the most frequently used commands when managing Docker containers.

## Outline

 - Basic commands
 - The `docker run` command
 - Keeping the overview of your containers
 - Inspecting containers

<br/>

## Basic commands

At the beginning, we want to do a high-level journey through the most basic Docker commands. Most of the following commands come with several options to tweak their behavior in different ways. We will look at them more closely in the next paragraphs.

<br/>

### The `docker run` command

The `docker run` command certainly belongs to the most important commands when it comes to managing the life cycle of a container. Its purpose is to start a process in a container, using an _image_ as a template:  

```
$ docker run hello-world
```

The general structure of the `docker run` command is defined as follows: <br/>

```
$ docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]
```

 - __IMAGE__: Denotes the Docker image the container is based on. If _IMAGE_ is not present on disk, it first gets "pulled" (i.e. downloaded) from an image registry (default is Docker Hub).

 - __COMMAND__ (optional): This parameter defines the executable to be launched on container startup. It is not mandatory as images usually specify a default command. If specified, _COMMAND_ takes precedence over the defaults so they can be overridden if necessary.  

 - __[ARG...]__ (optional): A list of _n_ arguments (with n >= 0) that shall be passed to _COMMAND_.

 - __TAG__ (optional): Points to a certain version of an _IMAGE_, defaults to _latest_.
 - __@DIGEST__ (optional): Provides another possbility to specify a certain image version by appending its hash (SHA-256) to the image name.  

 - __OPTIONS__ (optional): A list of options to modify the behavior of the `docker run` command.


We will examine these parameters in more detail and see them in action below.

<br/>

### The `docker ps` command

One of the most important aspects when running Docker containers is to keep an overview of all containers which are currently running. This is the purpose of the `docker ps` command, which returns a list of active containers by default. You can check it out yourself by launching the following commands in your terminal:

```
$ docker run -d debian sleep 100
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
c0497e7cad54        debian              "sleep 100"         4 seconds ago       Up 3 seconds                            pedantic_pare
```

You see that `docker ps` supplies some basic information most of which should be self-explanatory. As soon as the `sleep` process terminates or gets stopped explicitly the container is no longer visible in the list.
<br/>The general form of the command looks as follows:

```
$ docker ps [OPTIONS]
```

Again, this Docker command allows ist behavior to be tweaked by a range of options some of which we will meet later.

<br/>

### The `docker inspect` command

### The `docker logs` command

<br/>


## A closer look at `docker run`

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
 - `--interactive` : Keeps the container process's STDIN open even if no console/tty is attached.  
