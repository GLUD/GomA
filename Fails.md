En la confiracion del vagrant se presentaron algunos 
problemas que se debian al virtualizador de maquinas
virtuales por defecto (libvirt en vez de virtualbox)
para corregirlo:
echo "export VAGRANT_DEFAULT_PROVIDER=virtualbox" >> ~/.bashrc
source ~/.bashrc

Otro inconveniente fue debido a la no instalacion de
un plugin de virtualbox para vagrant, se corrigio:
sudo dnf install ruby-devel
sudo dnf install rpm-build
sudo dnf install rubygem-bundler
sudo gem install ffi -v '1.9.14'
sudo vagrant plugin install vagrant-vbguest
