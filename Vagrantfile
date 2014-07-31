# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.hostname = 'gsp'
  config.vm.box = "hashicorp/precise64"
  config.vm.network :private_network, type: :dhcp
  config.ssh.forward_agent = true

  config.vm.provider "virtualbox" do |v|
    v.gui = true
  end

  config.vm.provision :shell, path: '.vagrant-provision.sh'
end
