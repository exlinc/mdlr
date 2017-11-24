#!/usr/bin/env bash
echo "Downloading binaries ..."
CLONE_URL='https://git.exlhub.io/exlinc/tools-mdlr-dist'
BRANCH='master'
COMMIT='HEAD'
mkdir .mdlr-install
cd  .mdlr-install
git clone -b ${BRANCH} ${CLONE_URL} mdist
cd mdist
git checkout -b ${BRANCH} ${COMMIT}
echo "Successfully downloaded binaries"

echo "Selecting binary for this system ..."
INSTALL_ROOT='/usr/local/bin/'
MACHINE_TYPE=`uname -m`
ARCH='386'
OS='linux'
if [ ${MACHINE_TYPE} == 'x86_64' ]; then
    ARCH='amd64'
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    OS='darwin'
fi
echo "Successfully selected binary for this system"

echo "Preparing to install binary by elevating permissions ..."
sudo echo "Successfully elevated permissions to install binary"

echo "Installing binary ..."
sudo cp ${OS}-${ARCH}-mdlr ${INSTALL_ROOT}mdlr
echo "Successfully installed binary"

echo "Removing temp directories ..."
cd ../..
rm -rf .mdlr-install
echo "Removed temp directories"
