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

We assume to be on a amd64 Ubuntu 16.04 machine for the following installation procedure. If you prefer another Linux distribution, for instance Debian or Fedora, or if you'd like to work on another hardware platform like ARM, please refer to the [Docker documentation](https://docs.docker.com/engine/installation/linux/docker-ce) and read about the installation instructions for the Linux flavor of your choice.  

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

## Alternative 1: Setting up a VM with Vagrant

## Alternative 2: "Play with Docker"
