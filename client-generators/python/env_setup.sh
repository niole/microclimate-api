#!/bin/bash

brew update && brew install pyenv
pyenv install 3.8.0
pip3 uninstall virtualenv
pip3 install virtualenv
