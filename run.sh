#!/bin/bash

##################################################################
# DEFAULT PATH
##################################################################
PACKAGE_PATH=/root/env/package
GOROOT=${PACKAGE_PATH}/go
GOPATH=$(cd "$(dirname "$0")" && pwd)
GOBIN=${GOPATH}/bin
DOWNLOAD_PATH=${GOPATH}/var/download
PATH=$PATH:${GOROOT}/bin
STAGING=real
export GOROOT GOPATH PATH STAGING

##################################################################
# DIRECTORY
##################################################################
mkdir -p ${PACKAGE_PATH} ${GOPATH}/var/log ${GOPATH}/var/db ${GOPATH}/var/cert ${GOPATH}/pkg ${GOPATH}/bin

##################################################################
# COMMANDS
# [!} Local : ng serve
##################################################################
case $1 in
    compose)
        rm -rf ${DOWNLOAD_PATH}
        mkdir -p ${DOWNLOAD_PATH}
        cd ${DOWNLOAD_PATH}

        # go
        wget https://dl.google.com/go/go1.9.2.linux-amd64.tar.gz
        tar xvfz go1.9.2.linux-amd64.tar.gz
        rm -rf ${PACKAGE_PATH}/go
        mv go ${PACKAGE_PATH}

        # go package
        cd ${GOPATH}

        go get -u github.com/robfig/cron
        go get -u github.com/gorilla/mux
        go get -u github.com/google/uuid
        go get -u github.com/gorilla/websocket
        go get -u github.com/natefinch/lumberjack
        go get -u golang.org/x/crypto/acme/autocert

        # npm
        yum update openssl
        yum install epel-release
        yum install npm

        # angular cli
        npm install -g @angular/cli

        # angular third party
        cd ${GOPATH}/angular
        npm install
    ;;
    build)
        cd ${GOPATH}/src
        go build cpwp.go
        rm -rf ${GOPATH}/bin/cpwp
        mv ${GOPATH}/src/cpwp ${GOPATH}/bin

        cd ${GOPATH}/angular
        ng build --prod --aot --base-href /prod/ --output-path=prod_
        rm -rf prod
        mv prod_ prod
        #ng build --base-href /prod/ --output-path=prod
    ;;
    build-go)
        cd ${GOPATH}/src
        go build cpwp.go
        rm -rf ${GOPATH}/bin/cpwp
        mv ${GOPATH}/src/cpwp ${GOPATH}/bin
    ;;
    build-angular)
        cd ${GOPATH}/angular
        ng build --prod --aot --base-href /prod/ --output-path=prod_
        rm -rf prod
        mv prod_ prod
        #ng build --base-href /prod/ --output-path=prod
    ;;
    start)
        nohup ${GOPATH}/bin/cpwp > /dev/null &
    ;;
    stop)
        kill -9 $(ps aux | grep -v grep | grep cpwp | awk '{print $2;}')
    ;;
    *)
        echo './run.sh { compose | build | start | stop | build | build-go | build-angular }'
    ;;
esac