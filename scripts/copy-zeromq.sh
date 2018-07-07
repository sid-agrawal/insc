#!/bin/bash

# Download zeromq from 3rd-party-libs
#git clone -b 'v4.2.5' --single-branch --depth 1 https://github.com/sid-agrawal/insc-3rd-party-libs
git clone https://github.com/sid-agrawal/insc-3rd-party-libs

# copy lib to correct location
cd insc-3rd-party-libs
sudo ./install.sh
cd ../

