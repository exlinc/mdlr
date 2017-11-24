#!/usr/bin/env bash
sudo echo "Elevated permissions to install binary"
cd /usr/local/bin
URL='https://git.exlhub.io/exlinc/tools-mdlr/releases/TAG/files/'
MACHINE_TYPE=`uname -m`
ARCH='386'
OS='linux'
if [ ${MACHINE_TYPE} == 'x86_64' ]; then
    ARCH='amd64'
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    OS='darwin'
fi
echo "Downloading mdlr"
sudo curl -o mdlr ${URL}${OS}-${ARCH}-mdlr.tar.gz
echo "Downloaded, unpacking ..."
sudo tar -xzvf ${OS}-${ARCH}-mdlr.tar.gz
sudo rm ${OS}-${ARCH}-mdlr.tar.gz
sudo mv ${OS}-${ARCH}-mdlr ./mdlr
echo "Installed"
