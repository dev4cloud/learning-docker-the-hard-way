# Basic container administration

This fourth part of our Docker guide gives an overview of some of the most frequently used commands when managing Docker containers.

## Outline

 - [Creating & starting containers with `docker run`](#section-docker-run)
   - [Fundamentals](#docker-run-fundamentals)
   - [Assigning custom names to containers](#custom-container-names)
   - [Specifying a custom command](#custom-commands)
   - [Foreground mode and interactive containers](#foreground-mode)
   - [Running containers in detached mode](#detached-mode)
   - [Cleaning up containers automatically](#cleanup)
 - [Starting, stopping & removing existing containers](#section-start-stop-remove)
   - [Fundamentals](#start-stop-remove-fundamentals)
   - [Stopping an active Docker container](#docker-stop)
   - [Starting a stopped container](#docker-start)
   - [Removing an obsolete container](#docker-rm)
 - [Viewing containers with `docker ps`](#section-docker-ps)
   - [Fundamentals](#docker-ps-fundamentals)
   - [Printing active & stopped containers](#active-and-stopped)
   - [Filtering `docker ps` output](#container-filtering)
   - [Only show container IDs](#container-ids)
   - [Formatting output with Go templates](#docker-ps-formatting)
 - [Examining containers with `docker inspect`](#section-docker-inspect)
   - [Fundamentals](#docker-inspect-fundamentals)
   - [Formatting output](#docker-inspect-formatting)
 - [Viewing container logs with `docker logs`](#section-docker-logs)
   - [Fundamentals](#docker-logs-fundamentals)
   - [Tracking container logs](#docker-logs-follow)
   - [Controlling the number of logs displayed](#docker-logs-tail)


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
$ docker run -it --name shell dev4cloud/shell
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

<a name="section-start-stop-remove"></a>
## Starting, stopping & removing existing containers

<a name="start-stop-remove-fundamentals"></a>
### Fundamentals

Besides launching containers with `docker run`, we also need commands to stop them once they're no longer required. As their filesystem is kept on disk even if a container terminates (assumed that is has __NOT__ been intended for automatic cleanup with the `--rm` flag), stopped containers can be restarted at a later point in time. Otherwise, if we're sure that a stopped container will never be needed in the future, we can permanently delete it from disk.        

<br/>

<a name="docker-stop"></a>
### Stopping an active container

Use the `docker stop` command to stop an active Docker container. If the command returns the container ID or name, the execution has been successful.

```
$ docker stop copycat
copycat
```

The command's general structure reveals that more than one container can be scheduled for termination within a single command execution:

```
docker stop [OPTIONS] NAME|ID [NAME|ID ...]
```

<br/>

<a name="docker-start"></a>
### Starting an inactive container

If a container that has been stopped shall be started again, use the `docker start` command:

```
$ docker start copycat
```

The command's formal structure looks as follows:

```
docker start [OPTIONS] NAME|ID [NAME|ID ...]
```

Like `docker stop`, this command takes more than one container name or ID as an argument, i.e. multiple containers can be started with a single command.

If the stopped container had a tty allocated and you want to start it as an interactive container again, add the `--interactive` (short: `-i`)
flag to attach to the container's STDIN again:

```
$ docker start -i shell
/ #
```

The `docker start` command launches containers in the background by default. If you want to restart a container in foreground mode use the `--attach` (short: `-a`) flag which connects to the container's STDIN and STDERR and enables signal forwarding:

```
$ docker start -a copycat
Copycat listening at [::]:2000 ...
```

However, note that you can only detach your terminal from the container without stopping it again if it has initially been started with the `-i` and `-t` flags set. Additionally, you must add the `-i` flag to the `docker start` command for this to work:

```
$ docker start -i -a copycat
Copycat listening at [::]:2000 ...
^p^q
read escape sequence
```

In most situations, simply starting the container in the background and attaching to it via the [`docker attach` command](#detached-mode) might be the easiest way.

<br/>

<a name="docker-rm"></a>
### Removing an obsolete container

If a container and its filesystem are no longer used, it is always a good idea to delete them in order to free up some disk space. To remove a container permanently, pass it to the `docker rm` command:

```
$ docker rm copycat
copycat
```

Like the previous commands, the `docker rm` command accepts multiple containers as an argument:

```
docker rm [OPTIONS] NAME|ID [NAME|ID ...]
```

The execution has been successful if the Docker daemon returns the container name or ID back to you. Note that a container that shall be removed must be in stopped state. Otherwise, the Docker daemon responds with an error message:

```
$ docker rm copycat
Error response from daemon: You cannot remove a running container fda616caabda3a4cb4b4517cdd647b79d879b929cd9c5af98396151adb73df0f. Stop the container before attempting removal or force remove
$ docker stop copycat
copycat
$ docker rm copycat
copycat
```

Again, the Docker daemon returns the container name if the removal was successful. As having to use two commands (`stop` and `rm`) might be tedious, you can indeed force running containers to be stopped and removed with a single command:

```
$ docker rm -f copycat
copycat
```

The `--force` (short: `-f`) flag instructs the Docker daemon to forcfully kill the container process by sending it a _SIGKILL_ signal.

<br/>

<a name="section-docker-ps"></a>
## Viewing containers with `docker ps`

<a name="docker-ps-fundamentals"></a>
### Fundamentals

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

<a name="active-and-stopped"></a>
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

<a name="container-filtering"></a>
### Filtering `docker ps` output

The `--filter` (short: `-f`) flag gives us the opportunity to filter the list of containers by means of key-value pairs. For instance, we can create a query that searches for a container with a certain name:

```
$ docker ps -f name=copycat
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
81b61e021917        dev4cloud/copycat   "./copycat"         5 hours ago         Up 5 hours          0.0.0.0:9999->2000/tcp   copycat
```

A detailed list of valid filtering keys is available at the [docs](https://docs.docker.com/engine/reference/commandline/ps/#filtering).

<br/>

<a name="container-ids"></a>
### Only show container IDs

Sometimes it can be helpful to only list the IDs of containers and ignore other information. The `--quiet` (short: `-q`) flag serves that purpose:

```
$ docker ps -q
b370247ff648
```

An exemplary scenario where this is very helpful is when you want to stop and/or delete all running containers on your system:

```
$ docker stop $(docker ps -q)
b370247ff648
```

The command above fetches the IDs of all running containers in a subshell and then passes them to the `docker stop` command.  

<br/>

<a name="docker-ps-formatting"></a>
### Formatting output with Go templates

Docker's standard formatting for the `docker ps` output is very extensive and frequently only a subset of the properties of containers is actually needed. Docker offers us the possibility of customizing this output according to our needs by means of _Go templates_. Go templates consist of placeholders which represent a certain container attribute:

```
$ docker ps --format "{{.Names}}\t{{.Image}}\t{{.Size}}"
copycat    dev4cloud/copycat    0B (virtual 2.72MB)
```

Note that you need to specify the `table` directive within the template if you want the column heades to be printed:

```
$ docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Size}}"
NAMES               IMAGE               SIZE
copycat             dev4cloud/copycat   0B (virtual 2.72MB)
```

In order to make your custom `docker ps` formatting permanent and avoid having the specify it via the `--format` flag on every invocation, go to `~/.docker/config.json` and add your Go template pattern to the Docker CLI configuration:

```
$ cat ~/.docker/config.json
{
  "psFormat": "table {{.Names}}\t{{.Image}}\t{{.Size}}"
}
```

For a comprehensive list of valid template placeholders head over to the [documentation](https://docs.docker.com/engine/reference/commandline/ps/#formatting).

<br/>


<a name="section-docker-inspect"></a>
## The `docker inspect` command

<a name="docker-inspect-fundamentals"></a>
### Fundamentals

While the `docker ps` command gives us a high-level overview of running containers, we can use the `docker inspect` command to get more detailed insights into a specific container. While we solely use it to examine containers in this section, the `docker inspect` command can also be used to analyse other Docker objects like images. Applying the command to a container yields a very detailed, JSON formatted output with much information about it:


```
$ docker inspect copycat
[
    {
        "Id": "963a8ca447c12901902f3ddc474e195862a1f575c017095126e9678f53b107e5",
        "Created": "2017-12-17T14:54:45.211712127Z",
        "Path": "./copycat",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 15523,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2017-12-17T14:54:45.744835592Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
     ...            
```  

The `docker inspect` command is generally used as follows:

```
docker inspect [OPTIONS] NAME|ID [NAME|ID...]
```

 - __NAME__: The name of the container that shall be inspected.

 - __ID__: A container's UUID.

 - __[OPTIONS]__ (optional): Options to further adjust the commands behavior.  

Note that more than one container can be passed to the `docker inspect` command, as it optionally accepts more than one _NAME_ or _ID_ as arguments.  


<br/>

<a name="docker-inspect-formatting"></a>
### Formatting output

You already saw that the output of the `docker inspect` command is slightly extensive. In most cases, you'll probably be interested in only some specific container metadata, e.g. its IP address or hostname. In order to filter the JSON output for the interesting peace of information, one approach is to use `grep`:

```
$ docker inspect copycat | grep Hostname
    "HostnamePath": "/var/lib/docker/containers/963a8ca447c12901902f3ddc474e195862a1f575c017095126e9678f53b107e5/hostname",
        "Hostname": "963a8ca447c1",
```   

As you can see, `grep` basically works but also might return additional "noisy" information we didn't search for explicitly since `grep` does a simple string matching. Of course we couldfeed some regular expressions to `grep` to filter more precisely, but struggling with regex could be slow and annoying. Luckily, as we already did with `docker ps`, we can use Go templates in conjunction with the `docker inspect` command to make filtering more comfortable:  

```
$ docker inspect -f 'Hostname: {{.Config.Hostname}}'
Hostname: 963a8ca447c1
```

In the majority of cases, extracting information from the output JSON is as easy as following its structure and picking out the field(s) of interest. In the case of arrays or maps where you have an arbitrary number of elements, you can loop over their entries to produce a well-formatted text output:

```
$ docker inspect -f '{{range $net, $conf := .NetworkSettings.Networks}} {{$net}} -> {{.IPAddress}} {{end}}' copycat 
 bridge -> 172.17.0.2    
```

<br/>


<a name="section-docker-logs"></a>
## Viewing container logs with `docker logs`

<a name="docker-logs-fundamentals"></a>
### Fundamentals

The `docker logs` command shows the logs of a containerized application that writes to STDOUT and/or STDERR. For instance, this can be very helpful in case of error conditions as this command enables users to access a container's output easily and quickly.  

```
$ docker run -d -p 9999:2000 --name copycat dev4cloud/copycat
1a84e0b4c673d75938fd6884f9666231ba80602ce05c2f777353b42bd2900bf4
$ docker logs copycat
Copycat listening at [::]:2000 ...
Accepted incoming connection from: 172.17.0.1:44854
...
```

The command's general structure looks like this:

```
docker logs [OPTIONS] NAME|ID
```

 - __NAME__: Randomly generated or user-defined unique name of container (see above).

 - __ID__: Ccontainer's UUID (see above).

 - __[OPTIONS]__ (optional): Options to further adjust the commands behavior.  


<br/>

<a name="docker-logs-follow"></a>
### Tracking container logs

In order to understand the exact behavior of an application running inside a container, in can sometimes help to check the logs in real time instead of examining them after a crash or graceful termination. For that purpose, the `docker logs` command provides the `--follow` (short: `-f`) flag, which waits for incoming log records and displays them as they arrive:

```
$ docker logs -f copycat
Copycat listening at [::]:2000 ...
Accepted incoming connection from: 172.17.0.1:44854
Accepted incoming connection from: 172.17.0.1:45124
Accepted incoming connection from: 172.17.0.1:45128
Accepted incoming connection from: 172.17.0.1:45132
Accepted incoming connection from: 172.17.0.1:45136
Accepted incoming connection from: 172.17.0.1:45140
Accepted incoming connection from: 172.17.0.1:45144
...
```

<br/>

<a name="docker-logs-tail"></a>
### Controlling the number of logs displayed

Depending on the containerized process, its uptime and logging behavior, it might be hard to keep an overview of the logs returned by `docker logs` as the command simply returns everything that has been captured since the process has been started. To bypass this problem, the `docker logs` command allows us to specify the number of log entries to display, starting with the most recent one. In other words, the `--tail` flag makes the command to only show the last _n_ lines. The parameter _n_ is optional and defaults to _all_:

```
$ docker logs --tail 3 copycat
Accepted incoming connection from: 172.17.0.1:45136
Accepted incoming connection from: 172.17.0.1:45140
Accepted incoming connection from: 172.17.0.1:45144
```


