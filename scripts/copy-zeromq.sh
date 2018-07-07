#!/bin/bash

# Download zeromq from 3rd-party-libs
#git clone -b 'v4.2.5' --single-branch --depth 1 https://github.com/sid-agrawal/insc-3rd-party-libs
git clone https://github.com/sid-agrawal/insc-3rd-party-libs

# copy lib to correct location
sudo cp ./insc-3rd-party-libs/libzmq.so.5.1.5 /usr/local/lib/
sudo ln -s libzmq.so.5.1.5 /usr/local/lib/libzmq.so.5
sudo ln -s libzmq.so.5.1.5 /usr/local/lib/libzmq.so

# copy header file to correct locatio
sudo cp ./insc-3rd-party-libs/zmq.h /usr/local/include/
sudo cp ./insc-3rd-party-libs/zmq_utils.h /usr/local/include/

# Install zeromq driver on linux
sudo ldconfig

# Check installed
ldconfig -p | grep zmq

# Expected
############################################################
# libzmq.so.5 (libc6,x86-64) => /usr/local/lib/libzmq.so.5
# libzmq.so (libc6,x86-64) => /usr/local/lib/libzmq.so
############################################################
