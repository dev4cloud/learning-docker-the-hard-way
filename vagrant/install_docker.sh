#!/bin/bash

# Update package index.
sudo apt-get update

# Install installation dependencies.
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

# Import Docker GPG key.
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Add APT repository.
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

# Update package index again.
sudo apt-get update

# Install Docker CE. 
sudo apt-get install -y docker-ce

# Add default user to docker group.
sudo usermod -aG docker ubuntu


