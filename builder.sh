#!/bin/bash

cd ./cli || echo -e "\033[31;40mPath couldn't be found\033[0m\n"

# go building ...
go get -d
go build -v -o ksapify .
chmod +x ksapify

# moving binary to local/bin to make it executable
sudo mv -v ksapify /usr/local/bin/ksapify


echo -e "\033[32;40mINSTALLATION COMPLETED!\033[0m\n"

cd - || echo -e "\033[31;40mFailed to move to previous directory\033[0m\n"