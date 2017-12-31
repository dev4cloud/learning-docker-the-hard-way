# Container volumes

This part of our Docker tutorial is about how to share files and directories between multiple containers and the Docker host as well as how critical data can be persisted on the host filesystem beyond the life span of a container.

## Outline

 - [Introduction](#section-introduction)
 - [Bind mounts vs volumes](#section-bind-mounts-vs-volumes)
   - [Bind mount essentials](#bind-mounts)
   - [About Docker volumes](#docker-volumes)
 - How to use bind mounts in Docker
 - How to use volumes in Docker

<a name="section-introduction"></a>
## Introduction

After a container has exited because of an error or has been terminated delibaretly, all modifications which have been applied to its filesystem are gone irreversibly. Thus, it might be reasonable to place files or directories we'd like to survive a container crash or termination outside the container filesystem hierarchy (i.e. on the Docker host).

Another motivation for keeping files outside of containers is making sharing of, for instance, configuration files between multiple containers on the same host more efficient. Thereby, we can avoid repeated duplucation which usually impedes maintenance significantly.

Essentially, Docker offers two techniques, namley _bind mounts_ and _volumes_, that allow containers to access and persist files and directories on the Docker host machine. We'll cover both options and go through their characteristics and differences below.      

<br/>

<a name="section-bind-mounts-vs-volumes"></a>
## Bind mounts vs volumes

<a name="bind-mounts"></a>
### Bind mount essentials

What we usually understand by _mounting_ is integrating a filesystem located on a storage device (hard disk or USB drive) into our host system's filesystem hierarchy. The place within the host filesystem where a storage device is mounted is referred to as _mount point_.

Rather than a filesystem placed on an external device, a _bind mount_ mounts a local path to just another path in the same filesystem. On your Linux host or VM, you can create a bind mount by yourself with just a few commands:   

```
$ mkdir -p /tmp/src /tmp/target
$ sudo mount -o bind /tmp/src/ /tmp/target/
$ touch /tmp/target/hello
$ ls /tmp/src/
hello
```  

In the listing above, we create a `/tmp/src/` directory and mount it under the `/tmp/target/` directory. As you can see, changes applied to one of these folders are immediately reflected on the other one.

<br/>

<a name="docker-volumes"></a>
### About Docker volumes
