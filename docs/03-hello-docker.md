# "Hello Docker"

With a functional Docker installation at hand, you're now ready to get started
and launch your first Docker container. This is the goal of this part of the tutorial.

## Outline

 - [Running your first container](#running-your-first-container)
 - [What has happened under the hood?](#what-has-happened-under-the-hood)
 - [Troubleshooting](#troubleshooting)

<br/>

## Running your first container

In order to launch your first Docker container, copy the following line, paste it into a terminal and hit _ENTER_.

```
$ docker run hello-world
```

You should see the following output:

```
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
ca4f61b1923c: Pull complete
Digest: sha256:be0cd392e45be79ffeffa6b05338b98ebb16c87b255f48e297ec7f98e123905c
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://cloud.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/engine/userguide/
```

Congratulations, you've just launched your first Docker container successfully.

<br/>

## What has happended under the hood?

Fortunately, the output of our "Hello World" container provides valuable information on the steps that were necessary for it to appear on the console:

 1. As already explained in part 1, we used the Docker client to communicate with the Docker daemon.

 2. The output above tells us something about an _image_ which has been pulled on behalf of the Docker daemon when executing our command. Do not think about images for too long for now, just keep them in mind as templates Docker uses to create containers.

 3. We were immediately brought back to our shell prompt after the container had sent its output to the terminal. So it seems like the container terminated once its job had been finished.  

 4. The first lines of the output also show some weird sequences of numbers and letters like _ca4f61b1923c_. These are SHA-256 hashes which pose a cryptographic property of Docker images and, among others, serve the purpose of integrity verification. More on this later.

<br/>

## Troubleshooting  

If you don't receive the output shown above when executing the `docker run` command, make sure the following conditions are met:

#### 1. Ensure that Docker has been installed properly

```
$ docker -v
Docker version 17.09.0-ce, build afdb6d4
```

<br/>

#### 2. Check if the Docker daemon is active

```
$ sudo systemctl status docker.service
docker.service - Docker Application Container Engine
   Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
   Active: active (running) since So 2017-12-03 10:59:12 CET; 5h 10min ago
     Docs: https://docs.docker.com
 Main PID: 1350 (dockerd)
    Tasks: 21
   Memory: 83.5M
      CPU: 41.396s
```

If your Docker daemon is reported to be not running, try restarting it via your init system, e.g. systemd in Ubuntu:

```
$ sudo systemctl restart docker.service
```

<br/>

#### 3. Check if user has permissions to interact with the Docker daemon

In order to be allowed to interact with the Docker daemon, your user has to meet at least one of the following requirements:

##### a) User can run commands with superuser privileges

Start your Docker commands with `sudo`:
```
$ sudo docker run hello-world
```

##### b) User belongs to the _docker_ group

The _docker_ group is created during the installation process and allows its members to access the Docker daemon without `sudo`. If you haven't already, add your user to this group:

```
$ whoami
moby
$ sudo usermod -aG docker moby
```
Don't forget to sign out and in again for the modifications to be applied. Then, check again if your user is now part of the _docker_ group:

```
$ groups
moby sudo docker
```  

We generally recommend to follow approach b) since this saves you the need of having to type `sudo` again and again and significantly reduces the risk of causing damage on your system due to faulty commands executing with root permissions.    
