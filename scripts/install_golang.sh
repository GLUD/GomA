#!/bin/bash
echo 'Ejecutando: install_golang.sh'

go get -d -u github.com/hybridgroup/gobot/... && go install github.com/hybridgroup/gobot/platforms/firmata

wget https://s3.amazonaws.com/gort-io/0.6.2/gort_0.6.2_linux_amd64.tar.gz

tar -xgvf gort_0.6.2_linux_amd64.tar.gz 


