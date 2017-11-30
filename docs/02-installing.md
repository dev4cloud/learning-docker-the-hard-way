# Installing Docker

This part is about getting Docker up and running on your local machine. If, for some reason, you can't or don't want to install Docker on your local host directly, this guide also provides you with some alternatives that still offer a good experience.  

 - Docker Community (CE) and Enterprise (EE) edition
 - Docker on Linux
 - Docker on Windows/Mac
 - Alternative 1: Setting up a VM with Vagrant
 - Alternative 2: "Play with Docker"


## Docker Community (CE) and Enterprise (EE) edition

Docker is available in two editions: The Docker Community (CE) or Enterprise (EE) edition. While the free Docker CE version is perfectly suitable for development and inlcudes the most important features, the fee-based Docker EE variant comes with additional features that are only relevant for productive development. We will therefore adhere to Docker CE for this guide.

## Docker on Linux

We assume to be on a amd64 Ubuntu 16.04 machine for the following installation procedure. If you prefer another Linux distribution, for instance Debian or Fedora, or if you'd like to work on another supported hardware platform like ARM, please refer to the [Docker documentation](https://docs.docker.com/engine/installation/linux/docker-ce) and read about the installation instructions for the Linux and architecture flavor of your choice.  

 1) Install packages to enable `apt` accessing a respository over HTTPS:

 ```
 $ sudo apt-get update
 $ sudo apt-get install -y \
   apt-transport-https \
   ca-certificates \
   curl \
   software-properties-common
 ```

 2) Import Docker GPG key

 ```
 $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
 ```

3) Add amd64 package repository and update package index:

```
$ sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
$ sudo apt-get update   
```

4) Install Docker CE:

```
$ sudo apt-get install -y docker-ce
```

## Docker on Windows/Mac

Given you run on Windows, you can [download](https://docs.docker.com/docker-for-windows/install/) and run an `.msi` installer to get Docker on your machine. Note that Windows 10 Pro with Hyper-V is required to run Linux containers on a Win-based Docker host. More information on the requirements is available in the [docs](https://docs.docker.com/docker-for-windows/install/). For running Linux and Windows containers side by side, you'll find valuable input in [this blog post](https://stefanscherer.github.io/run-linux-and-windows-containers-on-windows-10/) by Stefan Scherer.  

In case you work on a Mac, there's a separate _Docker for Mac_ edition you can use. Please look [here](https://docs.docker.com/docker-for-mac/install/) for installation instructions.

## Alternative 1: Setting up a VM with Vagrant

If your Mac or Windows version is not compatible with Docker for Mac/Windows or you'd rather like to work on a VM anyway, we provide a `Vagrantfile` which enables you setting up an Ubuntu 16.04 VM with Docker installed fastly.

### Requirements

 - Oracle Virtualbox
 - [Vagrant](https://www.vagrantup.com/)

### Create the VM

1) Start the VM.

```
$ cd ../vagrant/
$ vagrant up
```

2) SSH into the running machine.

```
$ vagrant ssh
```

## Alternative 2: "Play with Docker"
