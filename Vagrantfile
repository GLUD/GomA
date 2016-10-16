# -*- mode: ruby -*-
# vi: set ft=ruby :

ENV['VAGRANT_DEFAULT_PROVIDER'] = 'virtualbox'

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure(2) do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://atlas.hashicorp.com/search.
  config.vm.box = "centos/7"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
   config.vm.network "forwarded_port", guest: 8080, host: 18080

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
   config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
   config.vm.synced_folder ".", "/vagrant", disabled: true
   config.vm.synced_folder "src", "/home/vagrant/src"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
   config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # vb.customize ["modifyvm", :id, "--uart1", "0x3F8", 4]
  # vb.customize ["modifyvm", :id, "--uartmode1", "server", "/tmp/my_tty"]
    vb.customize ["modifyvm", :id, "--usb", "on"]
    vb.customize ["modifyvm", :id, "--usbehci", "on"]
    #vb.customize ['modifyvm', :id, '--usb', 'on']
    vb.customize ['usbfilter', 'add', '0', '--target', :id, '--name', 'SmartCard', '--vendorid', '0x2a03', '--productid', '0x0043']

    # vb.customize ["controlvm", :id, "usbattach","ysfs:/sys/devices/pci0000:00/0000:00:1d.0/usb2/2-1/2-1.2//device:/dev/vboxusb/002/007"]
   end
  #
  #config.vm.provider :vmware_fusion do |v|
   # v.vmx["usb.autoConnect.device0"]  = "0x2a03:0x0043"
   # v.vmx["usb.autoConnect.device1"]  = "..."    # for another device
  #end
  # View the documentation for the provider you are using for more
  # information on available options.
  
  #Vagrant::Config.run do |config|
  # Map COM1 port in virtual machine to 1024 port on the host
  # config.serial.forward_com1 = 1024
  #
  # Map COM2 port in virtual machine to 1025 port on the host
  #  config.serial.forward_com2 = 1025
  #
  # Override sockets path
  # Default: ~/.vagrant.d/serial
  # config.serial.sockets_path = "/path/to/sockets/dir"
  # end
  
  # Define a Vagrant Push strategy for pushing to Atlas. Other push strategies
  # such as FTP and Heroku are also available. See the documentation at
  # https://docs.vagrantup.com/v2/push/atlas.html for more information.
  # config.push.define "atlas" do |push|
  #   push.app = "YOUR_ATLAS_USERNAME/YOUR_APPLICATION_NAME"
  # end

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   sudo apt-get update
  #   sudo apt-get install -y apache2
  # string_sudoers="Defaults env_keep += \""
  # string_sudoers+="\""
  # config.vm.provision "shell", inline: "source ~/.bashrc"
   config.vm.provision "shell", inline: <<-SHELL

  SHELL

    scripts_path="scripts/"
   # config.vm.provision "shell", path: scripts_path+"set_permisive.sh"
    config.vm.provision "shell", path: scripts_path+"install_utilities.sh"
   # config.vm.provision "shell", path: scripts_path+"install_golang.sh"
    config.vm.provider "virtualbox" do |vb|
    
           config.vm.provision "shell", path: scripts_path+"install_vboxguestaditions.sh"
    end

end
