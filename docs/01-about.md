# About Docker and container technology

This part of our Docker guide provides an introduction into Docker and the concepts behind container technology in general, and also explains why containers experienced such a massive growth of popularity over the last years.  

## Outline

 - [The motivation behind containers](#the-motivation-behind-containers)
 - [What exactly is a container?](#what-exactly-is-a-container?)
 - [What is Docker?](#what-is-docker?)


## The motivation behind containers

The increased importance of containers is closely tied to the growing popularity of microservices. The microservice architectural style drives monolithic applications towards being decomposed into several standalone functional units, running in their own processes and communicating over the network.<br/>
In order to achieve a acceptable level of isolation (i.e. separate failure and security domains) between the processes of different applications they are oftentimes hosted by dedicated virtual machines (VMs), with several VMs running side by side on a single physical commodity server. We call this _hardware virtualization_.<br/>
While hardware virtualization is a perfectly valid and still widely spread approach, it leaves lots of potential for utilizing the underlying physical resources much more efficiently. Because each VM houses a fully fledged operating system (OS), a large portion of the server hardware (physical/virtual CPUs, RAM, disk space, network bandwidth) is occupied by their OS kernels and also, not to forget, the hypervisor. This necessarily leads to the question if a certain degree of process isolation can be achieved without having to boot multiple OSes that bring along their own kernels.<br/>
In contrast to hardware level virtualization, containers provide a technique to virtualize on the OS level. A container establishes a dedicated and isolated environment for one or more userland processes. A enormous number of containers can be launched within a single host OS running a single kernel.       


## What exactly is a container?

## What is Docker?
