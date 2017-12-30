# Container volumes

## Outline

 - [Introduction](#section-introduction)
 - [Bind mounts vs volumes](#section-bind-mounts-vs-volumes)
 - How to use bind mounts in Docker
 - How to use volumes in Docker

<a name="section-introduction"></a>
## Introduction

After a container has exited because of an error or has been terminated delibaretly, all modifications which have been applied to its filesystem are gone irreversibly. Thus, it might be reasonable to place files or directories we'd like to survive a container crash or termination outside the container filesystem hierarchy (i.e. on the Docker host).

Another motivation for keeping files outside of containers is making sharing of, for instance, configuration files between multiple containers on the same host more efficient. Thereby, we can avoid repeated duplucation which usually impedes maintenance significantly.

Essentially, Docker offers two techniques, namley _bind mounts_ and _volumes_, that allow containers to access and persist files and directories on the Docker host machine.     

<a name="section-bind-mounts-vs-volumes"></a>
## Bind mounts vs volumes
