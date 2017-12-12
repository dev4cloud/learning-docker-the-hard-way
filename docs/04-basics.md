# Basic container administration

This fourth part of our Docker guide gives an overview of some of the most frequently used commands when managing Docker containers.

## Outline

 - [Creating & starting containers with `docker run`](#section-docker-run)
   - [`docker run` fundamentals](#docker-run-fundamentals)
   - [Assigning custom names to containers](#custom-container-names)
   - [Specifying a custom command](#custom-commands)
   - [Foreground mode and interactive containers](#foreground-mode)
   - [Running containers in detached mode](#detached-mode)
   - [Cleaning up containers automatically](#cleanup)

<br/>

<a name="section-docker-run"></a>
## Creating & starting containers with `docker run`

At the beginning, we want to do a high-level journey through the most basic Docker commands. Most of the following commands come with several options to tweak their behavior in different ways. We will look at them more closely in the next paragraphs.

<br/>

<a name="docker-run-fundamentals"></a>
### Fundamentals

The `docker run` command certainly belongs to the most important commands when it comes to managing the life cycle of a container. Its purpose is to start a process in a container, using an _image_ as a template:  

```
$ docker run dev4cloud/hello-docker
```

The general structure of the `docker run` command is defined as follows: <br/>

```
$ docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]
```

 - __IMAGE__: Denotes the Docker image the container is based on. If _IMAGE_ is not present on disk, it first gets "pulled" (i.e. downloaded) from an image registry (default is Docker Hub).

 - __COMMAND__ (optional): This parameter defines the executable to be launched on container startup. It is not mandatory as images usually specify a default command. If specified, _COMMAND_ takes precedence over the defaults so they can be overridden if necessary.  

 - __[ARG...]__ (optional): A list of _n_ arguments (with n >= 0) that shall be passed to _COMMAND_.

 - __[TAG]__ (optional): Points to a certain version of an _IMAGE_, defaults to _latest_.
 - __[DIGEST]__ (optional): Provides another possbility to specify a certain image version by appending its hash (SHA-256) to the image name.  

 - __[OPTIONS]__ (optional): A list of options to modify the behavior of the `docker run` command.


<br/>

<a name="custom-container-names"></a>
### Assigning custom names to containers

First, let's see how we can assign custom names to containers so we can address them by means of meaningful designations. For that purpose, we use the `--name` option:

```
$ docker run --name nightcap dev4cloud/nightcap
...
# Run from another terminal:
$ docker ps
CONTAINER ID        IMAGE                           COMMAND             CREATED             STATUS              PORTS               NAMES
50a7bf6c4dc9        dev4cloud/nightcap              "sleep 100"         3 seconds ago       Up 2 seconds                            nightcap

```

Custom names give us the opportunity to conveniently refer to specific containers in the context of other Docker commands without having to remember complicated container IDs. For example, we can inspect the container we previously launched by passing its name to the corresponding command:

```
$ docker inspect nightcap
```  

Note that container names must be unique. Consequently, the Docker daemon returns an error given that a container with a certain name already exists:

```
$ docker run --name nightcap dev4cloud/nightcap
...
# Run from another terminal:
$ docker run --name nightcap dev4cloud/nightcap
docker: Error response from daemon: Conflict. The container name "/nightcap" is already in use by container "50a7bf6c4dc93c760da50540ed43e8c49cff746eb454de88ab37614af72f618f". You have to remove (or rename) that container to be able to reuse that name.
```

<br/>

<a name="custom-commands"></a>
### Specifying a custom command   

As already mentioned, users can provide an executable that shall be launched when a Docker container gets started. The user-defined command takes precedence over an image's default process.
Paste the following Docker command into your terminal and give it a whirl:

```
$ docker run dev4cloud/hello-docker echo "Hello World!"
Hello World!
```

Here, we instruct our sample image to print "Hello World!" instead of "Hello Docker!" which is its standard behavior if nothing else is specified.

In some cases, we'll have to use the `--entrypoint` option to specify the executable that shall be launched in the container. The section about Dockerfiles will examine this further.    

<br/>

<a name="foreground-mode"></a>
### Foreground mode and interactive containers

By default, Docker starts containers in _foreground mode_, meaning that the current console is attached to STDOUT and STDERR if not specified otherwise. You can observe this behavior by running the following container:

```
$ docker run -p 9999:2000 --name copycat dev4cloud/copycat
Copycat listening at [::]:2000 ...
```

You should see that your terminal is blocked by _copycat_ waiting for some input. Open another console and use _netcat_ to send some messages to _copycat_:

```
$ echo "Hi Copycat" | nc localhost 9999
Hi Copycat
```

In order to additionally attach the containerized process's STDIN to the terminal, we have to add the `--interactive` flag (short: `-i`):

```
$ echo "hi" | docker run -i dev4cloud/hello-docker cat
hi
```    

If we'd like to go one step further and run an interactive shell in a container, we additionally have to allocate a tty or [pseudo-terminal (tty)](https://en.wikipedia.org/wiki/Pseudoterminal) and attach it to our container. Docker can be instructed to do that for us via the `--tty` (short: `-t`) flag. Accordingly, the complete command to launch a containerized shell looks like this:

```
$ docker run -it alpine sh
/ # echo "Hello Docker!"
Hello Docker!
```

<br/>

<a name="detached-mode"></a>
### Running containers in detached mode

For many applications running in containers, there's no need to keep them attached to the terminal. For instance, think of a web server which does not require a permanent link to a console as it is perfectly suitable to move the process to the background and direct STDOUT and STDERR to appropriate log files. A container running in the background is also said to be in _detached mode_. The Docker dameon must explicitly be told to start a container in detached mode by means of the `--detach` (short `-d`) flag:

```
$ docker run -d -p 9999:2000 --name copycat dev4cloud/copycat
fda616caabda3a4cb4b4517cdd647b79d879b929cd9c5af98396151adb73df0f
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
fda616caabda        dev4cloud/copycat   "./copycat"         2 minutes ago       Up 2 minutes        0.0.0.0:9999->2000/tcp   copycat
```

As you can see in the example above, instead of blocking your terminal the daemon simply responds with the container ID, giving you back the prompt immediately. You can use the `docker ps` command to convince yourself the container is actually in running state.

Assumed that a container has originally been started in detached mode, you can still attach to it afterwards by using the `docker attach` command:

```
$ docker run -d -p 9999:2000 --name copycat dev4cloud/copycat
fda616caabda3a4cb4b4517cdd647b79d879b929cd9c5af98396151adb73df0f
$ docker attach copycat
Accepted incoming connection from: 172.17.0.1:35830
...
```

In order to again detach from the container as soon as you're done, use the <kbd>Ctrl</kbd>+<kbd>p</kbd> + <kbd>Ctrl</kbd>+<kbd>q</kbd> escape sequence. However, be aware that detaching from a container requires it to have been started with the `-i` and `-t` flag set, since [escape sequences only work in TTY mode](https://groups.google.com/forum/#!msg/docker-user/nWXAnyLP9-M/kbv-FZpF4rUJ). <br/>
As an alternative, making use of the `--sig-proxy` flag of the `docker attach` command is also an option. Then, you can detach from the container via <kbd>Ctrl</kbd>+<kbd>c</kbd> without killing the container process:

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
fda616caabda        dev4cloud/copycat   "./copycat"         2 minutes ago       Up 2 minutes        0.0.0.0:9999->2000/tcp   copycat
$ docker attach --sig-proxy=false copycat
Accepted incoming connection from: 172.17.0.1:36106
...
^C
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
fda616caabda        dev4cloud/copycat   "./copycat"         3 minutes ago       Up 3 minutes        0.0.0.0:9999->2000/tcp   copycat
```

<br/>

<a name="cleanup"></a>
### Cleaning up containers automatically

Another important aspect to notice is that as soon as a container exits, it is not automatically cleaned up which means its file system is still present on disk. While this is a reasonable default behavior that allows debugging and inspecting containers after they have terminated, it may not always be desired and also might congest your disk space over time. In case you're sure you no longer need a container once it has terminated, add the `--rm` flag to the `docker run` command to make the container being _garbage collected_ automatically:

```
$ docker run --rm --name copycat dev4cloud/copycat
```

<br/>

## The `docker ps` command

One of the most important aspects when running Docker containers is to keep an overview of all containers which are currently running. This is the purpose of the `docker ps` command, which returns a list of active containers by default. You can check it out yourself by launching the following commands in your terminal:

```
$ docker run dev4cloud/nightcap
...
# Run from another terminal:
$ docker ps
CONTAINER ID        IMAGE                COMMAND             CREATED             STATUS              PORTS               NAMES
29c16ee80f01        dev4cloud/nightcap   "sleep 100"         3 seconds ago       Up 2 seconds                            happy_payne
```

You see that `docker ps` supplies some basic information most of which should be self-explanatory. As soon as the `sleep` process terminates or gets stopped explicitly the container is no longer visible in the list.
<br/>The general form of the command looks as follows:

```
$ docker ps [OPTIONS]
```

<br/>

### Printing active & stopped containers

By default, `docker ps` only shows containers which are currently running. In order to show both active containers and also containers which have exited but are not removed yet, use the `--all` (short: `-a`) flag:

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
81b61e021917        dev4cloud/copycat   "./copycat"         5 hours ago         Up 5 hours          0.0.0.0:9999->2000/tcp   copycat
$ docker ps -a
CONTAINER ID        IMAGE                    COMMAND                  CREATED             STATUS                     PORTS                    NAMES
5956a24ab62a        dev4cloud/hello-docker   "echo 'Hello Docker!'"   4 seconds ago       Exited (0) 3 seconds ago                            tender_easley
81b61e021917        dev4cloud/copycat        "./copycat"              5 hours ago         Up 5 hours                 0.0.0.0:9999->2000/tcp   copycat
```

<br/>

### Filtering for specific containers

The `--filter` (short: `-f`) flag gives us the opportunity to filter the list of containers by means of key-value pairs. For instance, we can create a query that searches for a container with a certain name:

```
$ docker ps -f name=copycat
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
81b61e021917        dev4cloud/copycat   "./copycat"         5 hours ago         Up 5 hours          0.0.0.0:9999->2000/tcp   copycat
```

A detailed list of valid filtering keys is available at the [docs](https://docs.docker.com/engine/reference/commandline/ps/#filtering).

<br/>

### Limit output to container IDs

<br/>

### Formatting output with Go templates


<br/>

### The `docker inspect` command

While the `docker ps` command gives us a high-level overview of running containers, we can use the `docker inspect` command to get a more detailed insight into a specific container. To specify the container you want to examine, use its container ID or the random name it gets assigned by the Docker Engine (we'll see how to name containers in a custom fashion shortly).

```
$ docker inspect happy_payne
[
    {
        "Id": "29c16ee80f01f954d1879a1939794e3dc8101ef213e21b3cad817f18897a1e7d",
        "Created": "2017-12-08T14:17:43.508158352Z",
        "Path": "sleep",
        "Args": [
            "100"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 8839,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2017-12-08T14:17:43.913842045Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
 ...            
```  

The `docker inspect` command is generally used as follows:

```
docker inspect [OPTIONS] NAME|ID [NAME|ID...]
```

 - __NAME__: The name of the container to inspect (either randomly generated by Docker Engine or explicitly defined by user, must be unique).

 - __ID__: The UUID of a container which can likewise be used instead of the name to point to a certain container.

 - __[OPTIONS]__ (optional): Options to further adjust the commands behavior.  

Note that more than one container can be passed to the `docker inspect` command, as it optionally takes more than one _NAME_ or _ID_ as its arguments.  


<br/>


### The `docker logs` command

The `docker logs` command allows to access the logs of e.g. a containerized application that writes its traces to STDOUT and STDERR. For instance, this might be useful for debugging error conditions.  

```
$ docker run dev4cloud/hello-docker
...
# Run from another terminal; For container ID see 'docker ps':
$ docker logs 246999
Hello Docker!
```

The command's general structure looks like this:

```
docker logs [OPTIONS] NAME|ID
```

 - __NAME__: Randomly generated or user-defined unique name of container (see above).

 - __ID__: UUID of a container (see above).

 - __[OPTIONS]__ (optional): Options to further adjust the commands behavior.  


<br/>


## Using `docker inspect` effectively

The output of the `docker inspect` command is slightly extensive and sometimes you're only interested in some defined part of the information it provides, e.g. its hostname. One option is to use `grep` in order to filter for the information you need:

```
$ docker inspect happy_payne | grep Hostname
"HostnamePath": "/var/lib/docker/containers/29c16ee80f01f954d1879a1939794e3dc8101ef213e21b3cad817f18897a1e7d/hostname",
    "Hostname": "29c16ee80f01",
```   

As you can see, using `grep` basically works but also might return additional "noisy" information we didn't search for, as `grep` does a simple string matching. Below, we'll examine how Go templates can be applied to be more precise with filtering without having to struggle with regular expressions.

## Examining container logs
