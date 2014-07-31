#!/bin/bash

export DEBIAN_FRONTEND=noninteractive

set -e
set -x

apt-get update -yq
apt-get install -yq \
  pkg-config \
  build-essential \
  xfce4 \
  gdb \
  git \
  curl \
  mercurial \
  vim-nox

if [[ ! $(go version | grep 1.2) ]] ; then
  curl -s -L 'http://golang.org/dl/go1.2.2.linux-amd64.tar.gz' | \
    tar xzf - -C /usr/local/
  ln -sfv /usr/local/go/bin/* /usr/local/bin/
fi

if [[ ! $(which liteide) ]] ; then
  curl -s -L 'http://downloads.sourceforge.net/project/liteide/X23.1/liteidex23.1.linux-64.tar.bz2' | \
    tar xjf - -C /usr/local/ --strip-components 1
fi

su - vagrant -c /vagrant/.vagrant-provision-as-vagrant.sh

echo 'Ding!'
