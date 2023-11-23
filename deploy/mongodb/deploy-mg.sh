#!/usr/bin/env bash

which docker > /dev/nill
if [ `echo $?` -ne 0 ];then
    echo "error: not install docker."
    exit 1
fi

function download_images(){
    docker pull mongo
}

function start_mg(){
    docker run -d \
    -p 27017:27017 \
    -v /all/docker_volume/mongo:/data/db \
    --name mongodb mongo
}

function main(){
    download_images
    start_mg
}
main