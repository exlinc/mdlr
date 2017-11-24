#!/usr/bin/env bash
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
curl -o mdlr ${URL}${OS}-${ARCH}-mdlr
echo "Installed"