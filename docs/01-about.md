# About Docker and container technology

This part of our Docker guide provides an introduction into Docker and the concepts behind container technology in general, and also explains why containers experienced such a massive growth of popularity over the last years.  

## Outline

 - [The motivation behind containers](#the-motivation-behind-containers)
 - [What exactly is a container?](#what-exactly-is-a-container?)
 - [What is Docker?](#what-is-docker?)


## The motivation behind containers

The increased importance of containers is closely tied to the growing popularity of microservices. The microservice architectural style drives monolithic applications towards being decomposed into several standalone functional units, running in their own processes and communicating over the network.<br/>
In order to achieve a acceptable level of isolation (i.e. separate failure and security domains) between the processes of different applications they are oftentimes hosted by dedicated virtual machines (VMs), with several VMs running side by side on a single physical commodity server. We call this _hardware virtualization_.
While hardware virtualization is a perfectly valid and still widely spread approach, it comes with several drawbacks:

 - Provisioning VMs is slow and elaborate. They are many tools like [Vagrant](https://www.vagrantup.com/) or [Ansible](https://www.ansible.com/) which help creating and configuring VMs in a highly automated fashion. Nevertheless, preparing a VM still requires at best several minutes until it is booted, configured and ready for use.  

 - VMs are heavyweight and introduce lots of overhead since each host OS comes with its own kernel. When placing multiple VMs on a single host machines, a large portion of the resources (physical/virtual CPUs, RAM, disk space, network bandwidth) is occupied by the OSes themselves plus the hypervisor and cannot be used for productive workload.


This brings us to the question if a certain degree of process isolation can alternatively be achieved without having to cope with such an overhead in terms of time, efforts and hardware resources.<br/>
In contrast to hardware level virtualization, containers provide a technique to virtualize on the OS level, establishing dedicated and isolated environments for one or more userland processes:

 - Containers are fast, meaning they can be created, started, stopped and destroyed in a fraction of seconds.

 - Containers are lightweight, i.e. they are just processes getting launched on the host OS.

 - Dozens of containers can be launched simultaneously on a single host OS, allowing for efficient sharing of the available computing resources.



## What exactly is a container?

As already stated above, containers provide a means to perform resource sharing on the operating system level. A container can be considered a sandbox for one or more processes, which isolates them from other processes on the host but also from the OS itself. Containers are essentially based upon two kinds of features integrated in the Linux kernel:

 - __namespaces:__ Isolating containers from the host OS is achieved by leveraging Linux kernel namespaces. By means of namespaces, system wide resources like user databases, mount points and network interfaces can be hidden from the container processes by supplying each sandbox with its own copies of the underlying structures.

- __cgroups:__  Cgroups (control groups) are responsible for managing the computing resources a container is allowed to use. If, for some reason, a container exceeds its assigned share, all its processes are automatically terminated by the host system.


## What is Docker?

To clarify this right at the beginning: Docker has not invented the idea of containers. Strictly speaking, the concept of isolating processes is pretty old, e.g. UNIX systems contain the `chroot` command for decades. The Linux Container (LXC) project was the first attempt to bring namespaces, cgroups etc. together to form a complete containerization solution.<br/>
But what is Docker then? At it's core, Docker came along with two basic features which made it a disruptive tool:

 - __High-level interface:__ Docker offers a simple and convenient but powerful interface for managing the entire life cycle of containers without having to use any low-level Linux kernel APIs.

 - __Reusable images:__ In Docker, containers are created from so-called _images_. An image is a read-only template that contains a rootfs usually hosting an application along with its dependencies. Docker images enable developers to create and share container templates by making them publicly available in image repositories.   


Basically, Docker is made up of three parts (see figure below):

  1. __Docker client:__ The Docker client offers a command-line interface for talking to the Docker daemon via HTTP.

  2. __Docker daemon:__ The Docker daemon takes care of creating, monitoring and destroying containers as well as managing images. By default, it listens on a UNIX domain socket, but can also be configured to bind to a TCP socket for remote access.

  3. __Docker Hub:__ Docker Hub is a publicly accessible repository where Docker images can be published and shared with other users.

<br/>
<p align="center">
 <img src="https://docs.docker.com/engine/article-img/architecture.svg" alt="Oops" width="75%" height="75%" align="center"/>
 <br/>
 Source: https://docs.docker.com/engine/article-img/architecture.svg
</p>
<br/>

Over the last years, the Docker project constantly evolved from an originally simple container administration tool towards a holistic tool stack that covers many aspects of operating distributed applications as containers. This is an incomplete list of peripheral tools and projects which make Docker a comprehensive container ecosystem:

 - __Swarm Mode (previously Docker Swarm):__ A container orchestration framework for reliable and fault-tolerant deployments of distributed applications.  

 - __Docker Machine:__ Enables convenient provisioning and administration of Docker hosts on local infrastructure as well as several cloud providers (AWS, Microsoft Azure, DigitalOcean, ...).      

 - __Docker Compose:__ Allows the definition and deployment of multi-container applications in YAML format.
