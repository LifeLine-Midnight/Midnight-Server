#!/bin/bash

APPNAME=midnightapisvr

if [ ! -d build ]; then
    mkdir build
fi

if [ ! -d build/sbin ]; then
    mkdir build/sbin
fi

if [ ! -d build/etc ]; then
    mkdir build/etc
fi

if [ ! -d build/bin ]; then
    mkdir build/bin
fi

if [ ! -d $HOME/$APPNAME ]; then
    mkdir $HOME/$APPNAME
fi

go build -o build/sbin/$APPNAME $APPNAME.go
cp -r build/sbin $HOME/$APPNAME
cp -r build/etc $HOME/$APPNAME
cp -r build/bin $HOME/$APPNAME