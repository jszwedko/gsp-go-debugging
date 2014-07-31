#!/bin/bash
set -x
set -e

mkdir -p ~/gopath

ln -svf /vagrant/.vagrant-skel/bashrc /home/vagrant/.bashrc
ln -svf /vagrant/.vagrant-skel/profile /home/vagrant/.profile
