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

 - VMs are heavyweight and introduce lots of overhead since each host OS comes with its own kernel. When placing multiple VMs on a single host machines, a large portion of the resources (physical/virtual CPUs, RAM, disk space, network bandwidth) is occupied by the OSes themselves and cannot be used for productive workload.


This brings us to the question if a certain degree of process isolation can alternatively be achieved without having to cope with such an overhead in terms of time, efforts and hardware resources.<br/>
In contrast to hardware level virtualization, containers provide a technique to virtualize on the OS level, establishing dedicated and isolated environments for one or more userland processes:

 - Containers are fast, meaning they can be created, started, stopped and destroyed in a fraction of seconds.

 - Containers are lightweight, i.e. they are just processes getting launched on the host OS.

 - Dozens of containers can be launched simultaneously on a single host OS, allowing for efficient sharing of the available computing resources.



## What exactly is a container?



## What is Docker?
