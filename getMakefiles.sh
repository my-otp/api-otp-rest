#!/bin/bash

if [ "$(command -v wget)" == "" ]; then
  echo "wget is required to use this"
  echo "download using:"
  echo "  arch:   sudo pacman -S wget"
  echo "  debian: sudo apt install wget"
  exit 1
fi

if [ "$(command -v tar)" == "" ]; then
  echo "tar is required to use this"
  echo "download using:"
  echo "  arch:   sudo pacman -S tar"
  echo "  debian: sudo apt install tar"
  exit 1
fi

wget -O makefiles.tar.gz https://codeload.github.com/my-otp/makefiles/tar.gz/refs/tags/v0.0.2

mkdir -p ./assets/tmp/makefiles
tar --strip-components 1 -xf makefiles.tar.gz -C ./assets/tmp/makefiles

mkdir -p makefiles
cp -r ./assets/tmp/makefiles/makefiles/* ./makefiles
cp ./assets/tmp/makefiles/Makefile .

rm -rf ./assets/tmp
rm -rf makefiles.tar.gz