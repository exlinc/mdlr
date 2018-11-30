#!/usr/bin/env bash
set -e
echo "Installer loaded ..."
GH_URL='https://github.com/exlinc/mdlr/releases/download'
VERSION='v1.0.0'
echo "Preparing to install mdlr@$VERSION ..."
echo "Creating temp directory ..."
mkdir .mdlr-install
cd  .mdlr-install
echo "Temp directory created ..."

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

echo "Downloading binary ..."
curl ${GH_URL}/${VERSION}/${OS}-${ARCH}-mdlr -o ${OS}-${ARCH}-mdlr
echo "Successfully downloaded binary"

echo "Installing binary ..."
cp ${OS}-${ARCH}-mdlr ${INSTALL_ROOT}mdlr
chmod +x /usr/local/bin/mdlr
echo "Successfully installed binary"

echo "Removing temp directories ..."
cd ..
rm -rf .mdlr-install
echo "Removed temp directories"

echo "Installation complete. Try running 'mdlr --help' to check your setup."
