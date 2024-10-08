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

wget -O makefiles.tar.gz https://codeload.github.com/my-otp/makefiles/tar.gz/refs/tags/v0.0.1
tar --strip-components 1 -xf makefiles.tar.gz -C .
rm -rf makefiles.tar.gz